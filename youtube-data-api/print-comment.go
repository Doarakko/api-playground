package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func newYoutubeService() *youtube.Service {
	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("YOUTUBE_API_KEY")},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	return service
}

func printComments(VideoID string) {
	service := newYoutubeService()
	call := service.CommentThreads.List("id,snippet").
		VideoId(VideoID).
		Order("relevance").
		SearchTerms("草").
		MaxResults(10)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}

	for _, item := range response.Items {
		authorName := item.Snippet.TopLevelComment.Snippet.AuthorDisplayName
		text := item.Snippet.TopLevelComment.Snippet.TextDisplay
		likeCnt := item.Snippet.TopLevelComment.Snippet.LikeCount
		replyCnt := item.Snippet.TotalReplyCount
		fmt.Printf("\"%v\" by %v\nいいね数: %v 返信数: %v\n\n", text, authorName, likeCnt, replyCnt)
	}
}

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	videoID := "_fJNAf6_4NU"
	printComments(videoID)
}
