// /Users/stuartclark/GolandProjects/toggler/data/store.go
package data

import "fmt"

// FlagOutput is a structure for returning flag data, suitable for GraphQL resolvers.
type FlagOutput struct {
	Name      string
	IsEnabled bool
}

// FeatureFlagsStore manages the feature flags.
// In a real application, this would interact with a database or configuration files.
type FeatureFlagsStore struct {
	flags map[string]bool
}

// NewFeatureFlagsStore creates and initializes a new FeatureFlagsStore.
func NewFeatureFlagsStore() *FeatureFlagsStore {
	return &FeatureFlagsStore{
		flags: map[string]bool{
			"new-checkout-flow": true,
			"beta-feature-x":    false,
			"dark-mode":         true,
		},
	}
}

// GetFlag retrieves a specific flag for GraphQL.
func (s *FeatureFlagsStore) GetFlag(name string) (*FlagOutput, error) {
	isEnabled, exists := s.flags[name]
	if !exists {
		return nil, fmt.Errorf("feature flag '%s' not found", name)
	}
	return &FlagOutput{Name: name, IsEnabled: isEnabled}, nil
}

// GetAllFlags retrieves all flags for GraphQL.
func (s *FeatureFlagsStore) GetAllFlags() ([]*FlagOutput, error) {
	var result []*FlagOutput
	for name, isEnabled := range s.flags {
		result = append(result, &FlagOutput{Name: name, IsEnabled: isEnabled})
	}
	return result, nil
}
