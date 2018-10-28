package parser

type Format struct {
	Identifier     string
	Name           string
	AlternateNames []string
}

var DEFAULT_FORMATS = []Format{
	Format{Identifier: "Cam", Name: "Cam", AlternateNames: []string{"CAMRip", "Cam Rip"}},
	Format{Identifier: "Telesync", Name: "Telesync", AlternateNames: []string{"TS", "HDTS", "PDVD", "PreDVDRip"}},
	Format{Identifier: "Workprint", Name: "Workprint", AlternateNames: []string{"WP"}},
	Format{Identifier: "Telecine", Name: "Telecine", AlternateNames: []string{"TC", "HDTC"}},
	Format{Identifier: "PPV", Name: "PayPerView Rip", AlternateNames: []string{"PPV", "PPVRip"}},
	Format{Identifier: "Screener", Name: "Screener", AlternateNames: []string{"SCR", "DVDSCR", "DVDSCREENER", "BDSCR"}},
	Format{Identifier: "DDC", Name: "Digital Distrobution Copy", AlternateNames: []string{"DDC"}},
	Format{Identifier: "R5", Name: "Region 5", AlternateNames: []string{"R5", "R5.LiNE", "R5.AC3.5.1.HQ"}},
	Format{Identifier: "DVDRip", Name: "DVD Rip", AlternateNames: []string{"DVDRip", "DVDMux"}},
	Format{Identifier: "HDTV", Name: "TV Rip", AlternateNames: []string{"DSR", "DSRip", "SATRip", "DTHRip", "DVBRip", "HDTV", "PDTV", "TVRip", "HDTVRip"}},
	Format{Identifier: "VODRip", Name: "VideoOnDemand Rip", AlternateNames: []string{"VODRip", "VODR"}},
	Format{Identifier: "WebDL", Name: "Web Download", AlternateNames: []string{"WEBDL", "Web DL", "Web-DL", "HDRIP", "Web-DLRip"}},
	Format{Identifier: "WebRip", Name: "Web Rip", AlternateNames: []string{"WebRip", "Web-Rip", "Web Rip", "Web"}},
	Format{Identifier: "WebCAP", Name: "Web Cap", AlternateNames: []string{"WebCAP", "Web-Cap", "Web Cap"}},
	Format{Identifier: "BDRip", Name: "Blu-ray Rip", AlternateNames: []string{"Blu-Ray", "BluRay", "BLURAY", "BDRip", "BRRip", "BDMV", "BDR", "BD25", "BD5", "BD5", "BD9"}},
	Format{Identifier: "WS", Name: "Widescreen", AlternateNames: []string{"WS"}},
}

var FORMAT_PARSERS = []formatParser{
	parseFormatName,
	parseFormatAlternateName,
}

type formatParser func(Parseable, []Format) (*Format, error)

func ParseFormat(object Parseable, formats ...Format) ([]Format, error) {
	formats = append(formats, DEFAULT_FORMATS...)
	foundFormats := []Format{}
	for _, parser := range FORMAT_PARSERS {
		format, err := parser(object, formats)
		if err != nil {
			return []Format{}, err
		}
		if format != nil {
			foundFormats = append(foundFormats, *format)
		}
	}
	return foundFormats, nil
}

func parseFormatName(object Parseable, formats []Format) (*Format, error) {
	for _, format := range formats {
		if matchTitle(object.Title(), format.Name) {
			return &format, nil
		}
	}
	return nil, nil
}

func parseFormatAlternateName(object Parseable, formats []Format) (*Format, error) {
	for _, format := range formats {
		for _, alternateName := range format.AlternateNames {
			if matchTitle(object.Title(), alternateName) {
				return &format, nil
			}
		}
	}
	return nil, nil
}
