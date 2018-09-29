package data

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Storage struct {
	storage   []Room
	gender    map[string][]*Room
	usernames map[string]int

	log    *logrus.Entry
	Access sync.Mutex
}

var (
	instance *Storage
	once     sync.Once
	err      error
)

func Instance() *Storage {
	once.Do(func() {
		instance = &Storage{}
		instance.init()
	})
	return instance
}
