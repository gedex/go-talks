package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Article struct {
	Id    int
	Title string
}

// START API
type articlesApi struct {
	articles []*Article
}

func (aa *articlesApi) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := json.NewEncoder(w).Encode(aa.articles); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (aa *articlesApi) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	article := aa.getById(id)
	if article == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

// END API

func (aa *articlesApi) getById(id int) (article *Article) {
	for _, a := range aa.articles {
		if a.Id == id {
			article = a
		}
	}
	return
}

// START MAIN
func main() {
	aa := new(articlesApi)
	aa.articles = append(aa.articles, &Article{1, "First Article"})
	aa.articles = append(aa.articles, &Article{2, "Second Article"})

	r := httprouter.New()
	r.GET("/articles", aa.Index)
	r.GET("/articles/:id", aa.Get)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// END MAIN
