package main

import (
	"fmt"
	"time"
)

type Article struct {
	Title string
	Body  string
	Date  time.Time
}

func main() {
	article := &Article{
		Title: "My Article",
	}

	fmt.Println(article.Title)
	article.Date = time.Now()
	fmt.Println(article.Date)
}
