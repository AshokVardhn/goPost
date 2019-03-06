package main

import "fmt"

func main() {
	PrintMe("a", "b")
	printme(1, 2)
}

func PrintMe(a, b string) {
	fmt.Println(a + b)
}
