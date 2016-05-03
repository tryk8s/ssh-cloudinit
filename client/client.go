package client

import "fmt"

func Run(conf *Config) error {
	err := executeCmds(conf.GetCmds(), conf)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	return nil
}
