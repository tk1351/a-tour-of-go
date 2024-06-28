package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("This value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("This value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("This value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("This value:", v, "Present?", ok)
}
