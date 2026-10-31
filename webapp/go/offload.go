package isureal

// まとめて返す画像の署名（Memories・エクスポート）と、受け取った JPEG の検証を、CPU の空いている台へ逃がす。
// JPEG の判定は手元と同じ jpegcheck.Check で行う。
// 逃がし先は、署名が ISUREAL_OFFLOAD_PEERS、JPEG が ISUREAL_OFFLOAD_JPEG_PEERS（頼む順、カンマ区切りの host:port）。
// 署名は europe1・america1 を先に使う。asia2 の PostgreSQL は書き込みを受けなくなったので、JPEG は asia2 にも頼む。失敗した逃がし先は数秒外し、
// 署名は、どの逃がし先も同時に頼める数（offloadPeerSlots）まで埋まっている・失敗した・遅いときは手元で計算する。
// JPEG は、逃がし先 1 台につき 32 件（往復を待つ間も逃がし先の CPU を埋める）まで同時に検証し、あふれた分は先着順に待たせる
// （一斉の投稿で全件を同時に計算すると、どの投稿も揃って遅れる）。手元の CPU は他の API に残し、逃がし先が外れたときだけ CPU 数の件数まで検証する。逃がし先は鍵を受け取るので、
// 受け付けはプライベートネットワーク（192.168.0.0/16）からだけにし、nginx には出さない。

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"example.com/isureal-company/isureal/jpegcheck"
)

const offloadPort = ":8090"

const (
	offloadSignTimeout = 2 * time.Second
	offloadJPEGTimeout = 2 * time.Second
	offloadPeerBackoff = 5 * time.Second
	offloadPeerSlots   = 8
)

type offloadPeer struct {
	url       string
	downUntil atomic.Int64
	slots     chan struct{}
}

var offloadPeerByAddr = map[string]*offloadPeer{}

func offloadPeersFromEnv(key string) []*offloadPeer {
	var peers []*offloadPeer
	for _, addr := range strings.Split(os.Getenv(key), ",") {
		if addr = strings.TrimSpace(addr); addr != "" {
			peer := offloadPeerByAddr[addr]
			if peer == nil {
				peer = &offloadPeer{url: "http://" + addr, slots: make(chan struct{}, offloadPeerSlots)}
				offloadPeerByAddr[addr] = peer
			}
			peers = append(peers, peer)
		}
	}
	return peers
}

var (
	offloadPeers     = offloadPeersFromEnv("ISUREAL_OFFLOAD_PEERS")
	offloadJPEGPeers = offloadPeersFromEnv("ISUREAL_OFFLOAD_JPEG_PEERS")
)

var offloadClient = &http.Client{Transport: &http.Transport{
	DialContext:         (&net.Dialer{Timeout: 300 * time.Millisecond}).DialContext,
	MaxIdleConns:        256,
	MaxIdleConnsPerHost: 64,
	IdleConnTimeout:     90 * time.Second,
}}

// 空きのある逃がし先へ近い順に頼む。つながらなければ次へ、時間切れなら手元で計算させる
func offloadCall(peers []*offloadPeer, path string, body []byte, timeout time.Duration) ([]byte, error) {
	now := time.Now().UnixNano()
	for _, peer := range peers {
		if peer.downUntil.Load() > now {
			continue
		}
		select {
		case peer.slots <- struct{}{}:
		default:
			continue
		}
		res, err := offloadPost(peer.url+path, body, timeout)
		<-peer.slots
		if err == nil {
			return res, nil
		}
		peer.downUntil.Store(time.Now().Add(offloadPeerBackoff).UnixNano())
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, err
		}
	}
	return nil, errors.New("no offload peer")
}

func offloadPost(url string, body []byte, timeout time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	res, err := offloadClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("offload status %d", res.StatusCode)
	}
	return data, nil
}

var errInvalidJPEG = errors.New("invalid JPEG")

// 逃がし先の JPEG の検証の空き
var jpegPeerSlots = func() chan *offloadPeer {
	ch := make(chan *offloadPeer, len(offloadJPEGPeers)*32)
	for range 32 {
		for _, peer := range offloadJPEGPeers {
			ch <- peer
		}
	}
	return ch
}()

// 手元の JPEG の検証の空き。逃がし先がすべて外れているときだけ使う
var jpegLocalSlots = make(chan struct{}, runtime.NumCPU())

func jpegPeerAvailable() bool {
	now := time.Now().UnixNano()
	for _, peer := range offloadJPEGPeers {
		if peer.downUntil.Load() <= now {
			return true
		}
	}
	return false
}

