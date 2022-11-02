package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4322258483633-4295116525767-J3C7MGCO8kl8tr1C56gm44og")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0493JLHUBU-4309634223698-95767e4fc6eab983449aaf7c18f4e1287007d98daeb3c332b72898a7c4902d7c")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Printf("Timestamp: %v - Command: %v - Parameters: %v - Event: %v\n",
			event.Timestamp, event.Command, event.Parameters, event.Event)
	}
}
