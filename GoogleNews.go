package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type searchResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []article `json:"articles"`
}

type article struct {
	Sources     source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//GetGoogleNews returns the search result on basis of keyword
func GetGoogleNews() {
	response, err := http.Get("https://newsapi.org/v2/everything?q=GoLang&sortBy=popularity&apiKey=1c858444e19547d9b8c5904c50f64816")
	var searchRes searchResponse
	if err != nil {
		log.Fatal("error retrieving data", err)
	} else {
		//read response
		data, _ := ioutil.ReadAll(response.Body)
		//parse response to object
		json.Unmarshal(data, &searchRes)
		//test print. remove it later
		fmt.Println("result:", searchRes.TotalResults)
		//save to a file
		file, _ := json.MarshalIndent(searchRes, "", " ")
		ioutil.WriteFile("outputfile.txt", file, 0644)
	}

}
