package connect

import (
	"fmt"
	"time"

	"github.com/jlaffaye/ftp"
)

func ConnectFtp() (*ftp.ServerConn, func(), error) {
	// FTPクライアントの設定
	c, err := ftp.Dial(fmt.Sprintf("%s:%d", HOST, PORT), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, nil, err
	}
	cancel := func() {
		_ = c.Quit()
	}
	err = c.Login(USER, PASS)
	if err != nil {
		cancel()
		return nil, nil, err
	}
	return c, cancel, nil
}
