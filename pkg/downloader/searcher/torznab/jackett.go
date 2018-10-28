package torznab

import (
	"fmt"
	"github.com/FreekingDean/go-newznab/newznab"
)

type Client struct {
	client newznab.Client
}

func NewJackettClient(hostname string, port int, apiKey string) *Client {
	return &Client{
		client: newznab.New(fmt.Sprintf("http://%s:%d", hostname, port), "/v2.0/indexers/all/results/torznab", apiKey, 1, false),
	}
}

type Torrentable struct {
	nzb newznab.NZB
}

func (c *Client) Query(query string) ([]Torrentable, error) {
	results := make([]Torrentable, 0)
	nzbs, err := c.client.SearchWithQuery([]int{5000}, query, "search")
	if err != nil {
		return results, err
	}
	for _, nzb := range nzbs {
		results = append(results, Torrentable{nzb: nzb})
	}
	return results, nil
}

func (t Torrentable) Seeders() int {
	return t.nzb.Seeders
}

func (t Torrentable) Name() string {
	return t.nzb.Title
}

func (t Torrentable) Title() string {
	return t.nzb.Title
}

func (t Torrentable) Size() int64 {
	return t.nzb.Size
}
