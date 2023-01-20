FROM golang:1.19.5-alpine3.17


# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /root/go/src/shorten_url

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Export necessary port
EXPOSE 8086

CMD ["url_shortener"]