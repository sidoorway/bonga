package data

import (
	"fmt"
	"math/rand"
)

func (s *Storage) Genders() (genders []string) {
	genders = make([]string, 0)
	for gender := range s.gender {
		genders = append(genders, gender)
	}
	return
}

func (s *Storage) Room(username string) (room Room, err error) {
	i, ok := s.usernames[username]
	if !ok {
		err = fmt.Errorf("Unknown user: %s", username)
		return
	}
	return s.storage[i], nil
}

func (s *Storage) Rooms(gender string, count int, seed int64) (rooms []Room) {

	rooms = make([]Room, 0)

	_, ok := s.gender[gender]
	if !ok {
		return
	}

	rnd := rand.New(rand.NewSource(seed))
	roomIDs := rnd.Perm(len(s.gender[gender]))

	for _, ID := range roomIDs {
		if s.gender[gender][ID].Status {
			rooms = append(rooms, *s.gender[gender][ID])
			if len(rooms) == count {
				break
			}
		}
	}
	return
}

func (s *Storage) Check(username string) bool {
	_, ok := s.usernames[username]
	return ok
}
