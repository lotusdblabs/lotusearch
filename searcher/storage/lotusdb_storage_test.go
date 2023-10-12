package storage

import (
	"fmt"
	"testing"
)

func TestNewLotusDbStorage(t *testing.T) {
	path := "./lotusdb_basic.db"
	lotusDb, err := NewLotusDbStorage(path)
	if err != nil {
		fmt.Println(err)
	}
	key := "test"
	value := "lotusdb"
	err = lotusDb.Set([]byte(key), []byte(value))
	if err != nil {
		fmt.Println(err)
	}
	exist := lotusDb.Exist([]byte(key))
	if exist {
		v, err := lotusDb.Get([]byte(key))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(v)
	}
}

func TestLotusDbStorage_Get(t *testing.T) {
	path := "./lotusdb_basic.db"
	lotusDb, err := NewLotusDbStorage(path)
	if err != nil {
		fmt.Println(err)
	}
	key := "test"
	exist := lotusDb.Exist([]byte(key))
	if exist {
		v, err := lotusDb.Get([]byte(key))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(v))

		err = lotusDb.Delete([]byte(key))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("delete successfully")

		exist2 := lotusDb.Exist([]byte(key))
		fmt.Println(exist2)
	}
}
