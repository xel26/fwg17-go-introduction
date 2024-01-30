package main

import "fmt"

func handlePanic(){
	a := recover();

	if 	a != ""{
		fmt.Println("RECOVER", a)
	}
}

func main() {
	defer handlePanic()

	PrintSegitiga(5)

	pass := "password"
	complex := "low"
	result, status := GenPass(&pass, &complex)
	if status == true {
		fmt.Println("success create password: ", result)
	}else{
		panic(result)
	}


	res, found := SistemMenonton(7)
	if found == true {
		fmt.Println("data ditemukan: ", res)
	}else{
		panic(res)
	}
}