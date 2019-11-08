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
func sendTweet() string {
	creds := getCredentials()
	var clnt, err = getClient(&creds)
	if err != nil {
		log.Fatal("error retrieving client", err)
	}
	tweet, resp, err := clnt.Statuses.Update("this is another test tweet", nil)
	if err != nil {
		log.Println(err)
		//return string(err.Error)
	}
	if resp.StatusCode == 200 {
		return "success"
	}

	fmt.Println("%+v\n", resp)
	fmt.Println("%+v\n", tweet)

	return "error sending tweet"
}

func SearchTweet() string {
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
	fmt.Println("resp: ", resp.StatusCode)
	fmt.Println(search.Statuses[1].Text)
	fmt.Println(search.Statuses[1].FullText)
	return resp.Status
}
func main() {
	fmt.Println("test bot")
	status := SearchTweet()
	fmt.Println(status)
}
