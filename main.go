package main

import "fmt"

func main() {
	PrintSegitiga(5)

	pass := "fazztrack"
	complex := "strong"
	fmt.Println(GenPass(&pass, &complex))

	fmt.Println(SistemMenonton(7))
}