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

## Create REST API using Golang

```
go get -d "github.com/gin-gonic/gin"
go mod init api
go install api.go
go run api.go
```
## API testing
- Download and Start ngrok on port 4000
