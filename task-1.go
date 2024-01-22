package main

import "fmt"

func printSegitiga(lines int) {
	fmt.Printf("printSegitiga %v \n", lines)
	output := ""

	for i := lines; i >=1; i-- {
		for j := lines - i; j > 0; j--{
			output += " "
		}

		for k := 1; k <= i * 2 - 1; k++{
			output += "*"
		}

		fmt.Println(output)
		output = ""
	}

}