package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goibibo/t-settings"
)

func GetMysqlClientFor(vertical string) *sql.DB {
	configs := settings.GetConfigsFor("mysql", vertical)
	pool := PoolManager{}.GetConnection(createMysqlPool, configs)
	return pool.(*sql.DB)
}

func createMysqlPool(configs dbConfig) interface{} {
	connectionUrl := settings.ConstructMysqlPath(configs)
	db, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		return nil
	}
	return db
}
