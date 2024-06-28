package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	w := strings.Fields(s)

	for _, ele := range w {
		wc[ele] = wc[ele] + 1
	}

	return wc
}

func main() {
	wc.Test(WordCount)
}
