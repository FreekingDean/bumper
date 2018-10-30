// Package configuration allows other packages to build a configuration
// using many different systems (flags, files, storage systems, etc)
package configuration

const (
	// TopLevelKey indicates it has no parent
	TopLevelKey = "-"
)

// Configuration is the domain object that holds all loaded configuration
// information
type configuration struct {
	Values map[string]Value
}

type Value struct {
	Data   interface{}
	Source Source
}

// Service holds all the nesecary information to use configuration
type Service struct {
	global map[string]Configuration

	self Configuration
}

// Source allows configuration to load conifguration
type source interface {
	StoreConfiguration(key string, value string) error
}

// Init creates a new configuration service as well as
// adds options for itself.
func Init() *Service {
	self := config.AddConfiguration(NewConfiguration(config, "configuration", nil))
	self.AddOption()

	svc := &service{
		global: map[string]Configuration{},
	}
	svc.AddConfiguration
}

// AddConfiguration adds a new configuration to the parent
func (service *Service) AddConfiguration(key string) *configuration {
	newConfig := &configuration{
		values: make(map[string]Value),
	}
	service.global[key] = newConfig
	return newConfig
}

func (config *configuration) AddOption(key string, defaultValue interface{}) {
}
