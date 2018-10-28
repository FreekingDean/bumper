package api

import (
	"encoding/json"
	"net/http"

	"github.com/FreekingDean/bumper/pkg/searcher"
)

type searchResponse struct {
	Media []media `json:"media"`
}

type media struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Media      string `json:"media"`
	PosterPath string `json:"poster_path"`
	Overview   string `json:"overview"`
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	mediaType := r.URL.Query().Get("media")
	query := r.URL.Query().Get("q")
	results, err := searcher.Search(mediaType, query)
	if err != nil {
		respondErr(err, w)
		return
	}
	medias := []media{}
	for _, result := range results {
		newMedia := media{
			Title:      result.Title(),
			ID:         result.SourceID(),
			Media:      result.Media(),
			PosterPath: result.PosterPath(),
			Overview:   result.Overview(),
		}
		medias = append(medias, newMedia)
	}
	response := searchResponse{
		Media: medias,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		respondErr(err, w)
		return
	}
}
