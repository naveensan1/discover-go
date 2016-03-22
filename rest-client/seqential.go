package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func readMovie() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Movie Name: ")
	movieName, _ := reader.ReadString('\n')
	fmt.Println(movieName)
	if movieName == "\n" {
		movieName = "Batman"
	}
	movies := searchMovies(movieName)

	startTime := time.Now()
	for _, movie := range movies {
		getMovieInfo(movie.ImdbID)
	}
	fmt.Printf("execution time is %s\n", time.Since(startTime).String())
}
