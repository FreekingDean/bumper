package main

import (
	"github.com/FreekingDean/bumper/configuration"
	"github.com/FreekingDean/bumper/database"
	"github.com/spf13/pflag"
)

// Four step approach into running the services:
// Init -> Configure -> Start -> Stop
// This allows us to allow packages to build their
// own configuration options.
func main() {
	config := configuration.DefaultConfig()
	config.AddSource(configuration.DefaultFlagSource)
	config.AddSource(configuration.DefaultFileSource)

	sotrageConfig := config.AddConfiguration("storage")
	store := storage.Init(storageConfig)
	config.AddSource(store)
	config.Init()

	err := config.Configure()
	err := store.Configure()
	//start-daemon
	//start-api
}

func printUsage() {
	fmt.Println("usage")
}
