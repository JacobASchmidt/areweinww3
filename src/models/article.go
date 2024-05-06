package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type NewsApiHeadline struct {
	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

func (h NewsApiHeadline) ToString() string {
	return fmt.Sprintf("Title: %s, Description: %s", h.Title, h.Description)
}

func setUpArticlesTable(db *sql.DB) {
	articlesQuery := `
		CREATE TABLE IF NOT EXISTS articles (
			id INTEGER PRIMARY KEY, 
			source_id TEXT,
			source_name TEXT,
			author TEXT,
			title TEXT,
			description TEXT,
			url TEXT,
			urlToImage TEXT,
			publishedAt TIMESTAMP,
			content TEXT
		)`

	_, err := db.Exec(articlesQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Database) GetArticles() []NewsApiHeadline {
	query := "SELECT * FROM articles LIMIT 10"
	rows, err := d.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var articles []NewsApiHeadline
	for rows.Next() {
		var article NewsApiHeadline
		var id int
		err = rows.Scan(
			&id,
			&article.Source.ID,
			&article.Source.Name,
			&article.Author,
			&article.Title,
			&article.Description,
			&article.URL,
			&article.URLToImage,
			&article.PublishedAt,
			&article.Content,
		)
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
	}
	return articles
}

func (d *Database) InsertArticle(article NewsApiHeadline) error {
	query := `INSERT INTO articles (source_id, source_name, author, title, description, url, urlToImage, publishedAt, content) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		article.Source.ID,
		article.Source.Name,
		article.Author,
		article.Title,
		article.Description,
		article.URL,
		article.URLToImage,
		article.PublishedAt,
		article.Content,
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) DeleteArticles() error {
	query := `DELETE FROM articles`
	_, err := d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
