package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDBSequencer struct {
	path        string
	isWriteSync bool
	optionWrite opt.WriteOptions
	db          *leveldb.DB
	mutex       sync.Mutex
}

func NewLevelDBSequencer(path string, isWriteSync bool) (*LevelDBSequencer, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	var optionWrite opt.WriteOptions
	if isWriteSync {
		optionWrite = opt.WriteOptions{Sync: true}
	}
	return &LevelDBSequencer{
		path: path, isWriteSync: isWriteSync,
		db: db, optionWrite: optionWrite}, nil
}

func (s *LevelDBSequencer) Reset(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	err := s.db.Put([]byte(key), []byte(strconv.FormatInt(-1, 10)), &s.optionWrite)
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
	return newSeq, s.db.Put([]byte(key), []byte(strconv.FormatInt(newSeq, 10)), &s.optionWrite)
}
