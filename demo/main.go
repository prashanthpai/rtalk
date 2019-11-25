package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
)

const (
	KB          = 1 * 1024
	maxCost     = 1 * KB * KB // 1 MB
	numCounters = 10 * maxCost
)

func main() {

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: numCounters,
		MaxCost:     maxCost,
		BufferItems: 64,
		Metrics:     true,
	})
	if err != nil {
		log.Fatalf("ristretto.NewCache() failed: %s", err.Error())
	}
	defer cache.Close()

	// Store 4 keys, each of size 256KB
	for i := 0; i < 4; i++ {
		b, err := getRandom(256 * KB)
		if err != nil {
			log.Fatalf("getRandom() failed: %s", err.Error())
		}
		cache.Set(i, b, int64(len(b)))
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("%s", cache.Metrics.String())
}

func getRandom(size int) ([]byte, error) {

	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
