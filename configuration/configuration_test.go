package configuration_test

import (
	"testing"

	"github.com/FreekingDean/bumper/configuration"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	_ = configuration.Init()
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
			s := configuration.Init()
			key1 := s.AddConfiguration(test.KeyNames[0])
			key2 := s.AddConfiguration(test.KeyNames[1])
			assert.True(t, (key1 == key2) == test.Equal)
		})
	}
}

func TestConfigurationAddOption(t *testing.T) {
	s := configuration.Init()
	c := s.AddConfiguration("test")
	c.AddOption("test", "t")
}

func TestServiceAddSource(t *testing.T) {
}
