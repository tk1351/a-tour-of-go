package main

import "fmt"

func main() {
	var s []int
	printSlice(s) // len=0 cap=0 []

	s = append(s, 0)
	printSlice(s) // len=1 cap=1 [0]

	s = append(s, 1)
	printSlice(s) // len=2 cap=2 [0 1]

	s = append(s, 2, 3, 4)
	printSlice(s) // len=5 cap=6 [0 1 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
