package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	resp, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		fmt.Println("Couldn't load url")
	}
	fmt.Printf("status code is %s\n", resp.Status)
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
