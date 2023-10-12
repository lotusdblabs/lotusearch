package storage

import (
	"log"
	"sync"
	"time"

	"github.com/lotusdblabs/lotusdb/v2"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

// LotusDbStorage 基于lotus db的存储 // TODO maybe we can replace leveldb by lotusdb
type LotusDbStorage struct {
	db       *lotusdb.DB
	path     string
	mu       sync.RWMutex // 加锁
	closed   bool
	lastTime int64
}

// NewLotusDbStorage 打开 LotusDb 数据库
func NewLotusDbStorage(path string) (*LotusDbStorage, error) {
	db := &LotusDbStorage{
		path:     path,
		closed:   true,
		lastTime: time.Now().Unix(),
	}

	return db, nil
}

func (s *LotusDbStorage) autoOpenDB() {
	if s.isClosed() {
		s.ReOpen()
	}
	s.lastTime = time.Now().Unix()
}

func (s *LotusDbStorage) openDB() (*lotusdb.DB, error) {
	if s.path == "" {
		return nil, errors.New("path is nil,please set the path first")
	}
	db, err := lotusdb.Open(lotusdb.Options{
		DirPath: s.path,
	})
	return db, err
}

func (s *LotusDbStorage) ReOpen() {
	if !s.isClosed() {
		log.Println("db is not closed")
		return
	}
	s.mu.Lock()
	db, err := s.openDB()
	if err != nil {
		panic(err)
	}
	s.db = db
	s.closed = false
	s.mu.Unlock()
}

func (s *LotusDbStorage) isClosed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.closed
}

// Exist determined the key exist or not
func (s *LotusDbStorage) Exist(key []byte) bool {
	s.autoOpenDB()
	has, err := s.db.Exist(key)
	if err != nil {
		panic(err)
	}
	return has
}

// Set store the key and value
func (s *LotusDbStorage) Set(key []byte, value []byte) (err error) {
	s.autoOpenDB()
	err = s.db.Put(key, value, nil)
	if err != nil {
		panic(err)
	}

	return
}

// Get the value by get key
func (s *LotusDbStorage) Get(key []byte) ([]byte, error) {
	s.autoOpenDB()
	buffer, err := s.db.Get(key)
	if err != nil {
		return nil, err
	}

	return buffer, err
}

// Delete the key
func (s *LotusDbStorage) Delete(key []byte) (err error) {
	s.autoOpenDB()
	return s.db.Delete(key, nil)
}

// Close 关闭
func (s *LotusDbStorage) Close() (err error) {
	if s.isClosed() {
		return nil
	}
	s.mu.Lock()
	err = s.db.Close()
	if err != nil {
		return
	}

	s.closed = true
	s.mu.Unlock()

	return
}
