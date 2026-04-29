package main

import (
	"os"
	"strings"
)

func loadBanner(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		logErr(err)
	}

	data := strings.Split(string(content), "\n")
	return data
}
