package connect

import (
	"fmt"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func ConnectSftp() (*sftp.Client, func(), error) {
	config := &ssh.ClientConfig{
		User: USER,
		Auth: []ssh.AuthMethod{
			ssh.Password(PASS),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", HOST, PORT), config)
	if err != nil {
		return nil, nil, err
	}
	sshCancel := func() {
		_ = conn.Close()
	}
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		sshCancel()
		return nil, nil, err
	}
	sftpCancel := func() {
		_ = sftpClient.Close()
		_ = conn.Close()
	}
	return sftpClient, sftpCancel, nil

}
