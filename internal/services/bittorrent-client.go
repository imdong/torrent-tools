package services

type (
	IBitTorrentClient interface{}

	BitTorrentClient struct {
		client IBitTorrentClient
	}
)

func NewBitTorrentClient(client IBitTorrentClient) BitTorrentClient {
	return BitTorrentClient{
		client: client,
	}
}
