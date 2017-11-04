FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get app/github.com/negratec

EXPOSE 8080
