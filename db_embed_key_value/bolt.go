package main

import (
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

type BoltDBSequencer struct {
	path   string
	db     *bolt.DB
	bucket []byte
}

func NewBoltDBSequencer(path string) (*BoltDBSequencer, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	bkt := []byte("sequences")
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bkt)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &BoltDBSequencer{path: path, db: db, bucket: bkt}, nil
}

func (s BoltDBSequencer) Reset(key string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(s.bucket)
		return bkt.Put([]byte(key), []byte(strconv.FormatInt(-1, 10)))
	})
}

func (s BoltDBSequencer) Incr(key string) (int64, error) {
	var newSeq int64
	err := s.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(s.bucket)
		oldB := bkt.Get([]byte(key))
		old, err := strconv.ParseInt(string(oldB), 10, 64)
		if err != nil {
			old = -1
		}
		newSeq = old + 1
		return bkt.Put([]byte(key), []byte(strconv.FormatInt(newSeq, 10)))
	})
	return newSeq, err
}
