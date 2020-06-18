package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/nlopes/slack"

	"ask8ball/pkg/lib"
	"ask8ball/pkg/personas"
	"ask8ball/pkg/triggerbot"
)

var (
	ask8ballConfiguration = triggerbot.Configuration{
		Username:  "Ask8Ball",
		Keywords:  personas.Ask8BallKeywords,
		IconURLs:  personas.Ask8BallIconURLs,
		HelloPhrases: personas.Ask8BallHelloPhrases,
		Phrases:   personas.Ask8BallPhrases,
		Reactions: personas.Ask8BallReactions,
	}
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if len(token) == 0 {
		log.Fatalln("SLACK_TOKEN env var is required!")
	} else {
		log.Print(token)
	}

	plugins := []lib.Plugin{
		triggerbot.New(ask8ballConfiguration),
	}

	api := slack.New(token, slack.OptionDebug(true))

	allowedChannels, err := parseGroups(api)
	panicIf(err)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.InvalidAuthEvent:
			log.Fatal("Invalid credentials")
			return
		case *slack.LatencyReport:
			log.Printf("LatencyReport: %v\n", ev.Value)
		case *slack.RTMError:
			log.Printf("Error: %s\n", ev.Error())
		case *slack.HelloEvent:
			for _, plugin := range plugins {
				if rand.Intn(5) > 3 {
					err := plugin.RandomMessage(api, rtm)
					if err == lib.ErrStop {
						break
					}
					if err != nil {
						log.Println("MessageEvent", plugin.Name(), err)
					}
				}
			}
		case *slack.MemberJoinedChannelEvent:
			for _, plugin := range plugins {
				err := plugin.WelcomeMessage(api, rtm, ev.User, ev.Channel)
				if err == lib.ErrStop {
					break
				}
				if err != nil {
					log.Println("MessageEvent", plugin.Name(), err)
				}
			}
		case *slack.ConnectingEvent:
			// Can't use api.SetUserCustomStatus for bots :(
		case *slack.ConnectedEvent:
			log.Printf("Connected! ConnectionCount=%d", ev.ConnectionCount)
		case *slack.MessageEvent:
			// log.Printf("Message: %+v\n", ev)
			if ev.Msg.Text == "stop-ask8ball" {
				os.Exit(1)
			}

			isInAllowedChannels := allowedChannels[ev.Channel]
			if !isInAllowedChannels {
				log.Printf("Channel not allowed %s %+v", ev.Channel, ev)
				continue
			}

			for _, plugin := range plugins {
				if ev.Msg.Username == plugin.Name() {
					continue
				}
				err := plugin.HandleMessage(api, rtm, ev)
				if err == lib.ErrStop {
					break
				}
				if err != nil {
					log.Println("MessageEvent", plugin.Name(), err)
				}
			}
		case *slack.UserTypingEvent:
			// Ignore
		case *slack.PresenceChangeEvent:
			// Ignore
		case *slack.FileSharedEvent:
			// Ignore
		default:
			// Ignore other events..
			log.Printf("Unhandled: %s\n", msg.Type)
		}
	}
}

func parseGroups(api *slack.Client) (map[string]bool, error) {
	const excludeArchived = true
	groups, err := api.GetGroups(excludeArchived)
	if err != nil {
		return nil, err
	}

	allowed := make(map[string]bool)
	for _, group := range groups {
		fmt.Printf("Allowed! ID=%q Name=%q\n", group.ID, group.Name)
		allowed[group.ID] = true
		allowed[group.Name] = true
	}

	allowed["GGYT2NJA1"] = true
	allowed["CGYT18RMK"] = true
	allowed["CJ87H4NNA"] = true

	return allowed, nil
}

func panicIf(err error) {
	if err != nil {
		log.Panic(err)
	}
}
