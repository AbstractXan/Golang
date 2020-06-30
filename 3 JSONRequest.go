/* 	Basic REST API
	Endpoints:
	8080/articles : Returns all articles
	https://tutorialedge.net/golang/creating-restful-api-with-golang/
*/
package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Article struct{
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article
var articles Articles

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(articles)
}

func homepage(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}


func handleRequests(){
	http.HandleFunc("/",homepage)
	http.HandleFunc("/articles",returnAllArticles)
   	http.ListenAndServe(":8080", nil)
}

func main() {
    articles = Articles{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}
