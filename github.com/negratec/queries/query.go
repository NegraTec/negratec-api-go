package queries

import (
  "github.com/graphql-go/graphql"
  "app/github.com/negratec/types"
)

var MentoraQuery *graphql.Object

func init() {
MentoraQuery = graphql.NewObject(
 graphql.ObjectConfig{
   Name: "MentoraQuery",
   Fields: graphql.Fields{
     "mentoraLista": &graphql.Field{
       Type: graphql.NewList(types.MentoraType),
       Description: "Lista das mentoras cadastradas.",
       Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				      return types.DbMentoras, nil
			  }}}})
}
