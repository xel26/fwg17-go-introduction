package main

import (
	"fmt"
	"math/rand"
	"strings"
)


func handlePanic(){
	if a := recover(); a != ""{
		fmt.Println("RECOVER", a)
	}
}

// 	pass := "rahasia"
//	complex := ""
//  password = &pass
//  fmt.println(*password)
//  fmt.prntln(&password)


func GenPass(password *string, complexity *string) string {
	// defer handlePanic()
	// fmt.Println("program tidak crash, dan dapat mengambil tindakan setelah terjadi panic")

	if *password == ""{
		panic("Error: password must be filled")
	}

	if len(*password) < 6{
		panic("Error: password must be at least 6 character")
	}

	if *complexity == ""{
		panic("Error: complexity must be filled ")
	}



	upperCase := strings.ToUpper(*password)
	number := "0123456789"
	symbol := "`@#$%&*()-[]|/"

	addUpperCase := func() string{
		amountUpperCase := rand.Intn(len(*password)-2) + 1

		passwordSlice := []string(strings.Split(*password, ""))

		for i := 1; i <= amountUpperCase; i++{
			randomIndex := rand.Intn(len(*password))
			passwordSlice[randomIndex] = string(upperCase[randomIndex])
		}

		*password = strings.Join(passwordSlice, "")
		return *password
	}



	addNumber := func() string{
		*password = addUpperCase()
		fmt.Println(*password)

		amountNumber := rand.Intn(len(*password)-3) + 1
		fmt.Println(amountNumber)

		result :=[]string{}

		headSlice := []string(strings.Split(*password, ""))

		backSlice := []string(strings.Split(*password, ""))

		for i := 1; i <= amountNumber; i++{
			randomIndex := rand.Intn(len(*password))
			fmt.Println(randomIndex)

			backSlice := backSlice[:randomIndex]

			backSlice = append(backSlice, string(number[randomIndex]))

			result = append(backSlice, strings.Join(headSlice[randomIndex:], ""))

			backSlice = result

			headSlice = result
		}

		*password = strings.Join(result, "")
		return *password
	}



	addSymbol := func() string{
		*password = addNumber()

		amountSymbol := rand.Intn(len(*password)-3) + 1

		result :=[]string{}

		headSlice := []string(strings.Split(*password, ""))

		backSlice := []string(strings.Split(*password, ""))

		for i := 1; i <= amountSymbol; i++{
			randomIndex := rand.Intn(len(*password))

			backSlice := backSlice[:randomIndex]

			backSlice = append(backSlice, string(symbol[randomIndex]))

			result = append(backSlice, strings.Join(headSlice[randomIndex:], ""))
			
			backSlice = result

			headSlice = result
		}
		*password = strings.Join(result, "")
		return *password
	}




	if *complexity == "low"{
		*password =  addUpperCase()
	} else if *complexity == "med"{
		*password = addNumber()
	}else{
		*password = addSymbol()
	}



	return *password
}