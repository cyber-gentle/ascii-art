package main

import "fmt"

func getCharRows(ch rune, bannerLines []string) []string {
	index := int(ch) - 32
	start := index * 9
	return bannerLines[start+1 : start+9]

}

func renderLine(word string, bannerLines []string) {
	var row []string

	wordsSlice := []rune(word)
	for _, char := range wordsSlice {
		row = getCharRows(char, bannerLines)
	}
	fmt.Println(row)
}
