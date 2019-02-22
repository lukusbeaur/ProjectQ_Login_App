package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

//User starts the user object
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	//add info for sign up page
}

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{})
	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		results := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(results)

		userType := graphql.NewObject(graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		})

		rootQ := graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Type: graphql.NewList(userType),
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return nil, nil
					},
				},
			},
		})
		fmt.Fprint(w, rootQ) //wtf is going on
	})

	http.ListenAndServe(":8080", nil)

	//testing because i dont fucking get GRAPHQL lets a go..mario

}

//Creating MockData ToDo:
//replace this code with DB garbage.
var user []User = []User{
	User{
		Email:    "test1@email.com",
		Password: "password",
	},
	User{
		Email:    "test2@email.com",
		Password: "password1",
	},
}
