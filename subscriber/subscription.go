package subscriber

import (
	"github.com/FreekingDean/bumper/pkg/downloader"
)

// Subscription is the object bumper uses to track
// what to download/update.
type Subscription struct {
	ID     uint   `db:"id" json:"id"`
	IMDBID string `db:"imdb_id" json:"imdb_id"`

	Downloads []downloader.Download `json:"downloads"`
}

func (s *Subscription) create(db Storer) error {
	err := db.Store("INSERT INTO subscriptions (imdb_id) VALUES (?)", s.imdbID)
	if err != nil {
		return err
	}
	return nil
}
