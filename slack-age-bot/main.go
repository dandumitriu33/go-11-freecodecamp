package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "<insert bot token>")
	os.Setenv("SLACK_APP_TOKEN", "<insert app token>")
	fmt.Println("BT: ", os.Getenv("SLACK_BOT_TOKEN"))
	fmt.Println("AT: ", os.Getenv("SLACK_APP_TOKEN"))

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 2002"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Println(err)
			}
			t := time.Now()
			currentYear := t.Year()
			age := currentYear - yob

			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

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
