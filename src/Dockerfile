# Golang Server's Dockerfile
FROM golang:1.15.6-alpine


RUN apk add --no-cache git

COPY . . 

ADD . . 

RUN go get github.com/lib/pq

RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]