package main

import (
	"ascii-art/gen2"
	"fmt"
	"os"
)

var (
	input_string       string
	first_banner_name  string
	second_banner_name string
	third_banner_name  string
)

func main() {

	switch len(os.Args) {
	case 1:
		fmt.Println("\n Usage: go run . <words to process to design.>  <file name of the design type.> \n E.g. go run . Hello  shadow.txt \n Or   go run . Hello")
		return

	case 2:
		input_string = os.Args[1]
		first_banner_name = "standard"

	case 3:
		input_string = os.Args[1]
		first_banner_name = gen2.LowerCase(os.Args[2])

	// case 4:
	// 	input_string = os.Args[1]
	// 	first_banner_name = gen2.LowerCase(os.Args[2])
	// 	second_banner_name = gen2.LowerCase(os.Args[3])

	// case 5:
	// 	input_string = os.Args[1]
	// 	first_banner_name = gen2.LowerCase(os.Args[2])
	// 	second_banner_name = gen2.LowerCase(os.Args[3])
	// 	third_banner_name = gen2.LowerCase(os.Args[4])

	default:
		fmt.Println("\n Usage: go run . <words to process to design.>  <file name of the design type.> \n E.g. go run . Hello  shadow.txt \n Or   go run . Hello")
		return
	}

	content, err := os.ReadFile("banners/" + first_banner_name + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))

}
