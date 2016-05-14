package main

import (
	"flag"
	"fmt"
	"github.com/tryk8s/ssh-cloudinit/client"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

var remote = flag.String("remote", "", "Remote cloud-init url")
var osType = flag.String("os", "ubuntu", "Server OS")
var port = flag.Int("port", 22, "Server SSH port")
var user = flag.String("user", "root", "Server SSH user")

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: ssh-cloudinit [options] <server>\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		return
	}
	fmt.Printf("%s@%s's password: ", *user, args[0])
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Print("\n")

	conf := &client.Config{
		Hostname: args[0],
		User:     *user,
		Password: string(bytePassword),
		Server:   *remote,
		Port:     *port,
		Os:       *osType,
		Stdout:   os.Stdout,
	}
	client.Run(conf)
}
