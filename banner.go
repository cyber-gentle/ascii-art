package main

import (
	"fmt"
	"os"
	"strings"
)

func getChar

func loadBanner(bannerName string) {
	banners, err := os.ReadFile(bannerName)
	// if err != nil {
	// 	return nil, err
	// }

	if err != nil {
		fmt.Println(err)
	}

	bannerLines := strings.Split(string(banners), "\\n")
	fmt.Println(bannerLines)
	
}

//  (map[rune]string, error)
