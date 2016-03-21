package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)

	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	}

	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Printf("I'm a student of the %s\n", school)
	}

	beautifulWeather := true
	if beautifulWeather {
		fmt.Println("It's a beautiful weather!")
	}

	holbertonFounder := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	for _, v := range holbertonFounder {
		fmt.Printf("%s is a founder\n", v)
	}
}
