package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// MESSAGE from env
		fmt.Fprintf(w, "Env MESSAGE: %s\n", os.Getenv("MESSAGE"))

		// MESSAGE from MySQL
		dbUser := os.Getenv("MYSQL_USER")
		dbPass := os.Getenv("MYSQL_PASSWORD")
		dbHost := os.Getenv("MYSQL_HOST") // set to 'mysql' in deployment
		dbName := os.Getenv("MYSQL_DATABASE")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Fprintf(w, "DB error: %v", err)
			return
		}
		defer db.Close()

		var message string
		appName := os.Getenv("APP_NAME")
		err = db.QueryRow("SELECT message FROM messages WHERE app_name=?", appName).Scan(&message)
		if err != nil {
			fmt.Fprintf(w, "Query error: %v", err)
			return
		}
		fmt.Fprintf(w, "MySQL MESSAGE: %s", message)
	})

	port := fmt.Sprintf(":%s", os.Getenv("LISTENANDSERVE"))
	http.ListenAndServe(port, nil)
}
