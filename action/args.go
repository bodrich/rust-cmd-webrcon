package action

import (
	"errors"
	"flag"
	"math/rand"
	"rust-cmd-webrcon/structures"
)

var Config structures.Configuration
var Message structures.Message

// парсим аргументы
func ParseArgs() error {
	flag.StringVar(&Config.Host, "host", "", "IP:Port WebRcon Server")
	flag.StringVar(&Config.Password, "password", "", "RCON password")
	flag.StringVar(&Message.Message, "command", "", "RCON command")
	flag.Parse()

	if Config.Host == "" {
		return errors.New("host not specified")
	}

	if Config.Password == "" {
		return errors.New("password not specified")
	}

	if Message.Message == "" {
		return errors.New("command not specified")
	}
	Message.Identifier = rand.Intn(100000000)
	Message.Name = "WebRcon"

	return nil
}
