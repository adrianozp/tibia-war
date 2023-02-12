package twitch

import (
	irc "github.com/fluffle/goirc/client"
	"fmt"
	//"strings"
)

var ircCfg Config
var c Client

type Config struct {
	name		string
	token		string
	channels	[]string
}

type Client struct {
	client *irc.Conn
	quitCh chan bool
}

func sendMessage(channel string, text string) {
	//log.Debug().Msg(fmt.Sprintf("Sending message: %s: %s", channel, text))
	fmt.Printf("%s: Sending message: %s", channel, text)
	c.client.Privmsg(channel, text)
}

func onConnect(conn *irc.Conn, line *irc.Line) {
	fmt.Printf("onConnect\n")
	for _, channel := range ircCfg.channels {
		//log.Debug().Msg(fmt.Sprintf("Joining %s channel", channel))
		c.client.Join(fmt.Sprintf("#%s", channel))
	}
	//log.Info().Msg("IRC Client Connected")
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

func Init() {
	fmt.Printf("INIT\n")
	cfg := irc.NewConfig(ircCfg.name)
	
	cfg.Me.Name = ircCfg.name
	cfg.Me.Ident = ircCfg.name
	cfg.Server = "irc.chat.twitch.tv:6667"
	cfg.Pass = fmt.Sprintf("oauth:%s", ircCfg.token)

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

func main() {
	var cfg Config
	cfg.name = "pantibiabot"
	cfg.token = "TOKEN"
	cfg.channels = []string{"panzp"}
	ircCfg = cfg
	Init()
}