package connect

import (
	"github.com/secsy/goftp"
)

func ConnectGoFtp() (*goftp.Client, func(), error) {
	// FTPクライアントの設定
	config := goftp.Config{
		User:     USER,
		Password: PASS,
	}
	client, err := goftp.DialConfig(config, HOST)
	if err != nil {
		return nil, nil, err
	}

	return client, func() { _ = client.Close() }, nil
}
