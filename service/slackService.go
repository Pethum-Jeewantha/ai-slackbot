/*
 * Copyright (c) 2023 - present Pethum Jeewantha. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */

package service

import (
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
)

type SlackService struct {
	bot            *slacker.Slacker
	witAIService   *WitAIService
	wolframService *WolframService
}

func (slackService *SlackService) NewSlackService() *SlackService {
	slackService.bot = slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	slackService.witAIService = (&WitAIService{}).newWitAIService()
	slackService.wolframService = (&WolframService{}).newWolframService()

	go slackService.printCommandEvents(slackService.bot.CommandEvents())

	return slackService
}

func (slackService *SlackService) GetBot() *slacker.Slacker {
	return slackService.bot
}

func (slackService *SlackService) CreateCommand() {
	slackService.bot.Command("Hey <message>", &slacker.CommandDefinition{
		Description: "Send any question to wolfram",
		Examples:    []string{"Who is the president of sri lanka"},
		Handler: func(botContext slacker.BotContext, request slacker.Request, writer slacker.ResponseWriter) {
			query := request.Param("message")
			query = slackService.witAIService.getReceivedQuery(query)

			res := slackService.wolframService.getAnswer(query)

			writErr := writer.Reply(res)
			if writErr != nil {
				log.Fatalln(writErr)
			}
		},
	})
}

func (slackService *SlackService) printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("==== Command Events ====")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
