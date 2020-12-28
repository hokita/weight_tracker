package main

import (
	"fmt"
	"os"

	"github.com/hokita/weight_tracker/http"
)

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
