package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vireshas/t-settings"
)

func GetMysqlClientFor(vertical string) *sql.DB {
	configs := settings.GetConfigsFor("mysql", vertical)
	connectionUrl := settings.ConstructMysqlPath(configs)
	pool := PoolManager{}.GetConnection(createMysqlPool, connectionUrl)
	return pool.(*sql.DB)
}

func createMysqlPool(hostNPorts []string) interface{} {
	db, err := sql.Open("mysql", hostNPorts[0])
	if err != nil {
		return nil
	}
	return db
}
