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
	usertime := time.Now()

	attachment := slack.Attachment{
		Pretext: "Date",
		Text:    "Hello!",
		Color:   "#36a64f",
		Fields: []slack.AttachmentField{
			{
				Title: "The current time is:",
				Value: usertime.Local().String(),
			},
		},
	}

	_, timestamp, err := api.PostMessage(channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message sent at %s\n", timestamp)
}
