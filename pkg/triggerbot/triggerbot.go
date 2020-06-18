package triggerbot

import (
	"ask8ball/pkg/lib"
	"strings"

	"github.com/nlopes/slack"
)

type Configuration struct {
	Username  string
	BotName   string
	Keywords  []string
	IconURLs  []string
	WelcomePhrases []string
	HelloPhrases []string
	Phrases   []string
	Reactions map[string]func(ev *slack.MessageEvent) string
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

func (b * Triggerbot) RandomMessage(api *slack.Client, rtm *slack.RTM) error {
	iconURL := lib.PickOne(b.config.IconURLs)
	_, _, err := lib.SendMessageAs(rtm, lib.Message{
		Text:     lib.PickOne(b.config.HelloPhrases),
		Username: b.config.Username,
		IconURL:  iconURL,
		Channel:  "CJ87H4NNA",
	})
	return err
}

func (b * Triggerbot) WelcomeMessage(api *slack.Client, rtm *slack.RTM, user string, channel string) error {
	iconURL := lib.PickOne(b.config.IconURLs)
	_, _, err := lib.SendMessageAs(rtm, lib.Message{
		Text: strings.Replace(lib.PickOne(b.config.WelcomePhrases), "@user", "@" + user, 1),
		Username: b.config.Username,
		IconURL:  iconURL,
		Channel:  channel,
	})
	return err
}

func (b *Triggerbot) HandleMessage(api *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent) error {
	text := strings.ToLower(ev.Msg.Text)
	var sampleText string

	isUpload := ev.Msg.Upload && b.config.OnUpload
	isKeyword := false
	for _, keyword := range b.config.Keywords {
		if strings.Contains(text, keyword) {
			sampleText = lib.PickOne(b.config.Phrases)
			isKeyword = true
			break
		}
	}
	for keyword, reactionFn := range b.config.Reactions {
		if strings.Contains(text, keyword) {
			sampleText = reactionFn(ev)
			isKeyword = true
			break
		}
	}

	if !(isKeyword || isUpload) {
		return nil
	}

	iconURL := lib.PickOne(b.config.IconURLs)

	_, _, err := lib.SendMessageAs(rtm, lib.Message{
		Text:     sampleText,
		Username: b.config.Username,
		IconURL:  iconURL,
		Channel:  ev.Channel,
	})
	return err
}
