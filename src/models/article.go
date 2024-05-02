package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Article struct {
	Id          int    `json:"id"`
	Headline    string `json:"headline"`
	Description string `json:"Description"`
}

func (a Article) ToString() string {
	return fmt.Sprintf("Headline: %s, Description: %s", a.Headline, a.Description)
}

func setUpArticlesTable(db *sql.DB) {
	articlesQuery := `
		CREATE TABLE IF NOT EXISTS articles (
			id INTEGER PRIMARY KEY, 
			headline TEXT, 
			description TEXT
		)`

	_, err := db.Exec(articlesQuery)
	if err != nil {
		log.Fatal(err)
	}

	//Insert some test articles
	articles := []Article{
		{Headline: "Article Title 1", Description: "Detailed analysis of recent global events."},
		{Headline: "Article Title 2", Description: "Expert opinions on geopolitical tensions."},
		{Headline: "Article Title 3", Description: "Historical context behind today's conflicts."},
		{Headline: "Article Title 4", Description: "Predictive insights into future global trends."},
	}
	for _, article := range articles {
		_, err := db.Exec("INSERT INTO articles (headline, description) VALUES (?, ?)", article.Headline, article.Description)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (d *Database) GetArticles() []Article {
	query := "SELECT * FROM articles LIMIT 10"
	rows, err := d.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Headline, &article.Description)
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
	}
	return articles
}
