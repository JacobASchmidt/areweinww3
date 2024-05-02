package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	models "github.com/JacobASchmidt/areweinww3/src/models"
	"github.com/gofiber/fiber/v3"
	openai "github.com/sashabaranov/go-openai"
)

func StatusController(c fiber.Ctx) error {
	log.Print("Status request.")
	db := models.GetInstance()
	lastStatus := db.GetMostRecentStatus()
	log.Printf("Last status: %v", lastStatus)

	if time.Since(lastStatus.Date) < 5*time.Minute {
		log.Print("Responding with fresh status from DB.")
		return c.JSON(lastStatus)
	}
	log.Print("Status is stale. Requesting new status from OpenAI.")

	articles := db.GetArticles()
	var stringArticles string
	for _, article := range articles {
		stringArticles += article.ToString() + "\n"
	}

	openAiClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := openAiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `You are a fictional oracle that predicts if the fictional world of Earth is currently in a third world war based on fictional articles provided by the user message. You respond in only JSON with the following structure 
					{
						"status": "YES" | "NO",
						"subLine": "Global Conflict" | "Regional Wars",
						"explanation": "Quite a long drawn out explanation
					}`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: stringArticles,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return c.Status(500).SendString(err.Error())
	}

	lastMessage := resp.Choices[0].Message.Content
	var latestStatus models.Status
	err = json.Unmarshal([]byte(lastMessage), &latestStatus)
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
		return c.Status(500).SendString(err.Error())
	}
	latestStatus.Date = time.Now()
	err = db.InsertStatus(latestStatus)

	return c.JSON(latestStatus)
}
