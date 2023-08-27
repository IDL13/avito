package request

import (
	"errors"
	"fmt"
	"os"

	"github.com/IDL13/avito/internal/config"
	"github.com/IDL13/avito/pkg/mysql"
)

type Set struct {
	id      int
	segment string
}

func CreateTables() error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from requests in Insert function: %v\n", err)
		os.Exit(1)
	}
	q1 := `CREATE TABLE Users (Id INT NOT NULL AUTO_INCREMENT, Name VARCHAR(255),CONSTRAINT PK_User_Id PRIMARY KEY (Id));`
	q2 := `CREATE TABLE Dependencies (Id INT NOT NULL AUTO_INCREMENT, UserId INT NOT NULL, Segment VARCHAR(512),CONSTRAINT PK_Segment_Id PRIMARY KEY (Id),CONSTRAINT FK_Segment_User FOREIGN KEY (UserId) REFERENCES Users (Id) ON DELETE CASCADE ON UPDATE CASCADE);`
	q3 := `CREATE TABLE Segments (Id INT NOT NULL AUTO_INCREMENT, Segment VARCHAR(255),CONSTRAINT PK_Segment_Id PRIMARY KEY (Id));`
	_, err = conn.Exec(q1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from CrateDatabase: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(q2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from CrateDatabase: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(q3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from CrateDatabase: %v\n", err)
		os.Exit(1)
	}
	return nil
}

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

func InserSegment(segment string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	valid := `SELECT Segment FROM Segments WHERE Segment = ?`
	row, err := conn.Query(valid, segment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while adding data: %v\n", err)
		os.Exit(1)
	}
	if row.Next() {
		err = errors.New("Duplicate")
		return err
	}
	q := `INSERT INTO Segments (Segment) VALUES (?)`
	_, err = conn.Exec(q, segment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while adding data: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func DeleteSegment(segment string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	q := `DELETE FROM Segments WHERE Segment = ?`
	_, err = conn.Exec(q, segment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while adding data: %v\n", err)
		os.Exit(1)
	}
	return nil
}

func SearchSegmentsForUser() (map[int][]string, error) {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	q := `SELECT Users.Id, Dependencies.Segment FROM Users RIGHT JOIN Dependencies ON Users.Id = Dependencies.UserId;`
	row, err := conn.Query(q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while adding data: %v\n", err)
		os.Exit(1)
	}
	m := make(map[int][]string)
	for row.Next() {
		var S Set
		err = row.Scan(&S.id, &S.segment)
		if err != nil {
			panic(err)
		}
		_, ok := m[S.id]
		if ok {
			m[S.id] = append(m[S.id], S.segment)
		} else {
			m[S.id] = append(m[S.id], S.segment)
		}
	}
	return m, nil
}

func InsertDependencies(UserId int, Segments []string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	q := `INSERT INTO Dependencies (UserId, Segment) VALUES (?, ?)`
	for iter := range Segments {
		_, err := conn.Exec(q, UserId, Segments[iter])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cyclic data append error: %v\n", err)
			os.Exit(1)
		}
	}
	return nil
}

func DeleteDependencies(UserId int, Segments []string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	q := `DELETE FROM Dependencies WHERE UserId = ? AND Segment = ?`
	for iter := range Segments {
		_, err := conn.Exec(q, UserId, Segments[iter])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cyclic data remove error: %v\n", err)
			os.Exit(1)
		}
	}
	return nil
}
