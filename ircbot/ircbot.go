package ircbot

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// BotFunc is a function that can be plugged in to handle commands.
type BotFunc func([]string) string

// Bot is a struct representing a bot.
type Bot struct {
	// Nick used by the bot.
	Nick string

	// ChannelNames is slice of names of channels that bot
	// is going to log in.
	ChannelNames []string
	// CommandPrefix is prefix used before command, e.g "~".
	CommandPrefix string

	conn net.Conn

	// Commands maps names of commands to bot functions.
	Commands map[string]BotFunc
}

// New creates new bot.
func New(nick string) *Bot {
	channelNames := []string{"#carneliantest", "#carneliantest2"}
	commandPrefix := ">>"
	commands := make(map[string]BotFunc)
	commands["echo"] = echo
	commands["count"] = count
	return &Bot{
		Nick:          nick,
		ChannelNames:  channelNames,
		CommandPrefix: commandPrefix,
		Commands:      commands,
	}
}

// Connect connects bot to irc server and joins channels.
func (b *Bot) Connect() {
	conn, err := net.Dial("tcp", "irc.freenode.net:6667")
	b.conn = conn
	if err != nil {
		panic(err)
	}
	message := fmt.Sprintf("USER %s %s %s : bot loaded!\n", b.Nick, b.Nick, b.Nick)

	fmt.Fprintf(b.conn, message)
	fmt.Fprintf(b.conn, "NICK "+b.Nick+"\n")

	for _, chanName := range b.ChannelNames {
		fmt.Fprintf(b.conn, "JOIN "+chanName+" \n")
		fmt.Println("joining... " + chanName)
	}
}

// ReadAndRespond is a main loop that handles messages from the
// bot's connection.
func (b *Bot) ReadAndRespond() {
	reader := bufio.NewReader(b.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Printf(message)

		if strings.HasPrefix(message, "PING") {
			b.Pong()
			continue
		}
		b.HandleCommand(message)
	}
}

func (b *Bot) Pong() {
	_, err := b.conn.Write([]byte("PONG :pingis\n"))
	if err != nil {
		panic(err)
	}
}

// HandleCommand gets a message as text, parses and handles it.
// TODO(att): extract parsing as separate method and test it!
func (b *Bot) HandleCommand(message string) {
	byColon := strings.Split(message, ":")
	if len(byColon) != 3 {
		return
	}

	message = byColon[2]

	if strings.HasPrefix(message, b.CommandPrefix) {
		afterPrefix := message[len(b.CommandPrefix):]
		words := strings.Split(afterPrefix, " ")
		command := words[0]
		args := words[1:]
		chann := strings.Split(byColon[1], " ")[2]
		if commandFunc, ok := b.Commands[command]; ok {
			result := commandFunc(args)
			_, err := b.conn.Write([]byte("PRIVMSG " + chann + " :" + result + "\n"))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// TODO(att): Move functions to a function registry.
func echo(args []string) string {
	return strings.Join(args, " ")
}

func count(args []string) string {
	return fmt.Sprintf("%d", len(args))
}
