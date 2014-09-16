package db

import (
	"github.com/vireshas/mantle"
	"github.com/vireshas/t-settings"
)

func GetRedisClientFor(vertical string) mantle.Mantle {
	configs := settings.GetConfigsFor("redis", vertical)
	connectionUrl := settings.ConstructRedisPath(configs)
	pool := PoolManager{}.GetConnection(createRedisPool, connectionUrl)
	return pool.(*mantle.Orm).Get()
}

func createRedisPool(hostNPorts []string) interface{} {
	pool := mantle.Orm{Driver: "redis", HostAndPorts: hostNPorts, Capacity: 100}
	return &pool
}
