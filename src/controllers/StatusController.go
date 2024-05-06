package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	models "github.com/JacobASchmidt/areweinww3/src/models"
	"github.com/gofiber/fiber/v3"
	openai "github.com/sashabaranov/go-openai"
)

type StatusResponse struct {
	Status   models.Status            `json:"status"`
	Articles []models.NewsApiHeadline `json:"articles"`
}

func StatusController(c fiber.Ctx) error {
	log.Print("Status request.")
	db := models.GetInstance()
	lastStatus := db.GetMostRecentStatus()
	log.Printf("Last status: %v", lastStatus)

	if time.Since(lastStatus.Date) < 5*time.Minute {
		log.Print("Responding with fresh status from DB.")
		articles := db.GetArticles()
		response := StatusResponse{
			Status:   lastStatus,
			Articles: articles,
		}
		return c.JSON(response)
	}
	log.Print("Status is stale. Fetching articles and and getting new status from OpenAI")

	//Get Articles
	newsApiKey := os.Getenv("NEWSAPI_API_KEY")
	if newsApiKey == "" {
		log.Fatal("NEWSAPI_API_KEY is not set.")
	}

	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&category=general&apiKey=%s", newsApiKey)
	newsApiResp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return c.Status(500).SendString(err.Error())
	}

	var newsApiResponse NewsApiResponse
	err = json.NewDecoder(newsApiResp.Body).Decode(&newsApiResponse)
	if err != nil {
		log.Print(err)
		return c.Status(500).SendString(err.Error())
	}
	newsApiResp.Body.Close()

	db.DeleteArticles()
	var stringArticles string
	for _, article := range newsApiResponse.Articles {
		db.InsertArticle(article)
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
	if err != nil {
		fmt.Printf("InsertStatus error: %v\n", err)
		return c.Status(500).SendString(err.Error())
	}

	statusResponse := StatusResponse{
		Status:   latestStatus,
		Articles: newsApiResponse.Articles,
	}
	return c.JSON(statusResponse)
}
