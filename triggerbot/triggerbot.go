package triggerbot

import (
	"strings"
	"nyanbot"
	"github.com/nlopes/slack"
)

type Configuration struct {
	Username  string
	BotName   string
	Keywords  []string
	IconURLs  []string
	Phrases   []string
	Reactions map[string]func() string
	OnUpload  bool
}

type Triggerbot struct {
	config Configuration
}

func New(config Configuration) *Triggerbot {
	return &Triggerbot{
		config: config,
	}
}

func (b *Triggerbot) Name() string {
	return b.config.Username
}

func (b *Triggerbot) HandleMessage(api *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent) error {
	text := strings.ToLower(ev.Msg.Text)
	var sampleText string

	isUpload := ev.Msg.Upload && b.config.OnUpload
	isKeyword := false
	for _, keyword := range b.config.Keywords {
		if strings.Contains(text, keyword) {
			sampleText = nyanbot.PickOne(b.config.Phrases)
			isKeyword = true
			break
		}
	}
	for keyword, reactionFn := range b.config.Reactions {
		if strings.Contains(text, keyword) {
			sampleText = reactionFn()
			isKeyword = true
			break
		}
	}

	if !(isKeyword || isUpload) {
		return nil
	}

	iconURL := nyanbot.PickOne(b.config.IconURLs)

	_, _, err := nyanbot.SendMessageAs(rtm, nyanbot.Message{
		Text:     sampleText,
		Username: b.config.Username,
		IconURL:  iconURL,
		Channel:  ev.Channel,
	})
	return err
}
