package CSV

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

//{"UserID", "Segment", "Add/Remove", "Date-Time"}

func CreateCSV() error {
	csvFile, err := os.OpenFile("./csv_log.csv", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from reading csv file: %v\n", err)
		os.Exit(1)
	}

	w := csv.NewWriter(csvFile)
	w.Write([]string{"UserID", "Segment", "Add/Remove", "Date-Time"})
	w.Flush()

	if err := w.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "Error from csv writer: %v\n", err)
		os.Exit(1)
	}

	defer csvFile.Close()

	return nil
}

func WriteInCSV(reconds []string) error {
	csvFile, err := os.OpenFile("./csv_log.csv", os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from reading csv file: %v\n", err)
		os.Exit(1)
	}

	w := csv.NewWriter(csvFile)
	w.Write(reconds)
	w.Flush()

	if err := w.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "Error from csv writer: %v\n", err)
		os.Exit(1)
	}

	defer csvFile.Close()

	return nil
}

func ReadInCSV(time string) map[string][][]string {
	if err := os.Chmod("./csv_log.csv", 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error from Chmod operation: %v\n", err)
		os.Exit(1)
	}

	csvFile, err := os.Open("./csv_log.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from reading csv file: %v\n", err)
		os.Exit(1)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ','

	m := make(map[string][][]string)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if record[3] == time {
			m[time] = append(m[time], record)
		}
	}

	return m
}
