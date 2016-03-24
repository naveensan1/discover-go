package main

import (
	"flag"
	"fmt"
)

func main() {
	movie := flag.String("movie", "", "name of the movie")
	blahblah := flag.Int("blah", 0, "naveen's")

	flag.Parse()

	fmt.Println(*movie)
	fmt.Println(*dicksize)

}
