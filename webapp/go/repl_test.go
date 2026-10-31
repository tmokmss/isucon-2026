package isureal

import (
	"bytes"
	"encoding/gob"
	"testing"
	"time"
)

func TestReplOpsRoundTrip(t *testing.T) {
	at := now()
	empty := ""
	zero := 0.0
	ops := []replOp{
		&opUserAdd{U: &UserRow{User: User{ID: "u1", Username: "a", ImageKey: []byte{1}}, CreatedAt: at}},
		&opLogin{UserID: "u1", Row: &LoginSessionRow{ID: "s", UserID: "u1", TokenHash: "h", Active: true, CreatedAt: at}},
		&opUpdateMe{UserID: "u1", Bio: opt(&empty)},
		&opWindowOpen{W: &WindowRow{PostingWindow: PostingWindow{ID: "w", Region: "asia", StartsAt: at}}, Notified: map[string]notificationRow{"u1": {ID: "n", CreatedAt: at}}},
		newOpPostAdd(&PostRow{ID: "p", AuthorID: "u1", WindowID: "w", Location: PostLocation{Kind: "k", Latitude: &zero, Locality: &empty}, PostedAt: at}),
		&opPostUpdate{ID: "p", Caption: opt(&empty)},
		&opReset{}, &opHeartbeat{}, &opLive{},
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	for i, op := range ops {
		if err := enc.Encode(&replMsg{Seq: uint64(i + 1), Op: op}); err != nil {
			t.Fatalf("encode %T: %v", op, err)
		}
	}
	dec := gob.NewDecoder(&buf)
	var got []replOp
	for range ops {
		var m replMsg
		if err := dec.Decode(&m); err != nil {
			t.Fatal(err)
		}
		got = append(got, m.Op)
	}
	if u := got[2].(*opUpdateMe); !u.Bio.Set || u.Bio.V != "" || u.DisplayName.Set {
		t.Fatalf("updateMe %+v", u)
	}
	s := NewStore()
	s.users = map[string]*UserRow{}
	s.usersByName = map[string]*UserRow{}
	s.sessions = map[string]*LoginSessionRow{}
	s.windows = map[string]*WindowRow{}
	s.activeWindows = map[string]*WindowRow{}
	s.maxLocalDate = map[string]time.Time{}
	s.posts = map[string]*PostRow{}
	s.postByImage = map[string]*PostRow{}
	s.profileByImage = map[string]*UserRow{}
	for _, op := range got {
		op.apply(s)
	}
	p := s.posts["p"]
	if p == nil || p.Location.Latitude == nil || p.Location.Locality == nil || p.Location.Longitude != nil || !p.PostedAt.Equal(at) {
		t.Fatalf("post %+v", p)
	}
	if s.sessions["h"] == nil || len(s.windows["w"].notified) != 1 || s.activeWindows["asia"] == nil {
		t.Fatal("state")
	}
}
