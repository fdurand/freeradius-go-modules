package modules

import (
	"fmt"

	"github.com/fdurand/freeradius-go"
)

// Creater function
type Creater func(string) (freeradius.Module, error)

var moduleLookup = map[string]Creater{
	"cache_eap": NewCacheEap,
}

// Create function
func Create(moduleType string, name string) (freeradius.Module, error) {
	if creater, found := moduleLookup[moduleType]; found {
		return creater(name)
	}

	return nil, fmt.Errorf("Module of %s not found", moduleType)
}
