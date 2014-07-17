package db

import (
	"github.com/vireshas/mantle"
)

func GetRedisClient(hostPort string) mantle.Mantle {
	pool := PoolManager{}.GetConnection(createRedisPool, hostPort)
	return pool.(*mantle.Orm).Get()
}

func createRedisPool(hostNPorts []string) interface{} {
	pool := mantle.Orm{Driver: "redis", HostAndPorts: hostNPorts, Capacity: 100}
	return &pool
}
