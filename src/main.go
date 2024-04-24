package main

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v3"
)

type Status struct {
	Status      string `json:"status"`
	TextColor   string `json:"textColor"`
	SubLine     string `json:"subLine"`
	Explanation string `json:"explanation"`
}

type Article struct {
	Headline    string `json:"headline"`
	Description string `json:"Description"`
}

func status() Status {
	//GET FROM GPT-4

	statuses := []Status{
		{
			Status:      "YES",
			TextColor:   "#e63946",
			SubLine:     "Global Conflict",
			Explanation: "Quite a long drawn out explanation about how israel nuked iran or something",
		},
		{
			Status:      "NO",
			TextColor:   "#4caf50",
			SubLine:     "Regional Wars",
			Explanation: "Quite a long drawn out explanation about how israel hasn't nuked iran or something",
		},
	}

	return statuses[rand.Intn(2)]
}

func articles() []Article {
	//GET FROM ???
	return []Article{
		{Headline: "Article Title 1", Description: "Detailed analysis of recent global events."},
		{Headline: "Article Title 2", Description: "Expert opinions on geopolitical tensions."},
		{Headline: "Article Title 3", Description: "Historical context behind today's conflicts."},
		{Headline: "Article Title 4", Description: "Predictive insights into future global trends."},
	}
}

func main() {
	app := fiber.New()

	app.Use(func(c fiber.Ctx) error {
		log.Print(c.Route().Method, c.Route().Path)
		return c.Next()
	})

	app.Static("/", "../build")

	app.Get("api/v1/articles", func(c fiber.Ctx) error {
		log.Print("in articles")
		return c.JSON(articles())
	})

	app.Get("api/v1/status", func(c fiber.Ctx) error {
		log.Print("in status")
		return c.JSON(status())
	})

	log.Fatal(app.Listen(":3000"))
}
