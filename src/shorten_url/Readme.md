# URL Shortener

This repo contains the Go source code for converting the given `URL` to a shortened format.

## Requirements Implemented

- API call to convert the given URL to shortened URL `http://localhost:8086/v1/<url_to_covert>`
- API call to fetch the previously given URL to get re-fetch shortened URL `http://localhost:8086/v2/<url_to_covert>`
- Go test files added to make sure the source code works fine, with mock test scenarios added in test file (by running `go test`)
- Created `Dockerfile` to execute the program using `docker build` and `docker run` command
- Image of `urlshortener` pushed into docker hub registry (`gsdockit/urlshortener`) for easy/quick access

## How to run

There are three ways the program can be executed

### Method 1

- Clone the repository `https://github.com/chefgs/golang`
- Go to the path `src/shorten_url`
- Download dependencies by `go mod tidy`
- Execute `go build cmd/url_shortener`
- It will create an executable file `url_shortener`. Run the file `./url_shortener`
- It will run the API program in `http://localhost:8086`

### Method 2

- Clone the repository `https://github.com/chefgs/golang`
- Go to the path `src/shorten_url`
- Run docker commands to build an and run the it in a container

```cmd
docker build -t urlshortener . --rm
docker run -d -p 8086:8086 --rm urlshortener
```

- It will run the program in `http://localhost:8086`

### Method 3

- A docker image is available in `hub.docker.io` for this application
- Users can directly run the command `docker run -d -p 8086:8086 --rm gsdockit/urlshortener`
- It will run the program in `http://localhost:8086`

## How to use URL shortener

- Once the program runs on `http://localhost:8086`, the API is ready is accept requests
- convert the URL by adding the `URL value` at the end of `host:port/url_to_convert` API call
- For example: If user wants to shorten the URL `mywebsite.com`, it can be done using `http://localhost:8086/v1/mywebsite.com`

## How to fetch stored URL

- The URL shortner stores the previously converted URLs in cache file
- While the program runs on `localhost`, Users can access it using the API call `http://localhost:8086/v1/fetchcachedurl/previously_converted_url_input`
- For example: If user wants to fetch the shortened URL of `mywebsite.com`, it can be fetched using `http://localhost:8086/v1/fetchcachedurl/mywebsite.com`
