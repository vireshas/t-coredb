package db

import (
	"github.com/goibibo/mantle"
	"github.com/goibibo/t-settings"
	"strconv"
)

func GetRedisClientFor(vertical string) mantle.Mantle {
	configs := settings.GetConfigsFor("redis", vertical)
	pool := PoolManager{}.GetConnection(createRedisPool, configs)
	return pool.(*mantle.Orm).New()
}

func foundOrDefault(configs dbConfig, key string, fallback string) string {
	value, ok := configs[key]
	if !ok {
		value = fallback
	}
	return value
}

func createRedisPool(configs dbConfig) interface{} {
	connectionUrl := settings.ConstructRedisPath(configs)
	db := foundOrDefault(configs, "db", "0")
	capacity, _ := strconv.Atoi(foundOrDefault(configs, "pool_size", "10"))
	options := map[string]string{"db": db}
	pool := mantle.Orm{
		Driver:       "redis",
		HostAndPorts: []string{connectionUrl},
		Capacity:     capacity,
		Options:      options}
	return &pool
}
