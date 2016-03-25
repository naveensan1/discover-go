package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func (u *user) Hello() {
	localUser := *u
	fmt.Printf("Hello %s\n", localUser.Name)
}

func (u *user) Info() {
	localUser := *u
	currentYear := time.Now().Year()
	userYearStr := strings.Split(localUser.DOB, ",")[1][1:]
	userYearInt, err := strconv.Atoi(userYearStr)
	if err != nil {
		fmt.Println("Couldn't convert user's age from str to int")
	}
	age := currentYear - userYearInt
	fmt.Printf("%s who was born in %s would be %d years old today\n", localUser.Name, localUser.City, age)
}

func main() {
	u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
	u.Hello()
	u.Info()
}
