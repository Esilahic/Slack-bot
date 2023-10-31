package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func HandleEventMessage(event slackevents.EventsAPIEvent, client *slack.Client) error {
    switch event.Type {
 
    case slackevents.CallbackEvent:
 
        innerEvent := event.InnerEvent
 
        switch evnt := innerEvent.Data.(type) {
            err := HandleAppMentionEventToBot(evnt, client)
            if err != nil {
                return err
            }
        }
    default:
        return errors.New("unsupported event type")
    }
    return nil
}

func main() {
	// TODO: add eventlistener for user commands in slack client
	godotenv.Load("../.env")
	token := os.Getenv("SLACK_API_TOKEN")
	websocket := os.Getenv("SLACK_APP_TOKEN")

	client := slack.New(token, slack.OptionAppLevelToken(websocket))

	socketClient := socketmode.New(
		client,
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")
				return
			case event := <-socketClient.Events:

				switch event.Type {

				case socketmode.EventTypeEventsAPI:

					eventsAPI, ok := event.Data.(slackevents.EventsAPIEvent)
					if !ok {
						log.Printf("Could not type cast the event to the EventsAPI: %v\n", event)
						continue
					}

					socketClient.Ack(*event.Request)
					err := HandleEventMessage(eventsAPI, client)
                    if err != nil {
                        log.Fatal(err)
                    }
				}
			}
		}
	}(ctx, client, socketClient)

	socketClient.Run()
}
