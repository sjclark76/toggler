// /Users/stuartclark/GolandProjects/toggler/main.go
package main

import (
	"log"
	"net/http"

	// Ensure your module name in go.mod is 'toggler' or adjust import paths.
	"toggler/data"        // Your data access package
	gql "toggler/graphql" // Your GraphQL package (aliased for clarity)

	"github.com/graphql-go/handler"
)

func main() {
	// 1. Initialize your data store.
	featureStore := data.NewFeatureFlagsStore()

	// 2. Create the GraphQL schema, injecting the data store.
	schema, err := gql.NewSchema(featureStore)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// 3. Set up the GraphQL HTTP handler.
	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", gqlHandler)

	// 4. Start the server.
	port := ":8080"
	log.Printf("GraphQL feature flag service starting on http://localhost%s/graphql", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
