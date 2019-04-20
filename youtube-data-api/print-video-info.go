package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

func printVideoInfo(videoID string) {
	service := newYoutubeService()
	call := service.Videos.List("id,snippet,Statistics").
		Id(videoID).
		MaxResults(1)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}

	item := response.Items[0]
	id := item.Id
	name := item.Snippet.Title
	description := item.Snippet.Description
	thumbnailURL := item.Snippet.Thumbnails.High.Url
	viewCount := item.Statistics.ViewCount
	commentCount := item.Statistics.CommentCount
	likeCount := item.Statistics.LikeCount
	dislikeCount := item.Statistics.DislikeCount
	channelID := item.Snippet.ChannelId
	categoryID := item.Snippet.CategoryId
	categoryName := getVideoCategory(categoryID)
	uploadDate, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("video id: %v\n\nタイトル: %v\n\n説明: \n%v\nサムネイルURL: %v\n\n再生回数: %v\n\nコメント数: %v\n\n高評価数: %v\n\n低評価数: %v\n\nchannel id: %v\n\nアップロード日時: %v\nカテゴリ ID: %v\n",
		id,
		name,
		description,
		thumbnailURL,
		viewCount,
		commentCount,
		likeCount,
		dislikeCount,
		channelID,
		uploadDate,
		categoryName,
	)
}

func getVideoCategory(categoryID string) string {
	service := newYoutubeService()
	call := service.VideoCategories.List("id,snippet").
		Id(categoryID)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}

	item := response.Items[0]
	categoryName := item.Snippet.Title

	return categoryName
}

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("【動画情報】")

	videoID := "xCSxGEYEcNk"
	printVideoInfo(videoID)

}
