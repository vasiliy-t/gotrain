package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		mysqlHost = os.Getenv("MYSQL_HOST")
		mysqlPort = os.Getenv("MYSQL_PORT")
		mysqlUser = os.Getenv("MYSQL_USER")
		mysqlPass = os.Getenv("MYSQL_PASSWORD")
		mysqlDB   = os.Getenv("MYSQL_DATABASE_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("mysql: failed to open")
	}

	defer db.Close()

	maxRetries := 30
	for {
		err = db.Ping()
		if err == nil {
			log.Printf("Connected to DB")
			break
		}
		log.Printf("Error: %s", err)
		if maxRetries == 0 {
			log.Fatalf("mysql: failed to connect")
		}
		maxRetries--
		<-time.After(time.Duration(1 * time.Second))
	}

	e, err := GetDemo(context.Background(), db)
	if err != nil {
		log.Fatalf("GetDemo: %s", err)
	}

	log.Printf("%+v", e)
}

type Querier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type entity struct {
}

func GetDemo(ctx context.Context, db Querier) (entity, error) {
	expr := "SELECT `id`, `title` FROM `demo` LIMIT 1"
	row := db.QueryRowContext(ctx, expr)
	log.Printf("%+v", row)
	return entity{}, nil
}
