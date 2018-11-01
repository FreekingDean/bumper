// Package configuration allows other packages to build a configuration
// using many different systems (flags, files, storage systems, etc)
package configuration

const (
	defaultConfigPath = "/etc/bumper/bumper.conf"
)

// Configuration is the domain object that holds all loaded configuration
// information
type configuration struct {
	values map[string]value
}

type value struct {
	Data   interface{}
	Source source
}

type source interface{}

type service struct {
	global map[string]*configuration

	self *configuration
}

// Source allows configuration to load conifguration
type source interface {
	StoreConfiguration(key string, value string) error
}

// Init creates a new configuration service as well as
// adds options for itself.
func Init() *service {
	svc := &service{
		global: map[string]*configuration{},
	}

	self := svc.AddConfiguration("config")
	svc.self = self

	self.AddOption("path", defaultConfigPath)
	return svc
}

// AddConfiguration adds a new configuration to the parent
func (service *service) AddConfiguration(key string) *configuration {
	if config, ok := service.global[key]; ok {
		return config
	}

	newConfig := &configuration{
		values: make(map[string]value),
	}

	service.global[key] = newConfig
	return newConfig
}

func (config *configuration) AddOption(key string, defaultValue interface{}) {
}
