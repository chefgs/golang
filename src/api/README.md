# Go REST API framework

## What is REST API
REST is acronym for REpresentational State Transfer. It is architectural style for distributed hypermedia systems and was first presented by Roy Fielding in 2000 in his famous dissertation.

Like any other architectural style, REST also does have it’s own 6 guiding constraints which must be satisfied if an interface needs to be referred as RESTful. These principles are listed below.

REST is not a standard but a set of recommendations and constraints for RESTful web services. 

These include:

- Client-Server: SystemA makes an HTTP request to a URL hosted by SystemB, which returns a response. 
It’s identical to how a browser works. The application makes a request for a specific URL. The request is routed to a web server that returns an HTML page. That page may contain references to images, style sheets, and JavaScript, which incur further requests and responses.

- Stateless: REST is stateless: the client request should contain all the information necessary to respond to a request. 
In other words, it should be possible to make two or more HTTP requests in any order and the same responses will be received.

- Cacheable: A response should be defined as cacheable or not.

- Layered: The requesting client need not know whether it’s communicating with the actual server, a proxy, or any other intermediary.

- Uniform interface – REST is defined by four interface constraints: identification of resources; manipulation of resources through representations; self-descriptive messages; and, hypermedia as the engine of application state.

- Code on demand (optional) – REST allows client functionality to be extended by downloading and executing code in the form of applets or scripts. This simplifies clients by reducing the number of features required to be pre-implemented.

## What makes a RESTful API
A REST API consists of,
1. API Endpoint URL
2. HTTP Method (or Verb)
3. Response return from API request

### API Endpoint URL
An application implementing a RESTful API will define one or more URL endpoints with a domain, port, path, and/or querystring — for example, `https://somedomain/book/bookname?format=json`

### HTTP Method
The above defined API endpoint URL will be serving HTTP methods, which can be perform create, read, update and delete functions defined within the application.

- GET	- read -	returns requested data
- POST - create -	creates a new record
- PUT - update - updates an existing record
- DELETE - delete -	deletes an existing record

### The Response
The response payload can be whatever is practical: data, HTML, an image, an audio file, and so on. Data responses are typically JSON-encoded, but XML, CSV, simple strings, or any other format can be used. You could allow the return format to be specified in the request either in JSON or XML format.

Each response header returns either of the HTTP status codes,
```
200 OK
201 Created
400 Bad Request
404 Not Found
401 Unauthorized
500 Internal server error
```
## Setup Golang development environment
- Refer the [README here](https://github.com/chefgs/golang), to setup Golang dev environment.

## Create REST API using Golang
### Why to choose Go for REST API
- It is fast
- It is simple to understand
- It is compiled and works well with Microservices
### Go modules
We will be using below modules for our development
- net/http
- [Gin-gonic](https://github.com/gin-gonic/gin) REST API framework. It is a Go opensource utility module, which can be readily available to use.
It provides the REST API controller framework.

## Start Coding
- Create file `api.go`
- Every go code starts with package name definition
`package main`

- Then we need to declare the module import section
```
import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)
```

### In function main
- Add gin controller to handle the HTTP events coming to API

`router := gin.Default()`

- Add functions to handle the POST and GET requests on URL route path "/"
```
  router.POST("/", PostMethod)
  router.GET("/", GetMethod)
```

- Listen to the specified port of localhost
```
  listenPort := "4000"
  // Listen and Server on the LocalHost:Port
  router.Run(":"+listenPort)
```

### POST and GET methods
Define the Post and Get methods to handle the requests coming to API in "/" route path

```
func PostMethod(c *gin.Context) {
  fmt.Println("\napi.go 'PostMethod' called")
  message := "PostMethod called"
  c.JSON(http.StatusOK, message)
}

func GetMethod(c *gin.Context) {
  fmt.Println("\napi.go 'GetMethod' called")
  message := "GetMethod called"
  c.JSON(http.StatusOK, message)
}
```

Please find the entire [source code for API here](https://github.com/chefgs/golang/tree/master/src/api).

### How to build the code
```
go mod init api
go get -d "github.com/gin-gonic/gin"
go build api.go # to create the executable in current path
or
go install api.go # to create the executable in GOBIN path
go run api.go
```
## Unit Testing of Go REST API
- Please refer to the Unit testing code added `api_test.go`
- It uses the Go "testing" module 
- There is mock POST and GET request implemented in the code
- Below command can be used to verify the Unit Testing of REST API, 
`go test`

## Github actions CI flow
- Please refer the Github action workflow file [located here](https://github.com/chefgs/golang/blob/master/.github/workflows/go_api.yml)
- This file shows, how to create Golang CI workflow using Github action yaml file.
- The action workflow file will be invoked, whenever there is code push happens in the repo "golang"

## API testing
- When we run the above code, it runs the REST API on http://localhost:4000/ of the server.

### Method 1
- To verify our REST API, we need to expose the localhost of the server to internet
- So we can use "ngrok" for this purpose
- Download [ngrok here](https://dl.equinox.io/ngrok/ngrok/stable) 
- Extract the ngrok executable in some location on your server.
- Start ngrok on port 4000(Port defined in go API code) as below,

`./ngrok http 4000`
- ngrok generates a dynamic URL. For ex: `http://123er5678.ngrok.io`

- The REST API can be tested by adding the URL in browser address bar, 
`http://123er5678.ngrok.io/` 

- The browser should show,
`GetMethod Called`

### Method 2
- Otherwise, we can use curl command inside same server to verify the endpoint URL, return success.
- In the same server, in which our REST API is running, execute the below commands, to verify the API

- POST method
`curl -X POST http://localhost:4000/`

- The console output should show,
`PostMethod Called`

- GET method
`curl -X GET http://localhost:4000/`

- The console output should show,
`GetMethod Called`

### Method 3
### Docker deployment
### How to run program as docker container
```
docker build -t apigopgm . --rm
docker image ls
docker run -p 4000:4000 --name apicontainer --rm apigopgm
docker container ls
docker stop apicontainer
```

### Method 4
### Create API service using Google Cloud Kubernetes Engine
- Login into Google Cloud Console
- Open Navigation menu from the left and choose `Compute > Kubernetes Engine`
- Choose `Clusters > Create Cluster` and follow instructions to create Cluster
- Next, choose `Workloads > Deploy` to create the container workload and attach it to the Cluster created above
- It is possible to create the container by linking the "github" repository
- Choose the appropriate `Dockerfile` to create image, in this repo, Dockerfile located under the path `src/api/`
- The image created as part of this deployment will be stored into `Google Container Registry (GCR)`
- Once the kube deployment is successful, we then have to expose the Service as `API Endpoint` to consume the REST API functions.
- Choose `expose` option and provide the port number, in which the REST API application has been configured in the code.
- In our case, we have exposed the API at `port 4000`
- After exposing the Kube workload, it generates the `publicIP:4000` for us to utilise the API.

Note: Refer the [Google documentation](https://cloud.google.com/source-repositories/docs/mirroring-a-github-repository) for connecting the Github repo with GCP.

Please reach out to me, in case if more details required.
