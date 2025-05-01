package yt

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Much faster than yt-dlp but only works on residental IPs
func ScraperGetVideoInfo(videoId string) (v VideoInfo, err error) {
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	v.Id = videoId

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		attr, exists := s.Attr("property")
		if exists {
			if attr == "og:title" && v.Title == "" {
				v.Title = s.AttrOr("content", "")
			} else if attr == "og:description" && v.Description == "" {
				v.Description = s.AttrOr("content", "")
			} else if attr == "og:image" && v.Thumbnail == "" {
				v.Thumbnail = s.AttrOr("content", "")
			} else if attr == "og:url" && v.OriginalUrl == "" {
				v.OriginalUrl = s.AttrOr("content", "")
			}
		}
	})

	doc.Find("#watch7-content span").Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("itemprop", "") == "author" {
			s.Find("link").Each(func(i int, s *goquery.Selection) {
				attr, exists := s.Attr("itemprop")
				if exists {
					if attr == "url" && v.UploaderUrl == "" {
						v.UploaderUrl = s.AttrOr("href", "")
					} else if attr == "name" && v.Uploader == "" {
						v.Uploader = s.AttrOr("content", "")
					}
				}
			})
		}
	})

	return
}
