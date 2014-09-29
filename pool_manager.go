package db

import (
	"sync"
)

//let's protect concurrent access to the map
var rwMutex sync.RWMutex
var IpToPoolMap = make(map[string]interface{})

type createPoolCallBack func(hostNPort []string) interface{}
type PoolManager struct{}

func (m PoolManager) GetConnection(cb createPoolCallBack, hostNPort string) interface{} {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	pool, ok := IpToPoolMap[hostNPort]
	if !ok {
		hostAndPort := []string{hostNPort}
		pool = cb(hostAndPort)
		IpToPoolMap[hostNPort] = pool
	}
	return pool
}
