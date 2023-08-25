package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/IDL13/avito/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewClient(config config.DbConfig) (conn *sql.DB, err error) {
	q := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	conn, err = sql.Open("mysql", q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databases: %v\n", err)
		os.Exit(1)
	}
	return conn, nil
}
