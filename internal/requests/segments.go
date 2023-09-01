package requests

import (
	"errors"
	"fmt"
	"os"

	"github.com/IDL13/avito/internal/config"
	"github.com/IDL13/avito/pkg/mysql"
)

func (s *server) InserSegment(segment string) error {
	if len(segment) > 1 {
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
	} else {
		return errors.New("empty segment")
	}
}

func (s *server) DeleteSegment(segment string) error {
	if len(segment) > 1 {
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
	} else {
		return errors.New("empty json")
	}
}
