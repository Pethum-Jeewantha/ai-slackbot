/*
 * Copyright (c) 2023 - present Pethum Jeewantha. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */

package service

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
	"os"
)

type WitAIService struct {
	client *witai.Client
}

func (witAIService *WitAIService) newWitAIService() *WitAIService {
	witAIService.client = witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	return witAIService
}

func (witAIService *WitAIService) getReceivedQuery(query string) (string, float64) {
	msg, _ := witAIService.client.Parse(&witai.MessageRequest{
		Query: query,
	})
	data, _ := json.MarshalIndent(msg, "", "    ")
	rough := string(data[:])

	value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
	confidence := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.confidence")

	return value.String(), confidence.Float()
}
