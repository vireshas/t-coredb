package db

import (
	"sync"
)

//let's protect concurrent access to the poolMap
var rwMutex sync.RWMutex

// unique key to pool mapping
var poolMap = make(map[string]interface{})

//this type is a syntactic sugar
type dbConfig map[string]string

//every db should implement a callback which will be called by pool_manager
type createPoolCallBack func(configs dbConfig) interface{}

//creates a unique key for every config
func createUniqKey(configs dbConfig) (key string) {
	for _, value := range configs {
		key += value
	}
	return
}

//get or set a connection from poolMap
func GetConnection(cb createPoolCallBack, configs dbConfig) interface{} {
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
