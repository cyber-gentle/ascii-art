package main

func getCharRows(ch rune, bannerLines []string) []string {
	index := int(ch) - 32
	start := index * 9
	return bannerLines[start+1 : start+9]

}

func renderLine() {

}
