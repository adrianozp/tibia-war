package twitch

import (
	l "github.com/adrianozp/tibia-war/internal/core/logger"
	irc "github.com/fluffle/goirc/client"
	"fmt"
	//"strings"
)

var ircCfg IrcConfig
var c IrcClient

type IrcConfig struct {
	Name		string
	Token		string
	Channels	[]string
}

type IrcClient struct {
	client *irc.Conn
	quitCh chan bool
}

func sendMessage(channel string, text string) {
	//l.Log().Debug().Msg(fmt.Sprintf("Sending message: %s: %s", channel, text))
	fmt.Printf("%s: Sending message: %s", channel, text)
	c.client.Privmsg(channel, text)
}

func onConnect(conn *irc.Conn, line *irc.Line) {
	fmt.Printf("onConnect\n")
	for _, channel := range ircCfg.Channels {
		//log.Debug().Msg(fmt.Sprintf("Joining %s channel", channel))
		c.client.Join(fmt.Sprintf("#%s", channel))
	}
	//log.Info().Msg("IRC Client Connected")
	sendMessage("#pantibiabot", "teste")
}

func onDisconnect(conn *irc.Conn, line *irc.Line) {
	fmt.Printf("onDisconnect\n")
	c.quitCh <- true 
}

func onNotice(conn *irc.Conn, line *irc.Line) {
	fmt.Printf("NOTICE: %v\n", line)
}

func onMsg(conn *irc.Conn, line *irc.Line) {
	fmt.Printf("MSG: %v\n", line)
	sendMessage("#panzp", line.Text())
}

func Init(i IrcConfig) {
	cfg := irc.NewConfig(i.Name)
	
	cfg.Me.Name = i.Name
	cfg.Me.Ident = i.Name
	cfg.Server = "irc.chat.twitch.tv:6667"
	cfg.Pass = fmt.Sprintf("oauth:%s", i.Token)

	//log.Debug().Msg(fmt.Sprintf("Setting IRC client for %s with token %s", *twitchName, *oauthToken))
	c.client = irc.Client(cfg)

	c.client.HandleFunc(irc.CONNECTED, onConnect)
	c.client.HandleFunc(irc.DISCONNECTED, onDisconnect)
	c.client.HandleFunc(irc.NOTICE, onNotice)
	c.client.HandleFunc(irc.PRIVMSG, onMsg)

	if err := c.client.Connect(); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}

	<-c.quitCh
}

func Config() IrcConfig {
	return ircCfg
}

func Client() IrcClient {
	return c
}