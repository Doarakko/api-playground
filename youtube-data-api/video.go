package main

import (
	"fmt"
	"log"
	"time"
)

func printVideoInfo(videoID string) {
	service := newYoutubeService(newClient())
	call := service.Videos.List([]string{"id", "snippet", "statistics"}).
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
	service := newYoutubeService(newClient())
	call := service.VideoCategories.List([]string{"id", "snippet"}).
		Id(categoryID)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("%v", err)
	}

	item := response.Items[0]
	categoryName := item.Snippet.Title

	return categoryName
}
