package searcher

type Searcher interface {
	Query(string) ([]Downloadable, error)
}
