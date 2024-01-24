package main

import "fmt"

func main() {
	PrintSegitiga(5)

	fmt.Println(GenPass("shella", "low"))
	fmt.Println(GenPass("shella", "med"))
	fmt.Println(GenPass("shella", "strong"))

	fmt.Println(SistemMenonton(7))
}