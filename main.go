package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	witai "github.com/wit-ai/wit-go/v2"
	"log"
	"os"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	dotenvErr := godotenv.Load(".env")
	if dotenvErr != nil {
		log.Fatalln(dotenvErr)
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{
		Description: "Send any question to wolfram",
		Examples:    []string{"Who is the president of sri lanka"},
		Handler: func(botContext slacker.BotContext, request slacker.Request, writer slacker.ResponseWriter) {
			query := request.Param("message")

			msg, _ := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])
			fmt.Println(rough)

			writErr := writer.Reply("Received")
			if writErr != nil {
				log.Fatalln(writErr)
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	botErr := bot.Listen(ctx)
	if botErr != nil {
		log.Fatalln(botErr)
	}
}
