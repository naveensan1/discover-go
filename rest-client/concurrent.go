package main

import (
	"fmt"
	"sync"
	"time"
)

func conMovieSearch() error {
	movieName := readCommandLine()
	startTime := time.Now()
	movies, err := searchMovies(movieName)
	if err != nil {
		fmt.Printf("conMovieSearch: %s\n", err)
		return err
	}
	var wg sync.WaitGroup

	for _, movie := range movies {
		wg.Add(1)
		go func(movie Movie) {
			defer wg.Done()
			getMovieInfo(movie.ImdbID)
		}(movie)
	}
	wg.Wait()
	fmt.Printf("execution time is %s\n", time.Since(startTime).String())
	return nil
}
