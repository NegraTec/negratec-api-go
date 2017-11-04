# NegraTec API Contribua!

## Utiliza

- Go + Graphql: https://github.com/graphql-go/graphql
- Docker

##  Construindo projeto

constroe a imagem
`docker build -t negratec-go`

abri bash do container
`docker run --rm -p 8080:8080 -v "$PWD":/go/src/app -w/go/src/app -it negratec-go /bin/bash`

sobe o servidor. http://localhost:8080/negratec?query={mentoraLista{github}}
`docker run --rm -p 8080:8080 -v "$PWD":/go/src/app -w /go/src/app -it negratec-go go run github.com/negratec/negratec.go`
