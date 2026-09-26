package main

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"sync"
	"time"
)

// 初期データ（mysqldump の 3-initial-data.sql.gz）を MySQL を通さずに読む。
// パースした結果はスナップショットの形でキャッシュし、/api/initialize のたびに復元する。

const (
	initialDataPath   = "../sql/3-initial-data.sql.gz"
	defaultPaymentURL = "http://localhost:12345" // 2-master-data.sql の settings
)

var (
	initialOnce sync.Once
	initialSnap []byte // gob。復元のたびにデコードして新しいオブジェクトを作る
	initialErr  error
)

func loadInitialStore() (*Store, error) {
	initialOnce.Do(func() {
		snap, err := parseInitialData(initialDataPath)
		if err != nil {
			initialErr = err
			return
		}
		var buf bytes.Buffer
		if err := gobEncode(&buf, snap); err != nil {
			initialErr = err
			return
		}
		initialSnap = buf.Bytes()
	})
	if initialErr != nil {
		return nil, initialErr
	}
	snap := &snapshot{}
	if err := gobDecode(bytes.NewReader(initialSnap), snap); err != nil {
		return nil, err
	}
	return restore(snap), nil
}

func parseInitialData(path string) (*snapshot, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(gz)
	if err != nil {
		return nil, err
	}

	snap := &snapshot{PaymentGatewayURL: defaultPaymentURL, ChairLoc: map[string]chairLocAgg{}}
	type loc struct {
		chairID  string
		lat, lon int
		at       time.Time
	}
	var locs []loc

	prefix := []byte("INSERT INTO `")
	for len(data) > 0 {
		i := bytes.Index(data, prefix)
		if i < 0 {
			break
		}
		data = data[i+len(prefix):]
		j := bytes.IndexByte(data, '`')
		table := string(data[:j])
		data = data[j:]
		k := bytes.Index(data, []byte(" VALUES "))
		data = data[k+len(" VALUES "):]

		var rows [][]sqlValue
		rows, data, err = parseValues(data)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", table, err)
		}
		for _, r := range rows {
			switch table {
			case "users":
				snap.Users = append(snap.Users, User{
					ID: r[0].str(), Username: r[1].str(), Firstname: r[2].str(), Lastname: r[3].str(),
					DateOfBirth: r[4].str(), AccessToken: r[5].str(), InvitationCode: r[6].str(),
					CreatedAt: r[7].time(), UpdatedAt: r[8].time(),
				})
			case "owners":
				snap.Owners = append(snap.Owners, Owner{
					ID: r[0].str(), Name: r[1].str(), AccessToken: r[2].str(), ChairRegisterToken: r[3].str(),
					CreatedAt: r[4].time(), UpdatedAt: r[5].time(),
				})
			case "chairs":
				snap.Chairs = append(snap.Chairs, Chair{
					ID: r[0].str(), OwnerID: r[1].str(), Name: r[2].str(), Model: r[3].str(),
					IsActive: r[4].int() != 0, AccessToken: r[5].str(), CreatedAt: r[6].time(), UpdatedAt: r[7].time(),
				})
			case "chair_locations":
				locs = append(locs, loc{chairID: r[1].str(), lat: r[2].int(), lon: r[3].int(), at: r[4].time()})
			case "payment_tokens":
				snap.PaymentTokens = append(snap.PaymentTokens, PaymentToken{UserID: r[0].str(), Token: r[1].str(), CreatedAt: r[2].time()})
			case "rides":
				ride := Ride{
					ID: r[0].str(), UserID: r[1].str(),
					PickupLatitude: r[3].int(), PickupLongitude: r[4].int(),
					DestinationLatitude: r[5].int(), DestinationLongitude: r[6].int(),
					CreatedAt: r[8].time(), UpdatedAt: r[9].time(),
				}
				if !r[2].null {
					ride.ChairID = sql.NullString{String: r[2].str(), Valid: true}
				}
				if !r[7].null {
					v := r[7].int()
					ride.Evaluation = &v
				}
				snap.Rides = append(snap.Rides, ride)
			case "ride_statuses":
				st := RideStatus{ID: r[0].str(), RideID: r[1].str(), Status: r[2].str(), CreatedAt: r[3].time()}
				if !r[4].null {
					t := r[4].time()
					st.AppSentAt = &t
				}
				if !r[5].null {
					t := r[5].time()
					st.ChairSentAt = &t
				}
				snap.Statuses = append(snap.Statuses, st)
			case "coupons":
				c := Coupon{UserID: r[0].str(), Code: r[1].str(), Discount: r[2].int(), CreatedAt: r[3].time()}
				if !r[4].null {
					v := r[4].str()
					c.UsedBy = &v
				}
				snap.Coupons = append(snap.Coupons, c)
			}
		}
	}

	// chair_locations は椅子ごとに created_at 順に畳む（元の LAG ... ORDER BY created_at と同じ）
	sort.SliceStable(locs, func(a, b int) bool { return locs[a].at.Before(locs[b].at) })
	for _, l := range locs {
		agg, ok := snap.ChairLoc[l.chairID]
		if ok {
			agg.TotalDistance += abs(l.lat-agg.Latitude) + abs(l.lon-agg.Longitude)
		}
		agg.Latitude, agg.Longitude, agg.UpdatedAt = l.lat, l.lon, l.at
		snap.ChairLoc[l.chairID] = agg
	}
	return snap, nil
}

