package connect

import (
	"fmt"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const HOST string = "ftp.ncbi.nlm.nih.gov"
const PORT int = 21
const USER string = "anonymous"
const PASS string = ""

func ConnectFtp() (*ftp.ServerConn, func() error, error) {
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

func ConnectSftp() (*sftp.Client, func() error, func() error, error) {
	config := &ssh.ClientConfig{
		User: USER,
		Auth: []ssh.AuthMethod{
			ssh.Password(PASS),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", HOST, PORT), config)
	if err != nil {
		return nil, nil, nil, err
	}
	sshCancel := func() error {
		return conn.Close()
	}
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		sshCancel()
		return nil, nil, nil, err
	}
	sftpCancel := func() error {
		return sftpClient.Close()
	}
	return sftpClient, sshCancel, sftpCancel, nil

}
