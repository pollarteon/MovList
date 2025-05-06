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
    MovieTitle  string `json:"Title"`
    Year   string `json:"Year"`
    IMDbID string `json:"imdbID"`
    Type   string `json:"Type"`
    Poster string `json:"Poster"`
	Watched bool `json:"watched"`
}

type SearchResponse struct {
    Search       []Movie `json:"Search"`
    TotalResults string  `json:"totalResults"`
    Response     string  `json:"Response"`
}


type SearchByIDResponse struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	Rated string `json:"Rated"`
	Released string `json:"Released"`
	Runtime string `json:"Runtime"`
	Genre string `json:"Genre"`
	Director string `json:"Director"`
	Writer string `json:"Writer"`
	Actors string `json:"Actors"`
	Plot string `json:"Plot"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Awards string `json:"Awards"`
	Poster string `json:"Poster"`
	Ratings []Ratings `json:"Ratings"`
	Metascore string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes string `json:"imdbVotes"`
	ImdbID string `json:"imdbID"`
	Type string `json:"Type"`
	Dvd string `json:"DVD"`
	BoxOffice string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website string `json:"Website"`
	Response string `json:"Response"`
}
type Ratings struct {
	Source string `json:"Source"`
	Value string `json:"Value"`
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

	res, err := http.DefaultClient.Do(req)
	if err!=nil{
		return SearchResponse{},err;
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	
	var apiResponse SearchResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return SearchResponse{}, err
	}

	return apiResponse, nil
}

func SearchByID(imdbId string)(SearchByIDResponse,error){
	//loading Movie API KEY
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err)
		return SearchByIDResponse{}, err
	}

	MOVIE_API_KEY :=os.Getenv("MOVIE_API_KEY");
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s&type=movie", MOVIE_API_KEY, imdbId)
	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err!=nil{
		return SearchByIDResponse{},err;
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var apiResponse SearchByIDResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return SearchByIDResponse{}, err
	}

	if apiResponse.Response=="False"{
		return apiResponse,nil
	}

	return apiResponse, nil
}
