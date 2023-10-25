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
	godotenv.Load(".env")
	token := os.Getenv("SLACK_API_TOKEN")
	channel := os.Getenv("CHANNEL_ID_TOKEN")

	api := slack.New(token)
	usertime := time.Now()

	attachment := slack.Attachment{
		Pretext: "SCTB!",
		Text:    "Hello",
		Color:   "#a85",
		Fields: []slack.AttachmentField{
			{
				Title: "The current Date and Time are:",
				Value: usertime.Local().String(),
			},
		},
	}

	_, timestamp, err := api.PostMessage(channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message sent at %s\n", timestamp)

	//TODO: add eventlistener for user commands in slack client
}
