package db

import (
	"github.com/goibibo/mantle"
	"github.com/goibibo/t-settings"
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
