package main

import (
	"ascii-art/gen2"
	"fmt"
	"os"
	"strings"
)

var (
	inputtedText     string
	firstBannerName  string
	secondBannerName string
	thirdBannerName  string
)

func main() {

	switch len(os.Args) {
	case 2:
		inputtedText = os.Args[1]
		firstBannerName = "standard"

	case 3:
		inputtedText = os.Args[1]
		firstBannerName = gen2.LowerCase(os.Args[2])

	// case 4:
	// 	inputtedText = os.Args[1]
	// 	firstBannerName = gen2.LowerCase(os.Args[2])
	// 	secondBannerName = gen2.LowerCase(os.Args[3])

	// case 5:
	// 	inputtedText = os.Args[1]
	// 	firstBannerName = gen2.LowerCase(os.Args[2])
	// 	secondBannerName = gen2.LowerCase(os.Args[3])
	// 	thirdBannerName = gen2.LowerCase(os.Args[4])

	default:
		fmt.Println("\n Usage: go run . <words to process to design.>  <file name of the design type.> \n E.g. go run . Hello  shadow.txt \n Or   go run . Hello")
		return
	}

	bannerLines := loadBanner("banners/" + firstBannerName + ".txt")

	inputtedTextSlice := strings.Split(inputtedText, `\n`)
	for i := range inputtedTextSlice {
		if inputtedTextSlice[i] == "" {
			fmt.Println()
		} else {
			renderLine(inputtedTextSlice[i], bannerLines)
		}
	}

}
