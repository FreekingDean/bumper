package subscriber

import (
	"github.com/FreekingDean/bumper/pkg/searcher"
)

// Subscribe will create a new bumper subscription
// based on an IMDB ID
func (serv SubscribeService) Subscribe(imdbID string) error {
	_, err := searcher.GetByIMDBID(imdbID)
	if err != nil {
		return err
	}

	subscription := &Subscription{
		imdbID: imdbID,
	}

	return subscription.create(serv.storer)
}

func (serv SubscribeService) AllSubscriptions() ([]Subscription, error) {
	subscriptions := make([]Subscription, 0)
	serv.storer.Get("SELECT * FROM subscriptions", &
}
