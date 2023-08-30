package dependencies

import (
	"fmt"
	"os"
	"strconv"

	"github.com/IDL13/avito/internal/CSV"
	"github.com/IDL13/avito/internal/config"
	"github.com/IDL13/avito/internal/requests"
	"github.com/IDL13/avito/internal/timer"
	"github.com/IDL13/avito/pkg/mysql"
)

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
		var S requests.Set
		err = row.Scan(&S.Id, &S.Segment)
		if err != nil {
			panic(err)
		}
		_, ok := m[S.Id]
		if ok {
			m[S.Id] = append(m[S.Id], S.Segment)
		} else {
			m[S.Id] = append(m[S.Id], S.Segment)
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
		err = CSV.WriteInCSV([]string{strconv.Itoa(UserId), Segments[iter], "Add", timer.TimeNow()})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cyclic CSV writing error: %v\n", err)
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
		err = CSV.WriteInCSV([]string{strconv.Itoa(UserId), Segments[iter], "Del", timer.TimeNow()})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cyclic CSV writing error: %v\n", err)
			os.Exit(1)
		}
	}
	return nil
}
