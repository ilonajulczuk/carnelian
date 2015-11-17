package main

import (
	"github.com/ilonajulczuk/carnelian/ircbot"
)

func main() {
	nick := "Carnelian"
	bot := ircbot.New(nick)
	bot.Connect()
	bot.ReadAndRespond()
}
