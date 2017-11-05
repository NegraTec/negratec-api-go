FROM golang:1.8

WORKDIR /go/src/app
COPY . .

# dependencies
RUN go get -v github.com/graphql-go/graphql
RUN go get -v github.com/graphql-go/handler
RUN go get -v github.com/stretchr/testify

EXPOSE 8080
