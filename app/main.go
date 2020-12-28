package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hokita/weight_tracker/http"
)

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	if err := http.Start(); err != nil {
		return err
	}
	return nil
}
