coredb
======

Library which helps in reusing connection pools.

DB schema: https://gist.github.com/vireshas/5f6cc3662ba3e0d95d47  
Redis value: https://gist.github.com/vireshas/a194abcd8cfbbb70fde5

    package main
    
    import (
            "fmt"
            "github.com/vireshas/t-coredb"
            "github.com/vireshas/t-settings"
    )
    
    func main() {
            settings.Configure()
            connection := db.GetRedisClientFor("r1")
            value := connection.Get("key3")
            fmt.Println(value)
    
            mysqldb := db.GetMysqlClientFor("m1")
            for i := 0; i < 10; i++ {
                    var msg string
                    err := mysqldb.QueryRow("SELECT value FROM bm WHERE id=?", i).Scan(&msg)
                    if err != nil {
                            fmt.Println("Database Error!")
                    } else {
                            fmt.Println(msg)
                    }
            }
    }
