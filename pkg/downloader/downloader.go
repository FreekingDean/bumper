package downloader

type Downloader interface {
	GetAllDownloads() ([]Download, error)
	//StartDownload(Downloadable) (Download, error)
}

type Download interface {
	Progress() float64
	Remove() error
	Delete() error
	Stop() error
}

//type Downloadable interface {
//}
