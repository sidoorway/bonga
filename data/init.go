package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *Storage) init() {
	s.log = logrus.WithField("mod", "data.Storage")
	log := s.log.WithField("func", "init")

	log.Info("Init bonga storage")

	s.storage = make([]Room, 0)
	s.gender = make(map[string][]*Room)
	s.usernames = make(map[string]int)

	data, err := ioutil.ReadFile("bonga.json")
	if err != nil {
		logrus.Warn(err)
		s.Reload()
		return
	}

	err = json.Unmarshal(data, &s.storage)
	if err != nil {
		logrus.Warn(err)
		return
	}

	for i := range s.storage {
		s.usernames[s.storage[i].UserName] = i
	}

	s.gender = s.parseGenders()

	log.Info("Found ", len(s.storage), " rooms")
	log.Info("Found ", len(s.gender), " genders")
}

func (s *Storage) load() (storage []Room, err error) {

	url := "https://promo-bc.com/promo.php?c=16778&lang=ru&type=api&api_v=1&api_type=json"

	client := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logrus.Warn(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		logrus.Warn(err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Warn(err)
		return
	}

	err = json.Unmarshal(data, &storage)
	if err != nil {
		logrus.Warn(err)
		return
	}

	s.usernames = make(map[string]int, 0)

	for i := range s.storage {
		s.usernames[s.storage[i].UserName] = i
	}

	return
}

func (s *Storage) parseGenders() (genders map[string][]*Room) {

	genders = make(map[string][]*Room)

	for i, room := range s.storage {
		_, ok := genders[room.Gender]
		if !ok {
			genders[room.Gender] = make([]*Room, 0)
		}
		genders[room.Gender] = append(genders[room.Gender], &s.storage[i])
	}

	return
}

func (s *Storage) Reload() {

	log := s.log.WithField("func", "reload")
	log.Info("Load rooms")

	storage, err := s.load()
	if err != nil {
		log.Warn(err)
		return
	}

	s.Access.Lock()

	for i := range storage {
		if _, ok := s.usernames[storage[i].UserName]; !ok {
			s.storage = append(s.storage, storage[i])
			s.usernames[storage[i].UserName] = len(s.storage) - 1
			continue
		}
		s.storage[s.usernames[storage[i].UserName]] = storage[i]
	}

	s.Access.Unlock()

	j, err := json.Marshal(s.storage)
	if err != nil {
		log.Warn(err)
	}

	err = ioutil.WriteFile("bonga.json", j, 0755)
	if err != nil {
		log.Warn(err)
	}

	s.gender = s.parseGenders()

	log.Info("Found ", len(s.storage), " rooms")
	log.Info("Found ", len(s.gender), " genders")
}
