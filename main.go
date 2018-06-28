package main

import (
	"gobot/commands"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/nlopes/slack"
)

var botID = " "

func main() {
	api := slack.New(os.Getenv("slackToken"))
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(false)

	rtm := api.NewRTM()
	commands.SetRTM(rtm)
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID
				response := "Moin Moin!"
				api.PostMessage(slackMessage.Channel, response, params)

			case *slack.MessageEvent:
				callerID := ev.Msg.User

				// only respond to messages sent to me by others on the same channel:
				if ev.Msg.Type == "message" && callerID != botID && ev.Msg.SubType != "message_deleted" &&
					(strings.Contains(ev.Msg.Text, "<@"+botID+">") ||
						strings.HasPrefix(ev.Msg.Channel, "K")) {
					originalMessage := ev.Msg.Text
					// strip out bot's name and spaces
					parsedMessage := strings.TrimSpace(strings.Replace(originalMessage, "<@"+botID+">", "", -1))
					r, n := utf8.DecodeRuneInString(parsedMessage)
					parsedMessage = string(unicode.ToLower(r)) + parsedMessage[n:]

					userInfo, _ := rtm.GetUserInfo(ev.Msg.User)
					userName := userInfo.Name
					logger.Printf("%s: %s\n", userName, parsedMessage)

					commands.CheckCommand(api, ev.Msg, parsedMessage)
				}

			case *slack.RTMError:
				logger.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				logger.Println("Invalid credentials")
				break

			default:
				// Ignore other events..
			}
		}
	}
}
