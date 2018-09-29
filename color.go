package main

import (
	"hash/crc32"
	"math/rand"
)

func color(domain string) string {
	rnd := rand.New(rand.NewSource(int64(crc32.ChecksumIEEE([]byte(domain)))))
	return config.Colors[rnd.Intn(len(config.Colors))]
}
