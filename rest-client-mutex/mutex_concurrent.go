package main

import (
	"fmt"
	"sync"
	"time"
	"net/url"
)

type movieList struct {
	sync.Mutex
	movies []string
}

var m = new(movieList)

func movieInfo(movieImdbID string) {
	baseURL, err := url.Parse("http://www.omdbapi.com/?plot=short&r=json")
	if err != nil {
		fmt.Println("Couldn't Parse url")
	}
	values := baseURL.Query()
	values.Add("i", movieImdbID)
	baseURL.RawQuery = values.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Println("Couldn't load url")
	}
	movie := new(Movie)
	decoder := json.NewDecoder(resp.Body)
	error := decoder.Decode(&movie)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	rating, err := strconv.ParseFloat(movie.ImdbRating, 32)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	m.Lock()
	defer m.Unlock()
	m.movies = append(m.movies, fmt.Sprintf("The movie : %s was released in %s - the IMBD rating is %.0f%% with %s votes\n", movie.Title, movie.Year, rating*10, movie.ImdbVotes))
}

func conMovieSearch() {
	movieName := readCommandLine()
	startTime := time.Now()
	movies := searchMovies(movieName)
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

}

func main() {
	for {
		conMovieSearch()
	}
}