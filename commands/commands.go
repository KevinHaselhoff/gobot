package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/nlopes/slack"
)

var rtm *slack.RTM

// SetRTM sets singleton
func SetRTM(rtmPassed *slack.RTM) {
	rtm = rtmPassed
}

// CheckCommand is now commented
func CheckCommand(api *slack.Client, slackMessage slack.Msg, command string) {
	args := strings.Fields(command)
	callingUserProfile, _ := api.GetUserInfo(slackMessage.User)
	params := slack.PostMessageParameters{AsUser: true}

	if args[0] == "zug" {
		response := "<https://img.srv2.de/customer/sbahnMuenchen/newsticker/newsticker.html|Aktuelles>"
		response += " | <" + mvvRoute("Freising", "München, Hauptbahnhof") + "|Ins Büro>"
		response += " | <" + mvvRoute("München, Hauptbahnhof", "Freising") + "|Nach Hause>"

		api.PostMessage(slackMessage.Channel, response, params)
	} else if args[0] == "wb" {
		response := ":partly_sunny_rain: <https://darksky.net/forecast/48.398,11.9227/ca24/de#week|Wetter Berglern>"
		api.PostMessage(slackMessage.Channel, response, params)
	} else if args[0] == "gce" {
		response := "GCE will be implemented shortly"
		api.PostMessage(slackMessage.Channel, response, params)
	} else if args[0] == "help" {
		response := ":sun_behind_rain_cloud: `wb`: Wetter Berglern\n" +
			":metro: `zug`: Aktuelles | Ins Büro | Nach Hause\n"
		api.PostMessage(slackMessage.Channel, response, params)
	} else {
		rtm.SendMessage(rtm.NewOutgoingMessage("This command is unknown <@"+callingUserProfile.Name+">? Try `help` instead...",
			slackMessage.Channel))
	}
}

func mvvRoute(origin string, destination string) string {
	loc, _ := time.LoadLocation("Europe/Berlin")
	date := time.Now().In(loc)

	yearObj := date.Year()
	monthObj := int(date.Month())
	dayObj := date.Day()
	hourObj := date.Hour()
	minuteObj := date.Minute()

	month := strconv.Itoa(monthObj)
	hour := strconv.Itoa(hourObj)
	day := strconv.Itoa(dayObj)
	minute := strconv.Itoa(minuteObj)
	year := strconv.Itoa(yearObj)

	return "http://efa.mvv-muenchen.de/mvv/XSLT_TRIP_REQUEST2?&language=de" +
		"&anyObjFilter_origin=0&sessionID=0&itdTripDateTimeDepArr=dep&type_destination=any" +
		"&itdDateMonth=" + month + "&itdTimeHour=" + hour + "&anySigWhenPerfectNoOtherMatches=1" +
		"&locationServerActive=1&name_origin=" + origin + "&itdDateDay=" + day + "&type_origin=any" +
		"&name_destination=" + destination + "&itdTimeMinute=" + minute + "&Session=0&stateless=1" +
		"&SpEncId=0&itdDateYear=" + year
}
