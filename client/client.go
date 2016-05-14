package client

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func Run(conf *Config) error {
	config := &ssh.ClientConfig{
		User: conf.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(conf.Password),
		},
	}
	server := fmt.Sprintf("%s:%d", conf.Hostname, conf.Port)
	err := executeCmds(server, config, conf.Stdout, conf.GetCmds())
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	return nil
}
