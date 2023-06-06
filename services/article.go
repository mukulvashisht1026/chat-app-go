package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description`
	Content     string `json:"content"`
}

type Articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit all articles ... ")
	articles := Articles{
		Article{
			Title:       "main title",
			Description: "some description",
			Content:     "some content",
		},
	}
	json.NewEncoder(w).Encode(articles)
}
