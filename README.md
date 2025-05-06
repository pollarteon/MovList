# MOVLIST - Movie Watchlist CLI Application

MOVLIST is a Command Line Interface (CLI) application that allows you to maintain a personal movie watchlist. You can add and remove movies from your list, search for movies using the OMDB API, and store those movies to your watchlist for later reference.

## Features

- **Add/Remove Movies from Watchlist**: Easily manage your movie watchlist by adding and removing movies.
- **Search Movies**: Use the OMDB API to search for movies by title, year, or other details.
- **Maintain a Personal Movie Collection**: Track movies you're interested in watching with an interactive CLI.
- **Requires an Internet Connection**: Fetch movie details from OMDB API in real-time.

## Requirements

1. **OMDB API Key**: You must have an active OMDB API key. You can obtain one by signing up at [OMDB API](http://www.omdbapi.com/apikey.aspx).
   
   - Once you get your key, make sure to set it in the environment variable `MOVIE_API_KEY`.

2. **Go Runtime**: This application is written in Go. Ensure that Go is installed on your system to build and run the application locally.

3. **Active Internet Connection**: The application requires an internet connection to search movies and fetch information from OMDB.

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/pollarteon/MovList.git
/movlist.git
cd movlist
```
2. Set up your OMDB API key:
Copy the .env.example to .env and set your API key:
```bash
cp .env.example .env
```
Edit the .env file and add your OMDB API key:

```vscode
MOVIE_API_KEY=your_api_key_here
```
3. Build and run the Application:
```bash
go build -o movlist
./movlist
```

## You will be able to interact with the application to:

- **Add movies**: Add a movie to your watchlist by searching for it using the OMDB API.
- **Remove movies**: Remove movies from your watchlist.
- **View watchlist**: View the current list of movies in your watchlist.
- **Search for movies**: Use the search feature to find movies by title or year.

## Example Commands:

- To **add** a movie to your watchlist, simply search for it and select the option to add it.
- To **remove** a movie, select the movie from the watchlist and choose the option to remove it.

