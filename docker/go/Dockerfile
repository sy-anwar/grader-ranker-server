FROM golang:latest

WORKDIR /go/src/github.com/sy-anwar/grader-ranker-server

COPY . /go/src/github.com/sy-anwar/grader-ranker-server

RUN go mod download

EXPOSE 8080
