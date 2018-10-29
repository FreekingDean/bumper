package transmission

import (
	"fmt"

	//"github.com/FreekingDean/bumper/api/searcher/torrent"
	"github.com/odwrtw/transmission"
)

type Client struct {
	config transmission.Config
	client *transmission.Client
}

func NewClient(host string, port int, username string, password string) (*Client, error) {
	config := transmission.Config{
		Address:  fmt.Sprintf("http://%s:%d/transmission/rpc", host, port),
		User:     username,
		Password: password,
	}
	client, err := transmission.New(config)
	if err != nil {
		return nil, err
	}
	return &Client{
		config: config,
		client: client,
	}, nil
}

func (c *Client) GetAllDownloads() ([]Download, error) {
	downloads := make([]Download, 0)
	torrents, err := c.client.GetTorrents()
	if err != nil {
		return downloads, err
	}
	for _, torrent := range torrents {
		downloads = append(downloads, Download{torrent: torrent})
	}
	return downloads, nil
}

//func (c *TransmissionClient) StartDownload(torrent torrent.Torrent) (TransmissionDownload
