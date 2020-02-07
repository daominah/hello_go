package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBSequencer struct {
	path  string
	db    *leveldb.DB
	mutex sync.Mutex
}

func NewLevelDBSequencer(path string) (*LevelDBSequencer, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBSequencer{path: path, db: db}, nil
}

func (s *LevelDBSequencer) Reset(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	err := s.db.Put([]byte(key), []byte(strconv.FormatInt(-1, 10)), nil)
	return err
}

func (s *LevelDBSequencer) Incr(key string) (int64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	oldB, err := s.db.Get([]byte(key), nil)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return 0, err
	}
	old, err := strconv.ParseInt(string(oldB), 10, 64)
	if err != nil {
		old = -1
	}
	newSeq := old + 1
	return newSeq, s.db.Put([]byte(key), []byte(strconv.FormatInt(newSeq, 10)), nil)
}
