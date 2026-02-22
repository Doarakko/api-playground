package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

type videoResult struct {
	VideoID         string
	VideoTitle      string
	ChannelName     string
	ChannelID       string
	SubscriberCount uint64
	ShopURL         string
}

var baseShopURLPattern = regexp.MustCompile(`https?://[a-zA-Z0-9\-]+\.base\.shop[/\S]*`)

func searchVideosByDescription(query string, maxResults int) {
	service := newYoutubeService(newClient())

	var allVideoIDs []string
	pageToken := ""

	// 1. search.list でページネーションしながら動画を収集
	for {
		searchCall := service.Search.List([]string{"snippet"}).
			Q(query).
			Type("video").
			MaxResults(50)
		if pageToken != "" {
			searchCall = searchCall.PageToken(pageToken)
		}
		searchResponse, err := searchCall.Do()
		if err != nil {
			log.Fatalf("Search API error: %v", err)
		}

		if pageToken == "" {
			fmt.Printf("検索結果（総数）: %d件\n\n", searchResponse.PageInfo.TotalResults)
		}

		for _, item := range searchResponse.Items {
			allVideoIDs = append(allVideoIDs, item.Id.VideoId)
		}

		if len(allVideoIDs) >= maxResults || searchResponse.NextPageToken == "" {
			break
		}
		pageToken = searchResponse.NextPageToken
	}

	if len(allVideoIDs) == 0 {
		fmt.Println("検索結果が見つかりませんでした")
		return
	}

	// 2. videos.list で説明欄を取得（50件ずつバッチ処理）
	var matched []videoResult
	channelIDs := make(map[string]bool)

	for i := 0; i < len(allVideoIDs); i += 50 {
		end := i + 50
		if end > len(allVideoIDs) {
			end = len(allVideoIDs)
		}
		batch := allVideoIDs[i:end]

		videosCall := service.Videos.List([]string{"snippet"}).
			Id(strings.Join(batch, ","))
		videosResponse, err := videosCall.Do()
		if err != nil {
			log.Fatalf("Videos API error: %v", err)
		}

		// 3. 説明欄から *.base.shop URL を含む動画をフィルタ
		for _, item := range videosResponse.Items {
			shopURL := baseShopURLPattern.FindString(item.Snippet.Description)
			if shopURL != "" {
				result := videoResult{
					VideoID:     item.Id,
					VideoTitle:  item.Snippet.Title,
					ChannelName: item.Snippet.ChannelTitle,
					ChannelID:   item.Snippet.ChannelId,
					ShopURL:     shopURL,
				}
				matched = append(matched, result)
				channelIDs[item.Snippet.ChannelId] = true
			}
		}
	}

	if len(matched) == 0 {
		fmt.Printf("説明欄に「%s」を含む動画が見つかりませんでした\n", query)
		return
	}

	// 4. channels.list でチャンネル登録者数を取得（50件ずつバッチ処理）
	subscriberMap := make(map[string]uint64)
	var channelIDList []string
	for id := range channelIDs {
		channelIDList = append(channelIDList, id)
	}

	for i := 0; i < len(channelIDList); i += 50 {
		end := i + 50
		if end > len(channelIDList) {
			end = len(channelIDList)
		}
		batch := channelIDList[i:end]

		channelsCall := service.Channels.List([]string{"statistics"}).
			Id(strings.Join(batch, ","))
		channelsResponse, err := channelsCall.Do()
		if err != nil {
			log.Fatalf("Channels API error: %v", err)
		}

		for _, ch := range channelsResponse.Items {
			subscriberMap[ch.Id] = ch.Statistics.SubscriberCount
		}
	}

	// マッチした動画に登録者数をセット
	for i := range matched {
		matched[i].SubscriberCount = subscriberMap[matched[i].ChannelID]
	}

	// 5. チャンネル登録者数の降順でソート
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].SubscriberCount > matched[j].SubscriberCount
	})

	// 6. 結果を出力
	fmt.Printf("=== 検索ワード: \"%s\" (ヒット数: %d) ===\n\n", query, len(matched))
	for i, v := range matched {
		fmt.Printf("%d. チャンネル名: %s (登録者数: %d)\n", i+1, v.ChannelName, v.SubscriberCount)
		fmt.Printf("   動画タイトル: %s\n", v.VideoTitle)
		fmt.Printf("   ショップURL: %s\n", v.ShopURL)
		fmt.Printf("   動画URL: https://www.youtube.com/watch?v=%s\n\n", v.VideoID)
	}
}
