package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "<insert bot token>")
	os.Setenv("CHANNEL_ID", "<insert app token>")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{""}
	fmt.Println(api)
	fmt.Println(channelArr)
	fmt.Println(fileArr)

}
