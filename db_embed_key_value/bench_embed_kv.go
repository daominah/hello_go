package main

import (
	"sync"
	"time"

	"github.com/mywrap/log"
)

type Sequencer interface {
	Reset(key string) error
	Incr(key string) (int64, error)
}

func main() {
	var seqer Sequencer
	//seqer, err := NewLevelDBSequencer("leveldb0.seq", true)
	seqer, err := NewLevelDBSequencer("leveldb1.seq", false)
	//seqer, err := NewBoltDBSequencer("boltdb.seq")
	if err != nil {
		log.Fatal(err)
	}
	key0 := "key0"
	var last int64

	startTime := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Add(-1)
			last, err = seqer.Incr(key0)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
	dur := time.Now().Sub(startTime)
	log.Infof("last: %v, dur: %v", last, dur)
}
