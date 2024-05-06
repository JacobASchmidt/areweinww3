package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JacobASchmidt/areweinww3/src/models"
	"github.com/gofiber/fiber/v3"
)

type NewsApiResponse struct {
	Status       string                   `json:"status"`
	TotalResults int                      `json:"totalResults"`
	Articles     []models.NewsApiHeadline `json:"articles"`
}

func ArticlesController(c fiber.Ctx) error {
	log.Print("Articles request.")
	// db := models.GetInstance()

	newsApiKey := os.Getenv("NEWSAPI_API_KEY")
	if newsApiKey == "" {
		log.Fatal("NEWSAPI_API_KEY is not set.")
	}

	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%s", newsApiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return err
	}
	defer resp.Body.Close()

	var newsApiResponse NewsApiResponse
	err = json.NewDecoder(resp.Body).Decode(&newsApiResponse)
	if err != nil {
		log.Print(err)
		return err
	}

	fmt.Printf("Status: %s\n", newsApiResponse.Status)
	fmt.Printf("Total Results: %d\n", newsApiResponse.TotalResults)
	for _, article := range newsApiResponse.Articles {
		fmt.Printf("Source ID: %s\n", article.Source.ID)
		fmt.Printf("Source Name: %s\n", article.Source.Name)
		fmt.Printf("Author: %s\n", article.Author)
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("Description: %s\n", article.Description)
		fmt.Printf("URL: %s\n", article.URL)
		fmt.Printf("URLToImage: %s\n", article.URLToImage)
		fmt.Printf("PublishedAt: %s\n", article.PublishedAt.String())
		fmt.Printf("Content: %s\n", article.Content)
		fmt.Println("-----------------------------")
	}

	return c.JSON(newsApiResponse.Articles)

}
