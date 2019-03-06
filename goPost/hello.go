package main

import "fmt"

func main() {
	PrintMe("a", "b")
}

func PrintMe(a, b string) {
	fmt.Println(a + b)
}
