package main

import (
	"fmt"
	"math/rand"
	"strings"
)


func GenPass(password *string, complexity *string) (string, bool) {
	
	if *password == ""{
		 return "Error: password must be filled", false
	}
	
	if len(*password) < 6{
		return "Error: password must be at least 6 character", false
	}

	if *complexity == ""{
		return "Error: complexity must be filled", false
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
		valuePassword := addUpperCase()
		fmt.Println("password dari addUpperCase ", valuePassword)

		amountNumber := rand.Intn(len(*password)-3) + 1

		


		for i := 1; i <= amountNumber; i++{
			randomIndex := rand.Intn(len(*password))

			valuePassword = valuePassword[:randomIndex] + string(number[randomIndex]) + valuePassword[randomIndex:]
		}

		*password = valuePassword
		return *password
	}



	addSymbol := func() string{
		valuePassword := addNumber()

		amountSymbol := rand.Intn(len(*password)-3) + 1



		for i := 1; i <= amountSymbol; i++{
			randomIndex := rand.Intn(len(*password))

			valuePassword = valuePassword[:randomIndex] + string(symbol[randomIndex]) + valuePassword[randomIndex:]
		}

		*password = valuePassword
		return *password
	}




	if *complexity == "low"{
		*password =  addUpperCase()
	} else if *complexity == "med"{
		*password = addNumber()
	}else{
		*password = addSymbol()
	}



	return *password, true
}