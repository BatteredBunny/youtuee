package yt

import (
	"errors"
	"fmt"

	"google.golang.org/api/youtube/v3"
)

var ErrVideoNotFound = errors.New("youtube video not found")
var ErrFailedToRetriveInfo = errors.New("youtube api didn't return video info")

func YtApiGetVideoInfo(client *youtube.Service, videoId string) (v VideoInfo, err error) {
	response, err := client.Videos.List([]string{"snippet"}).Id(videoId).Do()
	if err != nil {
		return
	}

	if len(response.Items) == 0 {
		err = ErrVideoNotFound
		return
	}

	videoData := response.Items[0]

	v.Id = videoId
	v.OriginalUrl = fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)

	if videoData.Snippet != nil {
		v.Title = videoData.Snippet.Title
		v.Description = FormatDescription(videoData.Snippet.Description)
		v.Uploader = videoData.Snippet.ChannelTitle
		v.UploaderUrl = fmt.Sprintf("https://www.youtube.com/channel/%s", videoData.Snippet.ChannelId)

		if videoData.Snippet.Thumbnails != nil {
			var thumbnail *youtube.Thumbnail = nil

			// Tries to pick highest resolution
			if videoData.Snippet.Thumbnails.Maxres != nil {
				thumbnail = videoData.Snippet.Thumbnails.Maxres
			} else if videoData.Snippet.Thumbnails.High != nil {
				thumbnail = videoData.Snippet.Thumbnails.High
			} else if videoData.Snippet.Thumbnails.Medium != nil {
				thumbnail = videoData.Snippet.Thumbnails.Medium
			} else if videoData.Snippet.Thumbnails.Standard != nil {
				thumbnail = videoData.Snippet.Thumbnails.Standard
			} else if videoData.Snippet.Thumbnails.Default != nil {
				thumbnail = videoData.Snippet.Thumbnails.Default
			}

			if thumbnail != nil {
				v.Thumbnail = thumbnail.Url
			}
		}
	} else {
		err = ErrFailedToRetriveInfo
	}

	return
}
