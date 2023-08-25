package request

import (
	"fmt"
	"os"

	"github.com/IDL13/avito/internal/config"
	"github.com/IDL13/avito/pkg/mysql"
)

func CreateTables() error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from requests in Insert function: %v\n", err)
		os.Exit(1)
	}
	q1 := `CREATE TABLE Users (Id INT NOT NULL, Name VARCHAR(255),CONSTRAINT PK_User_Id PRIMARY KEY (Id));`
	q2 := `CREATE TABLE Dependencies (Id INT NOT NULL, UserId INT NOT NULL, Segment VARCHAR(512),CONSTRAINT PK_Segment_Id PRIMARY KEY (Id),CONSTRAINT FK_Segment_User FOREIGN KEY (UserId) REFERENCES Users (Id) ON DELETE CASCADE ON UPDATE CASCADE);`
	q3 := `CREATE TABLE Segments (Id INT NOT NULL, Segment VARCHAR(255),CONSTRAINT PK_Segment_Id PRIMARY KEY (Id));`
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

func InsertUser(id int, name string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	q := `INSERT INTO Users (Id, Name) VALUES (?, ?)`
	_, err = conn.Exec(q, id, name)
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

func InserSegment(id int, segment string) error {
	config := config.GetConfig()
	conn, err := mysql.NewClient(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Connection: %v\n", err)
		os.Exit(1)
	}
	q := `INSERT INTO Segments (Id, Segment) VALUES (?, ?)`
	_, err = conn.Exec(q, id, segment)
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
