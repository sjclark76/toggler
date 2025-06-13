// /Users/stuartclark/GolandProjects/toggler/graphql/schema.go
package graphql

import (
	// Ensure your module name in go.mod is 'toggler' or adjust the import path.
	"toggler/data" // Import your data package

	"github.com/graphql-go/graphql"
)

// NewSchema creates and returns the GraphQL schema.
// It requires a FeatureFlagsStore to wire up resolvers.
func NewSchema(store *data.FeatureFlagsStore) (graphql.Schema, error) {
	// Instantiate your resolver, injecting the data store.
	resolver := &Resolver{Store: store}

	// Define the Root Query Type using types and resolvers from this package.
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"featureFlag": &graphql.Field{
					Type:        FeatureFlagType, // Defined in types.go
					Description: "Get a specific feature flag by name.",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type:        graphql.NewNonNull(graphql.String),
							Description: "Name of the feature flag to retrieve.",
						},
					},
					Resolve: resolver.ResolveFeatureFlag, // Method from resolvers.go
				},
				"allFeatureFlags": &graphql.Field{
					Type:        graphql.NewList(FeatureFlagType), // Defined in types.go
					Description: "Get all feature flags.",
					Resolve:     resolver.ResolveAllFeatureFlags, // Method from resolvers.go
				},
			},
		},
	)

	// Create the schema configuration.
	schemaConfig := graphql.SchemaConfig{
		Query: queryType,
	}
	return graphql.NewSchema(schemaConfig)
}
