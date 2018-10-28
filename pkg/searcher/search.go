package searcher

import (
	"fmt"
	"strconv"
)

type Result struct {
	title      string
	id         int
	media      string
	posterPath string
	overview   string
}

func (r Result) PosterPath() string {
	return r.posterPath
}

func (r Result) Media() string {
	return r.media
}

func (r Result) Title() string {
	return r.title
}

func (r Result) SourceID() string {
	return strconv.Itoa(r.id)
}

func (r Result) Overview() string {
	return r.overview
}

func (r Result) Source() string {
	return "tmdb"
}

func Search(media, query string) ([]Result, error) {
	return tmdbSearch(query)
}

func GetMedia(_ string, id string, media string) (Result, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return Result{}, err
	}
	if media == "movie" {
		return tmdbFind(idInt)
	} else if media == "tv" {
		return tvdbFind(idInt)
	}
	return Result{}, fmt.Errorf("WOMP!")
}

func GetByIMDBID(_ string) (Result, error) {
	return Result{}, nil
}
