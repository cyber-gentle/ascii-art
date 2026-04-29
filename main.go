package main

import (
	"ascii-art/gen2"
	"fmt"
	"os"
)

var (
	inputtedText       string
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

	content := loadBanner("banners/" + firstBannerName + ".txt")

}
