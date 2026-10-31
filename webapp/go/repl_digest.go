package isureal

// 送り漏れの検出: ISUREAL_REPL_DIGEST=1 の asia1 は、replDigestInterval ごとに s.mu の書きロックの中でストアの要約を作って流す。
// 複製側は同じ位置（それより前の変更を当て終えた状態）で自分の要約を作って比べ、違う区分を repl-digest のログに出す

import (
	"encoding/gob"
	"hash/fnv"
	"log"
	"os"
	"strconv"
	"time"
)

const replDigestInterval = 10 * time.Second

type opDigest struct{ Sums map[string]uint64 }

func (o *opDigest) apply(s *Store) {}

func init() { gob.Register(&opDigest{}) }

func h64(parts ...string) uint64 {
	h := fnv.New64a()
	for _, p := range parts {
		h.Write([]byte(p))
		h.Write([]byte{0})
	}
	return h.Sum64()
}

func tstr(t *time.Time) string {
	if t == nil {
		return "-"
	}
	return t.UTC().Format(time.RFC3339Nano)
}

func pstr(p *string) string {
	if p == nil {
		return "-"
	}
	return *p
}

func itoa(n int) string { return strconv.Itoa(n) }

// digestLocked は読み込み方の違い（削除済みのコメントが投稿の一覧に残るかなど）に左右されない部分だけを要約する
func (s *Store) digestLocked() map[string]uint64 {
	d := map[string]uint64{}
	for _, u := range s.users {
		d["users"] += h64(u.ID, u.Username, u.Status, u.AccountType, u.DisplayName, u.Bio, u.Region, pstr(u.ProfileImageID), tstr(u.ProfileImageCreatedAt))
		d["user_rel"] += h64(u.ID, itoa(len(u.friends)), itoa(len(u.following)), itoa(len(u.followers)), itoa(len(u.blockOut)), itoa(len(u.blockIn)), itoa(len(u.pendingOut)), itoa(len(u.pendingIn)))
		d["user_posts"] += h64(u.ID, itoa(len(u.posts)))
		for h := range u.sessions {
			d["user_sessions"] += h64(u.ID, h)
		}
	}
	d["users_sorted"] = uint64(len(s.usersSorted))
	d["users_by_name"] = uint64(len(s.usersByName))
	for h, r := range s.sessions {
		d["sessions"] += h64(h, r.UserID, strconv.FormatBool(r.Active))
	}
	for _, p := range s.posts {
		live := 0
		for _, c := range p.comments {
			if c.DeletedAt == nil {
				live++
			}
		}
		d["posts"] += h64(p.ID, p.AuthorID, p.WindowID, p.ImageID, p.Caption, p.Visibility, tstr(&p.PostedAt), tstr(p.DeletedAt), p.Location.Kind, pstr(p.Location.Locality))
		d["post_children"] += h64(p.ID, itoa(len(p.realmojis)), itoa(live))
	}
	for _, r := range s.realmojis {
		d["realmojis"] += h64(r.ID, r.PostID, r.ReactorID, r.ImageID, r.ReactionType, tstr(&r.CreatedAt), tstr(r.ReplacedAt), tstr(r.DeletedAt))
	}
	for _, c := range s.comments {
		d["comments"] += h64(c.ID, c.PostID, c.AuthorID, c.Body, tstr(&c.CreatedAt), tstr(c.DeletedAt))
	}
	for _, w := range s.windows {
		d["windows"] += h64(w.ID, w.Region, tstr(&w.StartsAt), tstr(w.EndedAt), itoa(len(w.posts)))
		for uid, n := range w.notified {
			d["notified"] += h64(w.ID, uid, n.ID)
		}
	}
	for r, w := range s.activeWindows {
		d["active_windows"] += h64(r, w.ID)
	}
	for _, r := range s.friendRequests {
		d["friend_requests"] += h64(r.ID, r.RequesterID, r.RecipientID, r.Status)
	}
	for _, f := range s.friendships {
		d["friendships"] += h64(f.ID, f.LowID, f.HighID)
	}
	for _, f := range s.follows {
		d["follows"] += h64(f.ID, f.FollowerID, f.FolloweeID)
	}
	for _, b := range s.blocks {
		d["blocks"] += h64(b.ID, b.BlockerID, b.BlockedID)
	}
	d["profile_by_image"] = uint64(len(s.profileByImage))
	d["post_by_image"] = uint64(len(s.postByImage))
	d["rm_by_image"] = uint64(len(s.rmByImage))
	return d
}

// StartReplDigest は asia1 で呼ぶ
func StartReplDigest(s *Store) {
	if os.Getenv("ISUREAL_REPL_DIGEST") != "1" {
		return
	}
	go func() {
		for range time.Tick(replDigestInterval) {
			s.mu.Lock()
			start := time.Now()
			op := &opDigest{Sums: s.digestLocked()}
			repl.publish(op, replAckTimeout)
			s.mu.Unlock()
			log.Printf("repl-digest: built in %s", time.Since(start))
		}
	}()
}

func (r *replica) checkDigestLocked(seq uint64, want map[string]uint64) {
	got := r.s.digestLocked()
	var bad []string
	for k, v := range want {
		if got[k] != v {
			bad = append(bad, k)
		}
	}
	if len(bad) == 0 {
		log.Printf("repl-digest: seq %d ok", seq)
		return
	}
	log.Printf("repl-digest: seq %d MISMATCH %v", seq, bad)
}
