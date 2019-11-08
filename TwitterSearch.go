package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Credentials struct does something. idk
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

//SearchResponse is response object for search tweet
type SearchResponse struct {
	Status   int
	Text     string
	FullText string
}

//SendTweetResponse is response for sending tweet
type SendTweetResponse struct {
	Status int
	ID     int64
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	return client, nil
}
func getCredentials() Credentials {
	return Credentials{
		AccessToken:       "125277653-iwDs38SZcAwYVQI76zn3O1MAHDesk0shEgSFvzKZ",
		AccessTokenSecret: "lgRCdt5ftq5xrZq3CvKjsYhkPglkkW9sXD2Le2tnZXNpF",
		ConsumerKey:       "xpedwliEFGhYQKNZFtG6Mm7BR",
		ConsumerSecret:    "texBxnSkmddS0a32iQLXmLttVJkXnrfmiB5YlmXg8BPcMUjChb",
	}
}

//SendTweet function sends tweet for authenticated user
func SendTweet() (SendTweetResponse, error) {
	creds := getCredentials()
	var clnt, err = getClient(&creds)
	if err != nil {
		log.Fatal("error retrieving client", err)
	}
	tweet, resp, err := clnt.Statuses.Update("this is another test tweet", nil)
	if err != nil {
		log.Println(err)
		return SendTweetResponse{Status: resp.StatusCode}, err
	}
	return SendTweetResponse{
		ID:     tweet.ID,
		Status: resp.StatusCode,
	}, nil
}

//SearchTweet searches a tweet or hashtag
func SearchTweet() SearchResponse {
	creds := getCredentials()
	var clnt, err = getClient(&creds)
	if err != nil {
		log.Fatal("error retrieving client", err)
	}
	var test = twitter.SearchTweetParams{
		Query:      "#GoLang",
		Count:      10,
		ResultType: "popular",
		Since:      "2019-01-01",
	}
	search, resp, err := clnt.Search.Tweets(&test)
	if err != nil {
		log.Println(err)
	}
	var searchResponse = SearchResponse{
		FullText: search.Statuses[0].FullText,
		Status:   resp.StatusCode,
		Text:     search.Statuses[0].Text,
	}
	return searchResponse

	// fmt.Println("resp: ", resp.StatusCode)
	// fmt.Println(search.Statuses[0].Text)
	// fmt.Println(search.Statuses[0].FullText)
	// return resp.Status
}
func main() {
	fmt.Println("test bot")
	status := SearchTweet()
	fmt.Println(status)
}
