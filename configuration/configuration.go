// Package configuration allows other packages to build a configuration
// using many different systems (flags, files, storage systems, etc)
package configuration

const (
	defaultConfigPath = "/etc/bumper/bumper.conf"
)

// Configuration is the domain object that holds all loaded configuration
// information
type Configuration struct {
	Values map[string]*Value
}

// Value holds the type of data in memory
type Value struct {
	Data    interface{}
	Default interface{}
	Source  Source
}

// Service is the confiugration service that allows for loading
// and retreiving data
type Service struct {
	Global map[string]*Configuration

	Self *Configuration
}

// Source allows configuration to load conifguration
type Source interface {
	StoreConfiguration(key string, value string) error
}

// Init creates a new configuration service as well as
// adds options for itself.
func Init() *Service {
	svc := &Service{
		Global: map[string]*Configuration{},
	}

	self := svc.AddConfiguration("config")
	svc.Self = self

	self.AddOption("path", defaultConfigPath)
	return svc
}

// AddConfiguration adds a new configuration to the parent
func (service *Service) AddConfiguration(key string) *Configuration {
	if config, ok := service.Global[key]; ok {
		return config
	}

	newConfig := &Configuration{
		Values: make(map[string]*Value),
	}

	service.Global[key] = newConfig
	return newConfig
}

// AddOption adds an option to a configuration
func (config *Configuration) AddOption(name string, defaultValue interface{}) {
	if _, ok := config.Values[name]; ok {
		config.Values[name].Default = defaultValue
	}
	config.Values[name] = &Value{
		Default: defaultValue,
	}
}
