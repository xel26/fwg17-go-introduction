package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"strings"
// )

// func handlePanic(){
// 	if a := recover(); a != ""{
// 		fmt.Println("RECOVER", a)
// 	}
// }

// // 	pass := "rahasia"
// //	complex := ""
// //  var password *string = &pass
// //  fmt.println(*password)
// //  fmt.prntln(&password)

// func GenPass(password *string, complexity *string) string {
// 	// defer handlePanic()
// 	// fmt.Println("program tidak crash, dan dapat mengambil tindakan setelah terjadi panic")

// 	if *password == ""{
// 		panic("Error: password must be filled")
// 	}

// 	if *complexity == ""{
// 		panic("Error: complexity must be filled ")
// 	}

// 	minLength := len(*password)
// 	if len(*password) < 6{
// 		minLength = 6
// 	}

// 	lowerAlphabet := strings.ToLower(*password)
// 	upperAlphabet := strings.ToUpper(*password)
// 	numeric := "0123456789"
// 	symbol := "`~!@#$%^&*()-=_+[]{}|;':,./<>?"

// 	generator := ""
// 	if *complexity == "low"{
// 		for i := 0; i < minLength; i++ {
// 			lower := lowerAlphabet[rand.Intn(len(lowerAlphabet))]
// 			upper := upperAlphabet[rand.Intn(len(upperAlphabet))]
// 			generator += string(lower) + string(upper)
// 		}
// 	}

// 	if *complexity == "med"{
// 		for i := 0; i < minLength; i++ {
// 			lower := lowerAlphabet[rand.Intn(len(lowerAlphabet))]
// 			upper := upperAlphabet[rand.Intn(len(upperAlphabet))]
// 			num := numeric[rand.Intn(len(numeric))]
// 			generator += string(upper) + string(num) + string(lower)
// 		}
// 	}

// 	if *complexity == "strong"{
// 		for i := 0; i < minLength; i++ {
// 			lower := lowerAlphabet[rand.Intn(len(lowerAlphabet))]
// 			upper := upperAlphabet[rand.Intn(len(upperAlphabet))]
// 			num := numeric[rand.Intn(len(numeric))]
// 			char := symbol[rand.Intn(len(symbol))]
// 			generator += string(upper) + string(lower) + string(char) + string(num)
// 		}
// 	}

// 	return generator
// }