package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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
	fmt.Printf("The movie : %s was released in %s - the IMBD rating is %.0f%% with %s votes\n", movie.Title, movie.Year, rating*10, movie.ImdbVotes)
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

func main() {
	for {
		//readMovie()
		readMovieConcurrent()
	}
}
