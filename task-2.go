package main

import (
	"math/rand"
	"strings"
)

func GenPass(password string, complexity string) string {
	minLength := 0
	if len(password) < 6{
		minLength = 6
	}else{
		minLength = len(password)
	}

	lowerAlphabet := strings.ToLower(password)
	upperAlphabet := strings.ToUpper(password)
	numeric := "0123456789"
	symbol := "`~!@#$%^&*()-=_+[]{}|;':,./<>?"

	generator := ""
	if complexity == "low"{
		for i := 0; i < minLength; i++ {
			lower := lowerAlphabet[rand.Intn(len(lowerAlphabet))]
			upper := upperAlphabet[rand.Intn(len(upperAlphabet))]
			generator += string(lower) + string(upper)
		}
	}

	if complexity == "med"{
		for i := 0; i < minLength; i++ {
			lower := lowerAlphabet[rand.Intn(len(lowerAlphabet))]
			upper := upperAlphabet[rand.Intn(len(upperAlphabet))]
			num := numeric[rand.Intn(len(numeric))]
			generator += string(upper) + string(num) + string(lower)
		}
	}

	if complexity == "strong"{
		for i := 0; i < minLength; i++ {
			lower := lowerAlphabet[rand.Intn(len(lowerAlphabet))]
			upper := upperAlphabet[rand.Intn(len(upperAlphabet))]
			num := numeric[rand.Intn(len(numeric))]
			char := symbol[rand.Intn(len(symbol))]
			generator += string(upper) + string(lower) + string(char) + string(num)
		}
	}

	return generator
}