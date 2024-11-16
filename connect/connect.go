package connect

import (
	"time"

	"github.com/jlaffaye/ftp"
)

const SERVER string = "ftp.ncbi.nlm.nih.gov:21"

func Connect() (*ftp.ServerConn, func() error, error) {
	c, err := ftp.Dial(SERVER, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, nil, err
	}
	cancel := func() error {
		return c.Quit()
	}
	err = c.Login("anonymous", "anonymous")
	if err != nil {
		return nil, cancel, err
	}

	return c, cancel, nil
}
