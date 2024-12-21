package connect

import (
	"fmt"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/secsy/goftp"
)

func ConnectGoFtp() (*ftp.ServerConn, func() error, error) {
	// FTPクライアントの設定
	config := goftp.Config{
		User:     USER,
		Password: PASS,
	}

	// goftpを用いたclientを開発しよう
	_, err := goftp.DialConfig(config, HOST)
	c, err := ftp.Dial(fmt.Sprintf("%s:%d", HOST, PORT), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, nil, err
	}
	cancel := func() error {
		return c.Quit()
	}
	err = c.Login(USER, PASS)
	if err != nil {
		cancel()
		return nil, cancel, err
	}
	return c, cancel, nil
}
