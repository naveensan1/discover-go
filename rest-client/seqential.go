package main

import (
	"fmt"
	"time"
)

func seqMovieSearch() error {
	movieName := readCommandLine()
	startTime := time.Now()
	movies, err := searchMovies(movieName)
	if err != nil {
		fmt.Printf("Error with searchMovies: %s", err)
		return err
	}
	for _, movie := range movies {
		getMovieInfo(movie.ImdbID)
	}
	fmt.Printf("execution time is %s\n", time.Since(startTime).String())
	return nil
}
