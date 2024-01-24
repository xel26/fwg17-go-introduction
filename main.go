package main

import "fmt"

func main() {
	PrintSegitiga(5)

	fmt.Println(GenPass("abcd", "low"))
	fmt.Println(GenPass("abcd", "med"))
	fmt.Println(GenPass("abcd", "strong"))

	fmt.Println(SistemMenonton(7))
}