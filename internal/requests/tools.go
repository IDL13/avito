package requests

import (
	"fmt"
	"os"

	"github.com/IDL13/avito/internal/config"
	"github.com/IDL13/avito/pkg/mysql"
)

func (s *server) CreateTables() error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from requests in Insert function: %v\n", err)
		os.Exit(1)
	}

	q1 := `CREATE TABLE Users (Id INT NOT NULL AUTO_INCREMENT, Name VARCHAR(255),CONSTRAINT PK_User_Id PRIMARY KEY (Id));`

	q2 := `CREATE TABLE Dependencies (Id INT NOT NULL AUTO_INCREMENT, UserId INT NOT NULL, Segment VARCHAR(512),
			CONSTRAINT PK_Segment_Id PRIMARY KEY (Id),
			CONSTRAINT FK_Segment_User FOREIGN KEY (UserId) REFERENCES Users (Id) ON DELETE CASCADE ON UPDATE CASCADE);`

	q3 := `CREATE TABLE Segments (Id INT NOT NULL AUTO_INCREMENT, Segment VARCHAR(255),CONSTRAINT PK_Segment_Id PRIMARY KEY (Id));`

	_, err = conn.Exec(q1)
	if err != nil {
		return err
	}

	_, err = conn.Exec(q2)
	if err != nil {
		return err
	}

	_, err = conn.Exec(q3)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) Count() (int, error) {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}

	q := `SELECT COUNT(*) FROM Users`

	var c int
	row := conn.QueryRow(q)
	row.Scan(&c)

	return c, nil
}

func (s *server) RandChoice(counter int, segment string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}

	q1 := `SELECT * FROM Users ORDER BY RAND() LIMIT ?`

	q2 := `INSERT INTO Dependencies (UserId, Segment) VALUES (?, ?)`

	row, err := conn.Query(q1, counter)
	for row.Next() {
		var s Set
		row.Scan(s)
		_, err := conn.Exec(q2, s.Id, segment)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cyclic data append error: %v\n", err)
			os.Exit(1)
		}
	}

	return nil
}
