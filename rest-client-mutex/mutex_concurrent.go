package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type movieList struct {
	sync.Mutex
	movies []string
}

var m = new(movieList)

func readCommandLine() string {
	movie := flag.String("movie", "Batman", "Name of the Movie")
	flag.Parse()
	return *movie
}

func searchMovies(movieName string) ([]Movie, error) {
	urlLocal, error := url.Parse("http://www.omdbapi.czm/")
	if error != nil {
		return nil, error
	}
	values := urlLocal.Query()
	values.Add("s", movieName)
	urlLocal.RawQuery = values.Encode()
	resp, error := http.Get(urlLocal.String())
	if error != nil {
		return nil, error
	}
	movieQuery := new(MovieQuery)
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&movieQuery); err != nil {
		return nil, err
	}
	return movieQuery.Search, nil
}

func movieInfo(movieImdbID string) error {
	baseURL, error := url.Parse("http://www.omdbapi.com/?plot=short&r=json")
	if error != nil {
		return error
	}
	values := baseURL.Query()
	values.Add("i", movieImdbID)
	baseURL.RawQuery = values.Encode()
	resp, error := http.Get(baseURL.String())
	if error != nil {
		return error
	}
	movie := new(Movie)
	decoder := json.NewDecoder(resp.Body)
	if error := decoder.Decode(&movie); error != nil {
		return error
	}

	m.Lock()
	defer m.Unlock()
	m.movies = append(m.movies, fmt.Sprintf("The movie : %s was released in %s - the IMBD rating is %.0f%% with %s votes\n", movie.Title, movie.Year, movie.ImdbRating*10, movie.ImdbVotes))
	return nil
}

func conMovieSearch() error {
	movieName := readCommandLine()
	startTime := time.Now()
	movies, error := searchMovies(movieName)
	if error != nil {
		return fmt.Errorf("searchMovies: %s\n", error)
	} else if movies == nil {
		return fmt.Errorf("searchMovies: Could not find any results for that movie\n")
	}

	var wg sync.WaitGroup
	for _, movie := range movies {
		wg.Add(1)
		go func(movie Movie) {
			defer wg.Done()
			movieInfo(movie.ImdbID)
		}(movie)
	}
	wg.Wait()
	for _, x := range m.movies {
		fmt.Printf("%s", x)
	}
	m = new(movieList)
	fmt.Printf("execution time is %s\n\n\n", time.Since(startTime).String())
	return nil
}

func main() {
	if error := conMovieSearch(); error != nil {
		fmt.Printf("%s", error)
	}
}
