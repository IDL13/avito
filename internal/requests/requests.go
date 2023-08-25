package request

import (
	"fmt"
	"os"

	"github.com/IDL13/avito/pkg/mysql"
)

func CreateTables() error {
	conn, err := mysql.NewClient("avito", "mysqlcontainer", "3306", "db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from requests in Insert function: %v\n", err)
		os.Exit(1)
	}
	q := `CREATE TABLE Users
(
    Id INT IDENTITY (1, 1) NOT NULL,
    Name NVARCHAR(MAX),
    CONSTRAINT PK_User_Id PRIMARY KEY (Id)
);

CREATE TABLE Segments
(
    Id INT NOT NULL,
    UserId INT NOT NULL,
    Segment VARCHAR(512),
    CONSTRAINT PK_Segment_Id PRIMARY KEY (Id),
    CONSTRAINT FK_Segment_User FOREIGN KEY (UserId)
        REFERENCES Users (Id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);`
	row, err := conn.Query(q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Insert date: %v\n", err)
		os.Exit(1)
	}
	defer row.Close()
	return nil
}

func Insert() error {
	conn, err := mysql.NewClient("avito", "mysqlcontainer", "3306", "db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from requests in Insert function: %v\n", err)
		os.Exit(1)
	}
	q := `INSERT INTO segments_table (user_id, segments) VALUES (?, ?) RETURNING id`
	row, err := conn.Query(q, "asdfasdf", "asdfsf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Insert date: %v\n", err)
		os.Exit(1)
	}
	defer row.Close()
	return nil
}
