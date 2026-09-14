package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// for this demo, were sorting teh article list in memory
// in a real applicarion, this list will most likely be fetch
// from a database or from static files
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 Body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 Body"},
}

func getAllArticels() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
