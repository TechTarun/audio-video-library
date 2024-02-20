package db

import(
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	var conn *sql.DB
	connStr := "user=<username> password=<password> dbname=recordings"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := conn.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Welcome, you are into the DB!")
	return conn
}