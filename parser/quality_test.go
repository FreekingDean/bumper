package parser_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/FreekingDean/bumper/api/parser"
)

type Titleable string

func (t Titleable) Title() string {
	return string(t)
}

func TestParseResolution(t *testing.T) {
	res100 := parser.Resolution{Identifier: "100", Name: "100K", Resolution: 100, AlternateNames: []string{"100k"}}
	resolutions := []parser.Resolution{
		res100,
		{Identifier: "200", Name: "200K", Resolution: 200},
	}
	t.Run("Without a quality in the title", func(t *testing.T) {
		titles := []Titleable{
			"Title of Movie",
			"Title of Movie100K",
			"Title of Movie100",
		}
		for _, title := range titles {
			resolution, err := parser.ParseResolution(title, resolutions...)
			assert.Nil(t, err)
			assert.Equal(t, parser.Resolution{}, resolution)
		}
	})
	t.Run("With a simple quality in the title", func(t *testing.T) {
		titles := []Titleable{
			"Title of Movie.100K",
			"Title of Movie.100k",
			"Title of Movie[100K]",
			"Title of Movie[100k]",
			"Title of Movie 100K",
			"Title of Movie 100k",
		}
		for _, title := range titles {
			resolution, err := parser.ParseResolution(title, resolutions...)
			assert.Nil(t, err)
			assert.Equal(t, res100, resolution, string(title))
		}
	})
	t.Run("With the resolution in the title", func(t *testing.T) {
		titles := []Titleable{
			"Title of Movie.100p",
			"Title of Movie 100P",
		}
		for _, title := range titles {
			resolution, err := parser.ParseResolution(title, resolutions...)
			assert.Nil(t, err)
			assert.Equal(t, res100, resolution, string(title))
		}
	})
}
