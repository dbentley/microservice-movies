FROM golang:1.10

ADD . /go/src/github.com/dbentley/microservice-movies/services/ratings
RUN go install github.com/dbentley/microservice-movies/services/ratings

ENTRYPOINT /go/bin/ratings