// jpegcheck.Check と同じ判定を返す。逃がし先が使えないときは手元で検証する
func checkJPEG(data []byte) (width, height int, err error) {
	if jpegPeerAvailable() {
		peer := <-jpegPeerSlots
		ok := peer.downUntil.Load() <= time.Now().UnixNano()
		var res []byte
		if ok {
			res, err = offloadPost(peer.url+"/jpeg", data, offloadJPEGTimeout)
			ok = err == nil && len(res) == 9
			if !ok {
				peer.downUntil.Store(time.Now().Add(offloadPeerBackoff).UnixNano())
			}
		}
		jpegPeerSlots <- peer
		if ok {
			if res[0] != 1 {
				return 0, 0, errInvalidJPEG
			}
			return int(binary.BigEndian.Uint32(res[1:5])), int(binary.BigEndian.Uint32(res[5:9])), nil
		}
	}
	jpegLocalSlots <- struct{}{}
	defer func() { <-jpegLocalSlots }()
	return jpegcheck.Check(bytes.NewReader(data))
}

type signTarget struct {
	key     []byte
	imageID string
}

// まとめて返す画像の署名を逃がし先で計算し、署名のキャッシュに入れておく。
// 失敗したら何もしない（signedImageURL が手元で計算する）
func prefetchImageSignatures(targets []signTarget) {
	if len(offloadPeers) == 0 {
		return
	}
	exp := imageURLExp()
	var body []byte
	var keys []imageSigKey
	seen := make(map[imageSigKey]bool, len(targets))
	imageSignatureCacheMu.RLock()
	gen := imageSignatureCache[exp]
	for _, t := range targets {
		k, ok := makeImageSigKey(t.key, t.imageID)
		if !ok || seen[k] {
			continue
		}
		seen[k] = true
		if _, ok := gen[k]; ok {
			continue
		}
		keys = append(keys, k)
		body = appendSignRequest(body, t.key, t.imageID, exp)
	}
	imageSignatureCacheMu.RUnlock()
	if len(keys) == 0 {
		return
	}
	res, err := offloadCall(offloadPeers, "/sign", body, offloadSignTimeout)
	if err != nil || len(res) != len(keys)*32 {
		return
	}
	imageSignatureCacheMu.Lock()
	putImageSignaturesLocked(exp, keys, res)
	imageSignatureCacheMu.Unlock()
}

// 1 件 = 鍵の長さ(1) 鍵 exp(8) 画像 ID の長さ(2) 画像 ID
func appendSignRequest(b, key []byte, imageID string, exp int64) []byte {
	b = append(b, byte(len(key)))
	b = append(b, key...)
	b = binary.BigEndian.AppendUint64(b, uint64(exp))
	b = binary.BigEndian.AppendUint16(b, uint16(len(imageID)))
	return append(b, imageID...)
}

type signRequest struct {
	key     []byte
	imageID string
	exp     int64
}

func parseSignRequests(b []byte) ([]signRequest, error) {
	var reqs []signRequest
	for len(b) > 0 {
		n := int(b[0])
		if len(b) < 1+n+8+2 {
			return nil, errors.New("short request")
		}
		key := b[1 : 1+n]
		exp := int64(binary.BigEndian.Uint64(b[1+n:]))
		m := int(binary.BigEndian.Uint16(b[1+n+8:]))
		b = b[1+n+8+2:]
		if len(b) < m {
			return nil, errors.New("short request")
		}
		reqs = append(reqs, signRequest{key: key, imageID: string(b[:m]), exp: exp})
		b = b[m:]
	}
	return reqs, nil
}

var privateNetwork = &net.IPNet{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(16, 32)}

func offloadHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /jpeg", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res := make([]byte, 9)
		if width, height, err := jpegcheck.Check(bytes.NewReader(data)); err == nil {
			res[0] = 1
			binary.BigEndian.PutUint32(res[1:5], uint32(width))
			binary.BigEndian.PutUint32(res[5:9], uint32(height))
		}
		w.Write(res)
	})
	mux.HandleFunc("POST /sign", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		reqs, err := parseSignRequests(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res := make([]byte, len(reqs)*32)
		var next atomic.Int64
		var wg sync.WaitGroup
		for range min(runtime.NumCPU(), len(reqs)) {
			wg.Go(func() {
				for i := int(next.Add(1) - 1); i < len(reqs); i = int(next.Add(1) - 1) {
					copy(res[i*32:], computeImageSignature(reqs[i].key, reqs[i].imageID, reqs[i].exp))
				}
			})
		}
		wg.Wait()
		w.Write(res)
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if ip := net.ParseIP(host); err != nil || ip == nil || !privateNetwork.Contains(ip) {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		mux.ServeHTTP(w, r)
	})
}

// 逃がし先として受け付ける（戻らない）
func ServeOffload() error {
	return http.ListenAndServe(offloadPort, offloadHandler())
}
