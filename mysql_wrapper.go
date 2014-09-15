package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlClient(hostPort string) *sql.DB {
	pool := PoolManager{}.GetConnection(createMysqlPool, hostPort)
	return pool.(*sql.DB)
}

func createMysqlPool(hostNPorts []string) interface{} {
	db, err := sql.Open("mysql", hostNPorts[0])
	if err != nil {
		return nil
	}
	return db
}
