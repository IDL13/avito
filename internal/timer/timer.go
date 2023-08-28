package timer

import (
	"fmt"
	"os"
	"time"
)

const timeLayout = "2006-01-02T15:04:00"

func CallAt(callTime string, f func(int, []string) error, UserId int, Segments []string) error {
	cTime, err := time.Parse(timeLayout, callTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshaling error:%v", err)
		os.Exit(1)
	}
	n := time.Now().Format(timeLayout)
	now, _ := time.Parse(timeLayout, n)
	duration := cTime.Sub(now)
	go func() error {
		time.Sleep(duration)
		err = f(UserId, Segments)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Eror from Isert or delet function:%v", err)
			os.Exit(1)
		}
		return nil
	}()
	return nil
}
