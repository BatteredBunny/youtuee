package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

type VideoInfo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	OriginalUrl string `json:"original_url"`

	Uploader    string `json:"uploader"`
	UploaderUrl string `json:"uploader_url"`
}

var ErrFailedToFetchVideo = errors.New("failed to fetch video")

func GetVideoInfo(ytdlpBinary string, videoId string) (v VideoInfo, err error) {
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)
	bs, err := exec.Command(ytdlpBinary, "-j", url).Output()
	if err != nil {
		err = ErrFailedToFetchVideo
		return
	}

	err = json.Unmarshal(bs, &v)
	if err != nil {
		return
	}

	v.Description = FormatDescription(v.Description)

	return
}
