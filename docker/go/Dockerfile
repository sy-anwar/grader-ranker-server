FROM golang:latest

WORKDIR /go/src/github.com/sy-anwar/grader-ranker-server

COPY . /go/src/github.com/sy-anwar/grader-ranker-server

RUN go mod download

EXPOSE 8080

RUN go build -o bin/main cmd/api/main.go
CMD [ "./bin/main" ]
