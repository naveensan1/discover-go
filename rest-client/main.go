package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

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

func searchMovies(movieName string) ([]Movie, error) {
	movieQuery := new(MovieQuery)
	urlLocal, err := url.Parse("http://www.omdbapi.com/")
	if err != nil {
		fmt.Printf("searchMovies: %s\n", err)
		return movieQuery.Search, err
	}
	values := urlLocal.Query()
	values.Add("s", movieName)
	urlLocal.RawQuery = values.Encode()
	resp, err := http.Get(urlLocal.String())
	if err != nil {
		fmt.Printf("searchMovies: %s\n", err)
		return movieQuery.Search, err
	}
	decoder := json.NewDecoder(resp.Body)
	error := decoder.Decode(&movieQuery)
	if error != nil {
		fmt.Println("we have problems")
	}
	return movieQuery.Search, nil
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
		fmt.Printf("Over here: %s", error)
		return
	}
	fmt.Printf("The movie : %s was released in %s - the IMBD rating is %.0f%% with %s votes\n", movie.Title, movie.Year, movie.ImdbRating*10, movie.ImdbVotes)
}

func main() {
	for {
		//seqMovieSearch()
		conMovieSearch()
	}
}
