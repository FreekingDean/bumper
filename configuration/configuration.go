// Package configuration allows other packages to build a configuration
// using many different systems (flags, files, storage systems, etc)
package configuration

const (
	defaultConfigPath = "/etc/bumper/bumper.conf"
)

type configuration struct {
	values map[string]*value
}

type value struct {
	data         interface{}
	defaultValue interface{}
	source       source
}

// Service is the confiugration service that allows for loading
// and retreiving data
type service struct {
	global  map[string]*configuration
	sources source

	self *configuration
}

// Source allows configuration to load conifguration
type loader interface {
	LoadConfiguration()
}

type source struct {
	name   string
	loader loader
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
		values: make(map[string]*value),
	}

	service.global[key] = newConfig
	return newConfig
}

// AddOption adds an option to a configuration
func (config *configuration) AddOption(name string, defaultValue interface{}) {
	if _, ok := config.values[name]; ok {
		config.values[name].defaultValue = defaultValue
	}
	config.values[name] = &value{
		defaultValue: defaultValue,
	}
}
