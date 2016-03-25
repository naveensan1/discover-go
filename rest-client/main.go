package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type movieList struct {
	sync.Mutex
	movies []string
}

var m = new(movieList)

func readCommandLine() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Movie Name: ")
	movieName, _ := reader.ReadString('\n')
	fmt.Println(movieName)
	if movieName == "\n" {
		movieName = "Batman"
	}
	return movieName
}

func searchMovies(movieName string) []Movie {
	urlLocal, err := url.Parse("http://www.omdbapi.com/")
	if err != nil {
		fmt.Println("Couldn't load url")
	}
	values := urlLocal.Query()
	values.Add("s", movieName)
	urlLocal.RawQuery = values.Encode()
	resp, err := http.Get(urlLocal.String())
	if err != nil {
		fmt.Println("Couldn't load url")
	}
	movieQuery := new(MovieQuery)
	decoder := json.NewDecoder(resp.Body)
	error := decoder.Decode(&movieQuery)
	if error != nil {
		fmt.Println("we have problems")
	}
	return movieQuery.Search
}

func getMovieInfo(movieImdbID string) {
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

func main() {
	for {
		//seqMovieSearch()
		conMovieSearch()
	}
}
