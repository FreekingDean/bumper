package main

import (
	"github.com/FreekingDean/bumper/configuration"
	"github.com/FreekingDean/bumper/database"
	"github.com/spf13/pflag"
)

// Four step approach into running the services:
// Init -> Start -> Stop
// This allows us to allow packages to build their
// own configuration options.
func main() {
	config := configuration.DefaultConfig()
	config.AddSource(configuration.DefaultFlagSource)
	config.AddSource(configuration.DefaultFileSource)

	databaseConfig := config.AddConfiguration("database")
	database := database.NewDatabase(storageConfig)
	config.AddSource(database)

	config.Init()

	err := config.Start()
	err := database.Start()
	//start-daemon
	//start-api
	database.Stop()
	config.Stop()
}

func printUsage() {
	fmt.Println("usage")
}
