package subscriber

type SubscribeService struct {
	storer   db
	searcher Searcher
}

type Searcher interface {
	GetByIMDBID(string) (Subscribeable, error)
}

type Subscribeable interface{}

func New(s Storer) *SubscribeService {
	return &SubscribeService{
		storer: s,
	}
}
