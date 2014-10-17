package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goibibo/t-settings"
)

func GetMysqlClientFor(vertical string) *sql.DB {
	configs := settings.GetConfigsFor("mysql", vertical)
	options := make(map[string]string)
	connectionUrl := settings.ConstructMysqlPath(configs)
	pool := PoolManager{}.GetConnection(createMysqlPool, connectionUrl, options)
	return pool.(*sql.DB)
}

func createMysqlPool(hostNPorts []string, options map[string]string) interface{} {
	db, err := sql.Open("mysql", hostNPorts[0])
	if err != nil {
		return nil
	}
	return db
}
