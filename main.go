package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	//godotenv load
	godotenv.Load(".env")
	token := os.Getenv("SLACK_API_TOKEN")
	channel := os.Getenv("CHANNEL_ID_TOKEN")
	//slack api token and channel id
	api := slack.New(token)
	//slack-go -read docs and create bot

	attachment := slack.Attachment{
		Pretext: "Date",
		Text:    "Current time is:",
		Color:   "#36a64f",
		Fields: []slack.AttachmentField{
			{
				Title: "Response",
				Value: time.Now().String(),
			},
		},
	}

	_, timestamp, err := api.PostMessage(channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message sent at %s", timestamp)
}
