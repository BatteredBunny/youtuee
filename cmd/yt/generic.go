package yt

type VideoInfo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	OriginalUrl string `json:"original_url"`

	Uploader    string `json:"uploader"`
	UploaderUrl string `json:"uploader_url"`
}
