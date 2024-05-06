package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/JacobASchmidt/areweinww3/src/controllers"
	"github.com/JacobASchmidt/areweinww3/src/models"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type Status struct {
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	SubLine     string    `json:"subLine"`
	Explanation string    `json:"explanation"`
	Date        time.Time `json:"date"`
}

type Article struct {
	Id          int    `json:"id"`
	Headline    string `json:"headline"`
	Description string `json:"Description"`
}

func status() Status {
	//GET FROM GPT-4

	statuses := []Status{
		{
			Status:      "YES",
			SubLine:     "Global Conflict",
			Explanation: "Quite a long drawn out explanation about how israel nuked iran or something",
		},
		{
			Status:      "NO",
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	db := models.GetInstance()
	defer db.Close()

	app.Use(func(c fiber.Ctx) error {
		log.Print(c.Route().Method, c.Route().Path)
		return c.Next()
	})

	app.Static("/", "../build")

	app.Get("api/v1/articles", controllers.ArticlesController)

	app.Get("api/v1/status", controllers.StatusController)

	log.Fatal(app.Listen(":3000"))
}
