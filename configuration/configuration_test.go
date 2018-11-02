package configuration_test

import (
	"testing"

	"github.com/FreekingDean/bumper/configuration"
	"github.com/stretchr/testify/assert"
)

type service interface {
	AddConfiguration(string) *configuration.Service
}

func TestInit(t *testing.T) {
	config := configuration.Init()
	assert.NotNil(t, config.Global["config"])
	assert.Equal(t, config.Global["config"], config.Self)
}

func TestServiceAddConfiguration(t *testing.T) {
	tests := []struct {
		Name     string
		KeyNames []string
		Equal    bool
	}{
		{"Similar keys should be the same", []string{"key", "key"}, true},
		{"Different keys should not be the same", []string{"key", "key2"}, false},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			s := newService()
			key1 := s.AddConfiguration(test.KeyNames[0])
			key2 := s.AddConfiguration(test.KeyNames[1])
			assert.True(t, (key1 == key2) == test.Equal)
		})
	}
}

func TestConfigurationAddOption(t *testing.T) {
	tests := []struct {
		Name         string
		OptionNames  []string
		TotalOptions int
	}{
		{"Similar keys should have 1 value", []string{"key", "key"}, 1},
		{"Different keys should have 2 values", []string{"key", "key2"}, 2},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := newConfiguration()
			c.AddOption(test.OptionNames[0], test.OptionNames[0])
			c.AddOption(test.OptionNames[1], test.OptionNames[1])
			assert.Equal(t, test.TotalOptions, len(c.Values))
		})
	}
}

func newService() *configuration.Service {
	return configuration.Init()
}

func newConfiguration() *configuration.Configuration {
	return newService().AddConfiguration("test")
}