type sqlValue struct {
	raw  string
	null bool
}

func (v sqlValue) str() string { return v.raw }
func (v sqlValue) int() int {
	n, err := strconv.Atoi(v.raw)
	if err != nil {
		panic(fmt.Sprintf("initial data: not an int: %q", v.raw))
	}
	return n
}
func (v sqlValue) time() time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05.999999", v.raw, time.UTC)
	if err != nil {
		panic(fmt.Sprintf("initial data: not a datetime: %q", v.raw))
	}
	return t
}

// parseValues は `(...),(...);` を読み、残りを返す。
func parseValues(data []byte) ([][]sqlValue, []byte, error) {
	var rows [][]sqlValue
	i := 0
	for {
		if i >= len(data) || data[i] != '(' {
			return nil, nil, fmt.Errorf("expected '(' at %d", i)
		}
		i++
		var row []sqlValue
		for {
			if data[i] == '\'' {
				i++
				var sb []byte
				for {
					c := data[i]
					if c == '\\' {
						i++
						switch data[i] {
						case 'n':
							sb = append(sb, '\n')
						case 'r':
							sb = append(sb, '\r')
						case 't':
							sb = append(sb, '\t')
						case '0':
							sb = append(sb, 0)
						case 'Z':
							sb = append(sb, 26)
						default:
							sb = append(sb, data[i])
						}
						i++
						continue
					}
					if c == '\'' {
						if i+1 < len(data) && data[i+1] == '\'' {
							sb = append(sb, '\'')
							i += 2
							continue
						}
						i++
						break
					}
					sb = append(sb, c)
					i++
				}
				row = append(row, sqlValue{raw: string(sb)})
			} else {
				j := i
				for data[j] != ',' && data[j] != ')' {
					j++
				}
				tok := string(data[i:j])
				row = append(row, sqlValue{raw: tok, null: tok == "NULL"})
				i = j
			}
			if data[i] == ',' {
				i++
				continue
			}
			// ')'
			i++
			break
		}
		rows = append(rows, row)
		if data[i] == ',' {
			i++
			continue
		}
		if data[i] == ';' {
			return rows, data[i+1:], nil
		}
		return nil, nil, fmt.Errorf("unexpected %q after row", data[i])
	}
}

// ---------------------------------------------------------------------------
// 並べ替え（復元時に元の ORDER BY と同じ順にする）
// ---------------------------------------------------------------------------

func sortByID(cs []Chair) {
	sort.SliceStable(cs, func(a, b int) bool { return cs[a].ID < cs[b].ID })
}

func sortRides(rs []Ride) {
	sort.SliceStable(rs, func(a, b int) bool {
		if !rs[a].CreatedAt.Equal(rs[b].CreatedAt) {
			return rs[a].CreatedAt.Before(rs[b].CreatedAt)
		}
		return rs[a].ID < rs[b].ID
	})
}

func sortStatuses(ss []RideStatus) {
	sort.SliceStable(ss, func(a, b int) bool {
		if !ss[a].CreatedAt.Equal(ss[b].CreatedAt) {
			return ss[a].CreatedAt.Before(ss[b].CreatedAt)
		}
		return ss[a].ID < ss[b].ID
	})
}

func sortCoupons(cs []Coupon) {
	sort.SliceStable(cs, func(a, b int) bool { return cs[a].CreatedAt.Before(cs[b].CreatedAt) })
}

// ---------------------------------------------------------------------------
// マスターデータ（椅子のモデルごとの速度）。固定なのでスナップショットには入れない
// ---------------------------------------------------------------------------

const masterDataPath = "../sql/2-master-data.sql"

var chairSpeeds = map[string]int{}

var chairModelRe = regexp.MustCompile(`\('([^']*)',\s*(\d+)\)`)

func loadChairSpeeds() error {
	b, err := os.ReadFile(masterDataPath)
	if err != nil {
		return err
	}
	i := bytes.Index(b, []byte("INSERT INTO chair_models"))
	if i < 0 {
		return fmt.Errorf("chair_models not found in %s", masterDataPath)
	}
	for _, m := range chairModelRe.FindAllSubmatch(b[i:], -1) {
		speed, err := strconv.Atoi(string(m[2]))
		if err != nil {
			return err
		}
		chairSpeeds[string(m[1])] = speed
	}
	return nil
}

// chairSpeed はモデルの速度。マスターデータに無いモデルは最も遅い 2 とみなす
func chairSpeed(model string) int {
	if s, ok := chairSpeeds[model]; ok {
		return s
	}
	return 2
}
