package main

import "fmt"

func SistemMenonton(n int) (string, bool) {
	data := []int{1, 7, 3, 4, 8, 9}

	var output string

	for index, value := range data {
		for i := 1; i < len(data) - index; i++ {
			
			if value + data[i] == n {
				output = fmt.Sprintf("%v dan %v", value, data[i])
				break
			}
		}

		if(output != ""){
			break
		}
	}

	if output == ""{
		return "data tidak ditemukan", false
	}

	return output, true
}