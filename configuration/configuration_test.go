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
	s := configuration.Init()
	c := s.AddConfiguration("key")
	nc := s.AddConfiguration("key")
	assert.True(t, c == nc)
	nc1 := s.AddConfiguration("key2")
	assert.False(t, c == nc1)
}
