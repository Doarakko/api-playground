package main

import (
	"fmt"
	"log"
	"net/http"

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

func printChannelInfo(channelID string) {
	service := newYoutubeService()
	call := service.Channels.List("snippet,contentDetails,statistics").
		Id(channelID).
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
	playlistID := item.ContentDetails.RelatedPlaylists.Uploads
	viewCount := item.Statistics.ViewCount
	subscriberCount := item.Statistics.SubscriberCount
	videoCount := item.Statistics.VideoCount

	fmt.Printf("channel id: %v\n\nチャンネル名: \n%v\n\n説明: %v\n\nサムネイルURL: %v\n\nplaylist id: %v\n\n総再生回数: %v\n\nチャンネル登録者数: %v\n\n動画数: %v\n",
		id,
		name,
		description,
		thumbnailURL,
		playlistID,
		viewCount,
		subscriberCount,
		videoCount,
	)
}

func main() {
	fmt.Println("【チャンネル情報】")
	channelID := "UC4YaOt1yT-ZeyB0OmxHgolA"
	printChannelInfo(channelID)
}
