package API

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

// Define structs to match the JSON response
type Movie struct {
    Title  string `json:"Title"`
    Year   string `json:"Year"`
    IMDbID string `json:"imdbID"`
    Type   string `json:"Type"`
    Poster string `json:"Poster"`
}

type SearchResponse struct {
    Search       []Movie `json:"Search"`
    TotalResults string  `json:"totalResults"`
    Response     string  `json:"Response"`
}

func Search(MovieTitle string) (SearchResponse, error) {

	// Load API Key
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err)
		return SearchResponse{}, err
	}


	MOVIE_API_KEY :=os.Getenv("MOVIE_API_KEY");
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&type=movie", MOVIE_API_KEY, MovieTitle)
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	// Parse JSON into the struct
	var apiResponse SearchResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return SearchResponse{}, err
	}

	return apiResponse, nil
}