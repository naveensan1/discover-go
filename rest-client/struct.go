package main

//Movie - a struct of a movie
type Movie struct {
	Title      string
	Year       string
	ImdbRating float32 `json:"ImdbRating,string"`
	ImdbVotes  string
	ImdbID     string
}

//MovieQuery - Search Results
type MovieQuery struct {
	Search []Movie
}
