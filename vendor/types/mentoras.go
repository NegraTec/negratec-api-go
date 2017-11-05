package types

import ("github.com/graphql-go/graphql")

var (
 DbMentoras []Mentora
 Schema graphql.Schema
 MentoraType *graphql.Object
)

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
  DbMentoras = append(DbMentoras, mentora1)


MentoraType = graphql.NewObject(
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
}
