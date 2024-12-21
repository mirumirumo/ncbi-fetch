package connect

import (
	"fmt"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

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
