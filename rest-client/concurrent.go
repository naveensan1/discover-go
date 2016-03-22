package main

import (
	"fmt"
	"time"
)

func conMovieSearch() {
	movieName := readCommandLine()
	startTime := time.Now()
	movies := searchMovies(movieName)
	defer fmt.Printf("execution time is %s\n", time.Since(startTime).String())
	for _, movie := range movies {
		go getMovieInfo(movie.ImdbID)
	}

}
