package main

import (
	"flag"
	"fmt"
	"github.com/tryk8s/ssh-cloudinit/client"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

var (
	remote string
	osType string
	port   int
	user   string
)

func init() {
	flag.StringVar(&remote, "remote", "", "Remote cloud-init url")
	flag.StringVar(&osType, "os", "ubuntu", "Server OS")
	flag.IntVar(&port, "port", 22, "Server SSH port")
	flag.StringVar(&user, "user", "root", "Server SSH user")
	flag.Usage = func() {
		fmt.Printf("Usage: ssh-cloudinit [options] <server>\n\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		return
	}
	fmt.Printf("%s@%s's password: ", user, args[0])
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Print("\n")

	conf := &client.Config{
		Hostname: args[0],
		User:     user,
		Password: string(bytePassword),
		Server:   remote,
		Port:     port,
		Os:       osType,
		Stdout:   os.Stdout,
	}
	client.Run(conf)
}
