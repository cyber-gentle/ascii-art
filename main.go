package main

import (
	"ascii-art/gen2"
	"fmt"
	"os"
)

var (
	input_string string
	banner_name  string
	banner_type  map[int]string
)

func main() {

	switch len(os.Args) {
	case 1:
		fmt.Println("\n Usage: go run . <words to process to design.>  <file name of the design type.> \n E.g. go run . Hello  shadow.txt \n Or   go run . Hello")
		return

	case 2:
		input_string = os.Args[1]

	case 3:
		input_string = os.Args[1]
		banner_name = gen2.LowerCase(os.Args[2])
	}

	fmt.Println(input_string)
	fmt.Println(banner_name)

}
