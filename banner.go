package main

import (
	"os"
	"strings"
)

func loadBanner(bannerName string) []string {
	content, err := os.ReadFile(bannerName)
	if err != nil {
		logErr(err)
	}

	bannerLines := strings.Split(string(content), "\n")
	return bannerLines
}
