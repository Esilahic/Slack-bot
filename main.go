package main

import (
	"os"

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

}
