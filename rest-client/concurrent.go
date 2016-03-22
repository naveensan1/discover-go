package main

import (
	"fmt"
	"sync"
	"time"
)

func conMovieSearch() {
	movieName := readCommandLine()
	startTime := time.Now()
	movies := searchMovies(movieName)
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

}
