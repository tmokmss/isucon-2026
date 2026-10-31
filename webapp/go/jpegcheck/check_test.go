package jpegcheck

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/fs"
	"math/rand/v2"
	"os"
	"path/filepath"
	"testing"
)

type verdict struct {
	ok     bool
	w, h   int
	panics bool
}

func stdVerdict(data []byte) (v verdict) {
	defer func() {
		if recover() != nil {
			v = verdict{panics: true}
		}
	}()
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return verdict{}
	}
	return verdict{ok: true, w: img.Bounds().Dx(), h: img.Bounds().Dy()}
}

func newVerdict(data []byte) (v verdict) {
	defer func() {
		if recover() != nil {
			v = verdict{panics: true}
		}
	}()
	w, h, err := Check(bytes.NewReader(data))
	if err != nil {
		return verdict{}
	}
	return verdict{ok: true, w: w, h: h}
}

func mutate(r *rand.Rand, data []byte, kind int) []byte {
	b := append([]byte(nil), data...)
	switch kind {
	case 0: // 途中で切る
		return b[:r.IntN(len(b))]
	case 1: // ランダムなバイトを書き換える
		for range 1 + r.IntN(8) {
			b[r.IntN(len(b))] = byte(r.IntN(256))
		}
		return b
	case 2: // マーカーを1つ消す（0xFF xx の 2 バイト）
		var idx []int
		for i := 0; i+1 < len(b); i++ {
			if b[i] == 0xff && b[i+1] != 0 && b[i+1] != 0xff {
				idx = append(idx, i)
			}
		}
		if len(idx) == 0 {
			return b
		}
		i := idx[r.IntN(len(idx))]
		return append(b[:i], b[i+2:]...)
	case 3: // 先頭付近（ヘッダ）を書き換える
		n := min(len(b), 700)
		for range 1 + r.IntN(4) {
			b[r.IntN(n)] = byte(r.IntN(256))
		}
		return b
	case 4: // 1 ビット反転
		i := r.IntN(len(b))
		b[i] ^= 1 << r.IntN(8)
		return b
	default: // バイトを挿入する
		i := r.IntN(len(b))
		return append(b[:i], append([]byte{byte(r.IntN(256))}, b[i:]...)...)
	}
}

func TestDiff(t *testing.T) {
	dir := os.Getenv("JPEG_SAMPLE_DIR")
	if dir == "" {
		t.Skip()
	}
	r := rand.New(rand.NewPCG(1, 2))
	var files, mutants, okFiles, okMutants, panics, mismatches int
	dims := map[string]int{}
	filepath.WalkDir(dir, func(path string, e fs.DirEntry, err error) error {
		if err != nil || e.IsDir() {
			return err
		}
		data, _ := os.ReadFile(path)
		files++
		a, b := stdVerdict(data), newVerdict(data)
		if a != b {
			mismatches++
			t.Errorf("original %s: std=%+v new=%+v", path, a, b)
		}
		if a.ok {
			okFiles++
			dims[fmt.Sprintf("%dx%d", a.w, a.h)]++
		}
		if len(data) == 0 {
			return nil
		}
		for k := range 6 {
			m := mutate(r, data, k)
			mutants++
			a, b := stdVerdict(m), newVerdict(m)
			if a.panics {
				panics++
			}
			if a.ok {
				okMutants++
			}
			if a != b {
				mismatches++
				if mismatches < 20 {
					t.Errorf("mutant %d of %s: std=%+v new=%+v", k, path, a, b)
				}
			}
		}
		return nil
	})
	t.Logf("files=%d ok=%d mutants=%d okMutants=%d stdPanics=%d mismatches=%d dims=%v", files, okFiles, mutants, okMutants, panics, mismatches, dims)
}
