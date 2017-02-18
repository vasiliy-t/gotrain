package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/ssl_tmp")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	ticker := time.NewTicker(time.Duration(1 * time.Second))

	for range ticker.C {
		fmt.Printf("ping: %s", db.Ping())
		rows, err := db.Query("SELECT 1")
		if err == nil {
			cols, _ := rows.Columns()
			fmt.Printf("%+v %s\n", cols, err)
		} else {
			fmt.Println("failed")
		}
	}
}
