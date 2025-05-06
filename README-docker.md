
# Docker Setup for MovList

  

## Prerequisites

1. [Docker](https://docs.docker.com/get-docker/) installed

2. Read the [main README.md](README.md) for general project setup

3. OMDB API key from [omdbapi.com](http://www.omdbapi.com/apikey.aspx)

  

## Quick Start

  

### Build and Run the Image

```bash
docker  build  -t  movlist  .
docker run -it -e MOVIE_API_KEY=<your_omdbapi_key> -v "$(pwd)/db1:/app/db" movlist
movlist

```

### OR
```bash
docker run -it -e  MOVIE_API_KEY=<your_omdbapi_key> -v "$(pwd)/db1:/app/db" prafulh/movlist
```

## Configuration

### Environment Variables

Variable Required

```bash

MOVIE_API_KEY  = <Your  OMDB  API  key>

```


  

### Troubleshooting

"Error loading .env file"

Expected behavior when using -e flags

  

Application prioritizes environment variables over .env file



Note: Replace <your_omdb_api_key> with your actual API key and <your_key> in examples.