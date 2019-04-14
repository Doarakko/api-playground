package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

const key = "AIzaSyAS1rtt0HrtBBAcww6HsrNnHd5nTT_hH_k"

func newYoutubeService() *youtube.Service {
	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
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
	uploadDate, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("video id: %v\n\nタイトル: %v\n\n説明: \n%v\nサムネイルURL: %v\n\n再生回数: %v\n\nコメント数: %v\n\n高評価数: %v\n\n低評価数: %v\n\nchannel id: %v\n\nアップロード日時: %v\n",
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
	)
}

func main() {
	fmt.Println("【動画情報】")
	videoID := "wT_GFTDpUno"
	printVideoInfo(videoID)
}
