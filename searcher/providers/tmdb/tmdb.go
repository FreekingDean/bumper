package searcher

import (
	tmdb "github.com/ryanbradynd05/go-tmdb"
)

func NewSearcher()

func tmdbSearch(query string) ([]Result, error) {
	tmdbSrv := tmdb.Init("c18ca146bbfb45902c9abe241d09d6c8")
	options := make(map[string]string)
	results, err := tmdbSrv.SearchMulti(query, options)
	found := []Result{}
	if err != nil {
		return found, err
	}
	for _, result := range results.Results {
		switch media := result.(type) {
		case *tmdb.MultiSearchMovieInfo:
			found = append(found, Result{title: media.Title, id: media.ID, media: media.MediaType, posterPath: media.PosterPath, overview: media.Overview})
		case *tmdb.MultiSearchTvInfo:
			found = append(found, Result{title: media.Name, id: media.ID, media: media.MediaType, posterPath: media.PosterPath, overview: media.Overview})
		}
	}
	return found, nil
}

func tvdbSearch(query string) ([]Result, error) {
	tmdb := tmdb.Init("c18ca146bbfb45902c9abe241d09d6c8")
	options := make(map[string]string)
	//options["include_adult"] = "true"
	results, err := tmdb.SearchTv(query, options)
	found := []Result{}
	if err != nil {
		return found, err
	}
	for _, result := range results.Results {
		found = append(found, Result{title: result.Name, id: result.ID})
	}
	return found, nil
}

func tmdbFind(id int) (Result, error) {
	tmdb := tmdb.Init("c18ca146bbfb45902c9abe241d09d6c8")
	movie, err := tmdb.GetMovieInfo(id, make(map[string]string))
	if err != nil {
		return Result{}, err
	}
	return Result{
		id:    movie.ID,
		title: movie.Title,
		media: "movie",
	}, nil
}

func tvdbFind(id int) (Result, error) {
	tmdb := tmdb.Init("c18ca146bbfb45902c9abe241d09d6c8")
	tv, err := tmdb.GetTvInfo(id, make(map[string]string))
	if err != nil {
		return Result{}, err
	}
	return Result{
		id:    tv.ID,
		title: tv.Name,
		media: "tv",
	}, nil
}
