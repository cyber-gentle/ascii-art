package main

import (
	"fmt"
	"os"
)

// func fileReader
func fileReader(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// func logErr that prints error to the standard output.
func logErr(err error) {
	fmt.Println(err)
}
