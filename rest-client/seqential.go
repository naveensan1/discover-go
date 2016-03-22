package main

import (
	"fmt"
	"time"
)

func seqMovieSearch() {
	movieName := readCommandLine()
	startTime := time.Now()
	movies := searchMovies(movieName)
	for _, movie := range movies {
		getMovieInfo(movie.ImdbID)
	}
	fmt.Printf("execution time is %s\n", time.Since(startTime).String())
}
