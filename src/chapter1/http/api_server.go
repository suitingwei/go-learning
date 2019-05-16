package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	_ = json.NewEncoder(w).Encode(articles)
	fmt.Printf("[%s] index page visited\n", time.Now().String())
}

func testPage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the test page")
	fmt.Printf("[%s] test page visited\n", time.Now().String())
}
func Run() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/test", testPage)
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
