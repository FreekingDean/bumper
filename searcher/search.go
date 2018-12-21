package searcher

import ()

type Provider interface {
	Name() string
	Type() string
	Query(string) ([]SearchResult, error)
}

type Searcher struct {
	providers []provider
}

func NewSearcher(providers ...Provider) (*Searcher, error) {
	return &Searcher{
		providers: providers,
	}, nil
}

func Search(provider string) ([]Result, error) {
}
