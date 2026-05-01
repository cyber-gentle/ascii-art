package main

import (
	"ascii-art/utils"
	"errors"
	"os"
	"strings"
)

func main() {
	var (
		//	inputtedText string
		bannerName string
		err        error
	)

	switch len(os.Args) {
	case 2:
		//	inputtedText = os.Args[1]
		bannerName = "standard.txt"

	case 3:
		//	inputtedText = os.Args[1]
		bannerName = os.Args[2]
		if !strings.HasSuffix(bannerName, ".txt") {
			bannerName += ".txt"
		}

	default:
		err = errors.New("Usage: go run . <\"text\"> <banner style>")
		utils.LogError(err)
		return
	}

	loadBanner("banners/" + bannerName)

}

func splitInput()

