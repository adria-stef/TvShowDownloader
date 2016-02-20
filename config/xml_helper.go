package configuration

import (
	"encoding/xml"
	"log"

	"github.com/adria-stef/TvShowDownloader/model"
)

//GetRssFeed returns rss feed
func GetRssFeed(rssFile []byte) model.RssFeed {
	rssFeed := model.RssFeed{}
	err := xml.Unmarshal(rssFile, &rssFeed)
	if err != nil {
		log.Printf("Error while parsing RssFeed: %v", err)
		return model.RssFeed{}
	}

	return rssFeed
}
