package graphql

import "github.com/graphql-go/graphql"

// FeatureFlagType defines the GraphQL object type for a FeatureFlag.
// It's exported (starts with an uppercase letter) to be accessible by other files in this package.
var FeatureFlagType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FeatureFlag",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the feature flag.",
			},
			"isEnabled": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Boolean),
				Description: "The current state of the feature flag.",
			},
		},
	},
)
