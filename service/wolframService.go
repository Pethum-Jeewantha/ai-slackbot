/*
 * Copyright (c) 2023 - present Pethum Jeewantha. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */

package service

import (
	"github.com/Krognol/go-wolfram"
	"log"
	"os"
)

type WolframService struct {
	client *wolfram.Client
}

func (wolframService *WolframService) newWolframService() *WolframService {
	wolframService.client = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}
	return wolframService
}

func (wolframService *WolframService) getAnswer(query string) string {
	res, wolframErr := wolframService.client.GetSpokentAnswerQuery(query, wolfram.Metric, 1000)
	if wolframErr != nil {
		log.Fatalln(wolframErr)
	}

	return res
}
