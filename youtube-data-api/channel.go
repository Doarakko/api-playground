package main

import (
	"fmt"
	"log"
)

func printChannelInfo(channelID string) {
	service := newYoutubeService(newClient())
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
