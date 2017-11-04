package main

import (
  "net/http"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "app/github.com/negratec/queries"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries.MentoraQuery,
})

func main() {
  h := handler.New(&handler.Config{
    Schema: &schema,
    Pretty: true,
  })

  http.Handle("/negratec", h)

  http.ListenAndServe(":8080", nil)
}
