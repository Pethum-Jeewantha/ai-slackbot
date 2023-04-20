/*
 * Copyright (c) 2023 - present Pethum Jeewantha. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */

package main

import (
	"ai-slackbot/service"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	dotenvErr := godotenv.Load(".env")
	if dotenvErr != nil {
		fmt.Println(".env file not found")
	}

	slackService := (&service.SlackService{}).NewSlackService()
	slackService.CreateCommand()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	botErr := slackService.GetBot().Listen(ctx)
	if botErr != nil {
		log.Fatalln(botErr)
	}
}
