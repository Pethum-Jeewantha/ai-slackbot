/*
 * Copyright (c) 2023 - present Pethum Jeewantha. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Krognol/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
	"log"
	"os"
)

var bot *slacker.Slacker
var client *witai.Client
var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("==== Command Events ====")
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
		fmt.Println(".env file not found")
	}

	bot = slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	client = witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}

	go printCommandEvents(bot.CommandEvents())

	CreateCommand()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	botErr := bot.Listen(ctx)
	if botErr != nil {
		log.Fatalln(botErr)
	}
}

func CreateCommand() {
	bot.Command("Hey <message>", &slacker.CommandDefinition{
		Description: "Send any question to wolfram",
		Examples:    []string{"Who is the president of sri lanka"},
		Handler: func(botContext slacker.BotContext, request slacker.Request, writer slacker.ResponseWriter) {
			query := request.Param("message")
			query = getReceivedQuery(query)

			res := getAnswer(query)

			writErr := writer.Reply(res)
			if writErr != nil {
				log.Fatalln(writErr)
			}
		},
	})
}

func getReceivedQuery(query string) string {
	msg, _ := client.Parse(&witai.MessageRequest{
		Query: query,
	})
	data, _ := json.MarshalIndent(msg, "", "    ")
	rough := string(data[:])
	value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")

	return value.String()
}

func getAnswer(query string) string {
	res, wolframErr := wolframClient.GetSpokentAnswerQuery(query, wolfram.Metric, 1000)
	if wolframErr != nil {
		log.Fatalln(wolframErr)
	}

	return res
}
