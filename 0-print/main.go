package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Hello we are Holberton School\n")
	fmt.Printf("the date is %s\n", time.Now())
	fmt.Printf("the year is %d\n", time.Now().Year())
	fmt.Printf("the month is %s\n", time.Now().Month())
	fmt.Printf("the day is %d\n", time.Now().Day())
}
