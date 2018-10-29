package main

import (
	"github.com/FreekingDean/bumper/config"
	"github.com/spf13/pflag"
)

func main() {
	configPath := pflag.String("config", "$HOME/.bumper", "--config /path/to/config")
	config := config.Load(configPath)
	//load-config
	//start-daemon
	//start-api
}
