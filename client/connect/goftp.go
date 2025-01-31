package connect

import (
	"github.com/mirumirumo/ncbi-cli/client/config"
	"github.com/secsy/goftp"
)

func ConnectGoFtp() (*goftp.Client, func(), error) {
	// FTPクライアントの設定
	c, err := config.FtpConfigs()
	if err != nil {
		return nil, nil, err
	}

	config := goftp.Config{
		User:     c.USER,
		Password: c.PASS,
	}
	client, err := goftp.DialConfig(config, HOST)
	if err != nil {
		return nil, nil, err
	}

	return client, func() { _ = client.Close() }, nil
}
