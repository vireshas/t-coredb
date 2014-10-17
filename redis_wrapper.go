package db

import (
	"github.com/goibibo/mantle"
	"github.com/goibibo/t-settings"
)

func GetRedisClientFor(vertical string) mantle.Mantle {
	configs := settings.GetConfigsFor("redis", vertical)
	db, ok := configs["db"]
	if !ok {
		db = "0"
	}
	options := map[string]string{"db": db}
	connectionUrl := settings.ConstructRedisPath(configs)
	pool := PoolManager{}.GetConnection(createRedisPool, connectionUrl, options)
	return pool.(*mantle.Orm).New()
}

func createRedisPool(hostNPorts []string, params map[string]string) interface{} {
	pool := mantle.Orm{Driver: "redis", HostAndPorts: hostNPorts, Capacity: 100, Options: params}
	return &pool
}
