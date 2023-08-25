package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewClient(password, host, port, db string) (conn *sql.DB, err error) {
	q := fmt.Sprintf("root:%s:@tcp(%s:%s)/%s", password, host, port, db)
	conn, err = sql.Open("mysql", q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	return conn, nil
}
