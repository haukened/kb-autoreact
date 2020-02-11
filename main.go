package main

import (
	"flag"
	"fmt"
	"os"

	"samhofi.us/x/keybase"
)

var (
	k    = keybase.NewKeybase()
	who  string
	what string
)

func init() {
	flag.StringVar(&who, "username", "", "username to react to")
	flag.StringVar(&what, "reaction", "", "reaction, with colons if its an emoji")
	flag.Parse()
	if who == "" {
		fmt.Println("missing -username flag")
		os.Exit(1)
	}
	if what == "" {
		fmt.Println("missing -reaction flag")
		os.Exit(2)
	}
}

func main() {
	fmt.Println("running!")
	k.Run(handler)
}

// handler handles each incoming message.
func handler(m keybase.ChatAPI) {
	if m.ErrorListen != nil {
		fmt.Println(fmt.Sprintf("Error: %s", *m.ErrorListen))
		return
	}

	if m.Msg.Sender.Username == who {
		c := k.NewChat(m.Msg.Channel)
		c.React(m.Msg.ID, what)
	}
}
