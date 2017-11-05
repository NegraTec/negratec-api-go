package main

import (
  "github.com/graphql-go/graphql"
  "github.com/stretchr/testify/assert"
  "testing"
  "queries"
)

var (dbMentoras []interface{})

func TestMentorasLista(t *testing.T) {
  mentora1 := map[string]interface{}{"id": "a"}
  dbMentoras = append(dbMentoras, mentora1)

  var expected = &graphql.Result{
    Data: map[string]interface{}{
      "mentoraLista": dbMentoras}}

  schema, _ := graphql.NewSchema(graphql.SchemaConfig{
  	Query:    queries.MentoraQuery,
  })

  query := "{mentoraLista{id}}"

  result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}
	assert.Equal(t, expected.Data, result.Data)
}
