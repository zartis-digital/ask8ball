package main

import (
	"../../personas"
	"../../triggerbot"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"nyanbot"
	"os"
)

var (
	ask8ballConfiguration = triggerbot.Configuration{
		Username:  "Ask8Ball",
		Keywords:  personas.Ask8BallKeywords,
		IconURLs:  personas.Ask8BallIconURLs,
		Phrases:   personas.Ask8BallPhrases,
		Reactions: personas.Ask8BallReactions,
	}
)

func main() {
	token := os.Getenv("NYANBOT_TOKEN")
	if len(token) == 0 {
		log.Fatalln("NYANBOT_TOKEN env var is required!")
	} else {
		log.Print(token)
	}

	plugins := []nyanbot.Plugin{
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
			// Ignore
		case *slack.ConnectingEvent:
			// Can't use api.SetUserCustomStatus for bots :(
		case *slack.ConnectedEvent:
			log.Printf("Connected! ConnectionCount=%d", ev.ConnectionCount)
		case *slack.MessageEvent:
			// log.Printf("Message: %+v\n", ev)
			if ev.Msg.Text == "stop-nyanbot" {
				os.Exit(1)
			}

			isInAllowedChannels := allowedChannels[ev.Channel]
			if !isInAllowedChannels {
				log.Printf("Channel not allowed %s %+v", ev.Channel, ev);
				continue
			}

			for _, plugin := range plugins {
				if ev.Msg.Username == plugin.Name() {
					continue
				}
				err := plugin.HandleMessage(api, rtm, ev)
				if err == nyanbot.ErrStop {
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
		if group.Name == "ask8ball" || group.Name == "ask8ball-testing" {
			fmt.Printf("Allowed! ID=%q Name=%q\n", group.ID, group.Name)
			allowed[group.ID] = true
		} else {
			fmt.Printf("NotAllowed! ID=%q Name=%q\n", group.ID, group.Name)
		}
	}

	allowed["GGYT2NJA1"] = true;
	allowed["CGYT18RMK"] = true;

	return allowed, nil
}

func panicIf(err error) {
	if err != nil {
		log.Panic(err)
	}
}
