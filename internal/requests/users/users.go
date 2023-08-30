package users

import (
	"fmt"
	"os"

	"github.com/IDL13/avito/internal/config"
	"github.com/IDL13/avito/pkg/mysql"
)

func InsertUser(name string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	q := `INSERT INTO Users (Name) VALUES (?)`
	_, err = conn.Exec(q, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while adding data: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func DeleteUser(name string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	q := `DELETE FROM Users WHERE Name = ?`
	_, err = conn.Exec(q, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while adding data: %v\n", err)
		os.Exit(1)
	}
	return nil
}
