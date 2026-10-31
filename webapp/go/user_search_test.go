package isureal

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

// UserRow をたどって走査する実装。検索用の配列を使う実装と結果を比べる
func (s *Store) searchUsersScanLocked(viewer, query, beforeUsername, beforeID string) []User {
	list := make([]*UserRow, 0, len(s.users))
	for _, u := range s.users {
		list = append(list, u)
	}
	sort.Slice(list, func(i, j int) bool { return userLess(list[i], list[j]) })
	users := make([]User, 0, 51)
	v := s.users[viewer]
	for _, u := range list {
		if beforeUsername != "" && !(u.Username > beforeUsername || (u.Username == beforeUsername && u.ID > beforeID)) {
			continue
		}
		if !u.active() || u.ID == viewer {
			continue
		}
		if !strings.Contains(u.normUsername, query) && !strings.Contains(u.normDisplayName, query) {
			continue
		}
		if v != nil && (v.blockOut[u.ID] != nil || v.blockIn[u.ID] != nil) {
			continue
		}
		users = append(users, u.User)
		if len(users) == 51 {
			break
		}
	}
	return users
}

var searchTestWords = []string{"ab", "Ab", "ＡＢ", "straße", "STRASSE", "x", "café", "café", "ba", "a"}

func (f *feedFixture) searchWord() string {
	return searchTestWords[f.rng.Intn(len(searchTestWords))] + searchTestWords[f.rng.Intn(len(searchTestWords))]
}

func TestSearchUsersMatchesScan(t *testing.T) {
	for seed := int64(1); seed <= 5; seed++ {
		f := newFeedFixture(seed, 400)
		s := f.s
		for _, u := range s.users {
			u.Username = fmt.Sprintf("%c%d", 'a'+rune(f.rng.Intn(3)), f.rng.Intn(50))
			u.DisplayName = f.searchWord()
			u.setNormalized()
		}
		s.rebuildIndexes()
		check := func(step int) {
			for k := 0; k < 30; k++ {
				viewer := f.user()
				query := normalizeSearchText([]string{searchTestWords[f.rng.Intn(len(searchTestWords))], "a1", "b", "c2"}[f.rng.Intn(4)])
				var beforeUsername, beforeID string
				if f.rng.Intn(2) == 0 {
					u := s.users[f.user()]
					beforeUsername, beforeID = u.Username, u.ID
				}
				got := s.searchUsersLocked(viewer, query, beforeUsername, beforeID)
				want := s.searchUsersScanLocked(viewer, query, beforeUsername, beforeID)
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("seed %d step %d query %q: got %d users, want %d", seed, step, query, len(got), len(want))
				}
			}
		}
		check(0)
		for step := 1; step <= 50; step++ {
			switch f.rng.Intn(3) {
			case 0:
				u := &UserRow{User: User{ID: f.id("u"), Username: fmt.Sprintf("%c%d", 'a'+rune(f.rng.Intn(3)), f.rng.Intn(50)), DisplayName: f.searchWord(), Status: "active"}}
				u.setNormalized()
				s.addUserLocked(u)
				f.ids = append(f.ids, u.ID)
			case 1:
				u := s.users[f.user()]
				u.DisplayName = f.searchWord()
				u.normDisplayName = normalizeSearchText(u.DisplayName)
				s.refreshUserSearchLocked(u)
			case 2:
				u := s.users[f.user()]
				u.Status = "deleted"
				s.refreshUserSearchLocked(u)
			}
			check(step)
		}
	}
}
