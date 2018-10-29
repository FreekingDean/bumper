package transmission

import (
	"github.com/odwrtw/transmission"
)

type Download struct {
	torrent *transmission.Torrent
}

func (d Download) Progress() float64 {
	//d.torrent.Update()
	return 1 - (float64(d.torrent.LeftUntilDone) / float64(d.torrent.TotalSize))
}

func (d Download) Name() string {
	return d.torrent.Name
}
