package parser

import (
	"fmt"
)

type Resolution struct {
	Identifier     string
	Name           string
	PrimaryName    string
	AlternateNames []string
	Resolution     int
	MaxBitrate     int //bits per second
	MinBitrate     int //bits per second
}

var DEFAULT_RESOLUTIONS = []Resolution{
	Resolution{Identifier: "8K", Name: "8K", PrimaryName: "8K", Resolution: 4320, MaxBitrate: 90000000, MinBitrate: 40000000},
	Resolution{Identifier: "4K", Name: "4K", Resolution: 2160, MaxBitrate: 80000000, MinBitrate: 20000000},
	Resolution{Identifier: "2K", Name: "2K", Resolution: 1440, MaxBitrate: 30000000, MinBitrate: 10000000},
	Resolution{Identifier: "1080", Name: "1080p", Resolution: 1080, MaxBitrate: 20000000, MinBitrate: 5000000},
	Resolution{Identifier: "720", Name: "720p", Resolution: 720, MaxBitrate: 10000000, MinBitrate: 2000000},
}

var RESOLUTION_PARSERS = []resolutionParser{
	parseResolutionName,
	parseResolutionAlternateName,
	parseResolutionResolution,
}

type resolutionParser func(Parseable, []Resolution) (*Resolution, error)

func ParseResolution(object Parseable, resolutions ...Resolution) (Resolution, error) {
	resolutions = append(resolutions, DEFAULT_RESOLUTIONS...)
	for _, parser := range RESOLUTION_PARSERS {
		resolution, err := parser(object, resolutions)
		if err != nil {
			return Resolution{}, err
		}
		if resolution != nil {
			return *resolution, nil
		}
	}
	return Resolution{}, fmt.Errorf("Could not find resolution")
}

func parseResolutionName(object Parseable, resolutions []Resolution) (*Resolution, error) {
	for _, resolution := range resolutions {
		if matchTitle(object.Title(), resolution.Name) {
			return &resolution, nil
		}
	}
	return nil, nil
}

func parseResolutionAlternateName(object Parseable, resolutions []Resolution) (*Resolution, error) {
	for _, resolution := range resolutions {
		for _, alternateName := range resolution.AlternateNames {
			if matchTitle(object.Title(), alternateName) {
				return &resolution, nil
			}
		}
	}
	return nil, nil
}

func parseResolutionResolution(object Parseable, resolutions []Resolution) (*Resolution, error) {
	for _, resolution := range resolutions {
		if matchTitle(object.Title(), fmt.Sprintf("%dP", resolution.Resolution)) {
			return &resolution, nil
		}
		if matchTitle(object.Title(), fmt.Sprintf("%dp", resolution.Resolution)) {
			return &resolution, nil
		}
	}
	return nil, nil
}
