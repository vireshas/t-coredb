package db

import (
	"sync"
)

//let's protect concurrent access to the map
var rwMutex sync.RWMutex
var poolMap = make(map[string]interface{})

type dbConfig map[string]string

type createPoolCallBack func(configs dbConfig) interface{}
type PoolManager struct{}

func createUniqKey(configs dbConfig) (key string) {
	for _, value := range configs {
		key += value
	}
	return
}

func (m PoolManager) GetConnection(cb createPoolCallBack, configs dbConfig) interface{} {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	key := createUniqKey(configs)
	pool, ok := poolMap[key]
	if !ok {
		pool = cb(configs)
		poolMap[key] = pool
	}
	return pool
}
