package main

import (
  "net/http"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
)

var dbMentoras []Mentora

type Mentora struct {
  ID string `json:"id"`
  Nome string `json:"nome"`
  Email string `json:"email"`
  Github string `json:"github"`
  Twitter string `json:"twitter"`
  CorRaca string `json:"corRaca"`
  Genero string `json:"genero"`
}

func init() {
  mentora1 := Mentora{ID: "a", Nome: "Rosangela", Email: "algo@gmail.com", Github: "github", Twitter: "", CorRaca: "negra", Genero: "feminino"}
  dbMentoras = append(dbMentoras, mentora1)
}

var mentoraType = graphql.NewObject(
  graphql.ObjectConfig{
    Name: "Mentora",
    Fields: graphql.Fields{
      "id": &graphql.Field{
			     Type: graphql.String,
         },
      "nome": &graphql.Field{
			     Type: graphql.String,
         },
      "email": &graphql.Field{
			     Type: graphql.String,
         },
      "github": &graphql.Field{
			     Type: graphql.String,
         },
      "twitter": &graphql.Field{
			     Type: graphql.String,
         },
      "corRaca": &graphql.Field{
			     Type: graphql.String,
         },
      "genero": &graphql.Field{
			     Type: graphql.String,
         }}})

var mentoraQuery = graphql.NewObject(
  graphql.ObjectConfig{
    Name: "MentoraQuery",
    Fields: graphql.Fields{
      "mentoraLista": &graphql.Field{
        Type: graphql.NewList(mentoraType),
        Description: "Lista das mentoras cadastradas.",
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				      return dbMentoras, nil
			  }}}})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    mentoraQuery,
})

func main() {
  h := handler.New(&handler.Config{
    Schema: &schema,
    Pretty: true,
  })

  http.Handle("/negratec", h)

  http.ListenAndServe(":8080", nil)
}
