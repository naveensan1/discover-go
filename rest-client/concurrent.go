package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func readMovieConcurrent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Movie Name: ")
	movieName, _ := reader.ReadString('\n')
	if movieName == "\n" {
		movieName = "Batman"
	}
	startTime := time.Now()
	movies := searchMovies(movieName)
	for _, movie := range movies {
		go getMovieInfo(movie.ImdbID)
	}
	fmt.Printf("execution time is %s\n", time.Since(startTime).String())
}
