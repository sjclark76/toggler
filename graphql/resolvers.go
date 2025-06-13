// /Users/stuartclark/GolandProjects/toggler/graphql/resolvers.go
package graphql

import (
	"fmt"
	// Ensure your module name in go.mod is 'toggler' or adjust the import path.
	"toggler/data" // Import your data package

	"github.com/graphql-go/graphql"
)

// Resolver holds dependencies for GraphQL resolvers, such as the data store.
type Resolver struct {
	Store *data.FeatureFlagsStore
}

// ResolveFeatureFlag handles the logic for the 'featureFlag' query.
func (r *Resolver) ResolveFeatureFlag(params graphql.ResolveParams) (interface{}, error) {
	name, ok := params.Args["name"].(string)
	if !ok || name == "" {
		return nil, fmt.Errorf("argument 'name' is required and must be a string")
	}
	return r.Store.GetFlag(name)
}

// ResolveAllFeatureFlags handles the logic for the 'allFeatureFlags' query.
func (r *Resolver) ResolveAllFeatureFlags(p graphql.ResolveParams) (interface{}, error) {
	return r.Store.GetAllFlags()
}
