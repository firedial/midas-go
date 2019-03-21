package db

import (
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
    db, err := sql.Open("mysql", "root:@/midas?parseTime=true")
    if err != nil {
      panic(err.Error())
    }

    return db
}
