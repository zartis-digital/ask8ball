package lib

import (
	"errors"
	"math/rand"
	"time"

	"github.com/nlopes/slack"
)

var ErrStop = errors.New("stop")

type Plugin interface {
	Name() string
	HandleMessage(api *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent) error
}

func PickOne(list []string) string {
	return list[rand.Intn(len(list))]
}

type Message struct {
	Text     string
	Username string
	IconURL  string
	Channel  string
}

func SendMessageAs(rtm *slack.RTM, msg Message) (channelID, timestamp string, err error) {
	postParams := slack.PostMessageParameters{
		Username: msg.Username,
		AsUser:   false,
		IconURL:  msg.IconURL,
	}

	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	return rtm.PostMessage(msg.Channel,
		slack.MsgOptionText(msg.Text, false),
		slack.MsgOptionAsUser(false),
		slack.MsgOptionParse(true),
		slack.MsgOptionEnableLinkUnfurl(),
		slack.MsgOptionPostMessageParameters(postParams))
}
