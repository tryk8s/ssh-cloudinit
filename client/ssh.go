package client

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func executeCmds(commands []string, conf *Config) (err error) {
	config := &ssh.ClientConfig{
		User: conf.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(conf.Password),
		},
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", conf.Hostname, conf.Port), config)
	if err != nil {
		return
	}
	for _, command := range commands {
		session, err := conn.NewSession()
		if err != nil {
			return err
		}
		defer session.Close()

		session.Stdout = conf.Stdout

		err = session.Run(command)
		if err != nil {
			return err
		}
	}

	return
}
