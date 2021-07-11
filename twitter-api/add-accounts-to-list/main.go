package main

import (
	"log"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func newTwitterClient() *twitter.Client {
	config := oauth1.NewConfig(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client
}

func searchTweet(q string, tweetType string, queryHashtag bool) []twitter.Tweet {
	// exclude retweet
	q = q + " AND -filter:retweets"
	// Add hashtag
	if queryHashtag {
		q = "#" + q
	}

	client := newTwitterClient()
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:      q,
		Count:      100,
		ResultType: tweetType,
	})
	time.Sleep(2 * time.Second)
	if err != nil {
		log.Fatalf("[searchTweet error] %v %v", resp.Status, err)
	}

	log.Printf("Get %v tweets", len(search.Statuses))
	return search.Statuses
}

func listExists(listName string, userScreenName string) int64 {
	client := newTwitterClient()
	lists, resp, err := client.Lists.List(&twitter.ListsListParams{
		ScreenName: userScreenName,
	})
	time.Sleep(2 * time.Second)
	if err != nil {
		log.Fatal("[listExists error] ", resp.Status)
	}

	for i := 0; i < len(lists); i++ {
		if lists[i].Name == listName {
			log.Printf("%v is already created", listName)
			return lists[i].ID
		}
	}
	return -1
}

func createList(name string, mode string) int64 {
	client := newTwitterClient()
	list, resp, err := client.Lists.Create(name, &twitter.ListsCreateParams{
		Mode: mode,
	})
	time.Sleep(2 * time.Second)
	if err != nil {
		log.Fatal("[makeList error] ", resp.Status)
	}

	log.Printf("Create list")
	return list.ID
}

func addMember(listID int64, userID int64) {
	client := newTwitterClient()
	resp, err := client.Lists.MembersCreate(&twitter.ListsMembersCreateParams{
		ListID: listID,
		UserID: userID,
	})
	time.Sleep(2 * time.Second)
	if err != nil {
		log.Fatal("[addMember error] ", resp.Status)
	}
	log.Printf("[Add member] %v", userID)
}

func memberExists(listID int64, userID int64) bool {
	client := newTwitterClient()
	_, resp, err := client.Lists.MembersShow(&twitter.ListsMembersShowParams{
		ListID: listID,
		UserID: userID,
	})
	time.Sleep(2 * time.Second)
	if err != nil {
		if resp.StatusCode == 404 {
			log.Printf("%v is already member", userID)
			return true
		}
		log.Fatalf("[memberExists error] %v %v", resp.Status, err)
	}

	return false
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	query := os.Getenv("QUERY")
	listName := os.Getenv("LIST_NAME")
	listMode := os.Getenv("LIST_MODE")
	accountID := os.Getenv("ACCOUNT_ID")
	tweetType := os.Getenv("TWEET_TYPE")

	// Add hashtag
	var queryHashtag bool
	if os.Getenv("QUERY_HASHTAG") == "true" {
		queryHashtag = true
	} else {
		queryHashtag = false
	}

	tweetList := searchTweet(query, tweetType, queryHashtag)

	listID := listExists(listName, accountID)
	if listID == -1 {
		listID = createList(listName, listMode)
	}

	var userID int64
	for i := 0; i < len(tweetList); i++ {
		userID = tweetList[i].User.ID
		addMember(listID, userID)
	}
}
