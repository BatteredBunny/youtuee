package yt

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

var ErrFailedToFetchVideo = errors.New("failed to fetch video")

func YtDlpGetVideoInfo(ytdlpBinary string, videoId string) (v VideoInfo, err error) {
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
