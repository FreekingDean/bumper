package api

import (
	"encoding/json"
	"net/http"

	"github.com/FreekingDean/bumper/pkg/subscriber"
)

type subscribeMediaRequest struct {
	IMDBID string `json:"imdb_id"`
}

type subscribeMediaResponse struct {
	Ok bool `json:"ok"`
}

func handleSubscribe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	request := &subscribeMediaRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		respondErr(err, w)
		return
	}
	s := subscriber.New(db)
	err = s.Subscribe(request.IMDBID)
	if err != nil {
		respondErr(err, w)
		return
	}
	response := subscribeMediaResponse{Ok: true}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		respondErr(err, w)
		return
	}
}

type subscriptionsResponse struct {
	Subscriptions []subscriber.Subscription `json:"subscriptions"`
}

type subscription struct {
	ID               string        `json:"id"`
	IMDBID           string        `json:"imdb_id"`
	Active           bool          `json:"active"`
	TargetResolution resolution    `json:"target_resolution"`
	Downloads        []download    `json:"download"`
	DiskVersions     []diskVersion `json:"disk_versions"`
}

type download struct {
	Path string `json:"path"`
}

type diskVersion struct {
	Resolution resolution `json:"resolution"`
	Formats    []format   `json:"formats"`
	Path       string     `json:"path"`
}

type resolution struct {
	Name string `json:"name"`
}

type format struct {
	Name string `json:"name"`
}

func handleGetSubscribtions(w http.ResponseWriter, r *http.Request) {
	serv := subscriber.New(db)
	subscriptions, err := serv.AllSubscriptions()
	if err != nil {
		respondErr(err, w)
		return
	}
	err = encodeSubscriptions(w, subscriptions)
	response := subscriptionsResponse{Subscriptions: subscriptions}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		respondErr(err, w)
		return
	}
	w.WriteHeader(200)
}
