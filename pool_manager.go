package db

import (
	"sync"
)

//let's protect concurrent access to the map
var rwMutex sync.RWMutex
var IpToPoolMap = make(map[string]interface{})

type createPoolCallBack func(hostNPort []string, options map[string]string) interface{}
type PoolManager struct{}

func (m PoolManager) GetConnection(cb createPoolCallBack, hostNPort string, options map[string]string) interface{} {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	pool, ok := IpToPoolMap[hostNPort]
	if !ok {
		hostAndPort := []string{hostNPort}
		pool = cb(hostAndPort, options)
		IpToPoolMap[hostNPort] = pool
	}
	return pool
}
