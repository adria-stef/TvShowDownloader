//Package configuration provides ptimitives for xml and yml unmarshaling
package configuration

import (
	"encoding/xml"
	"log"

	"github.com/adria-stef/TvShowDownloader/model"
)

//GetRssFeed unmarrshales rss.xml and returns rss feed struct
func GetRssFeed(rssFile []byte) model.RssFeed {
	rssFeed := model.RssFeed{}
	err := xml.Unmarshal(rssFile, &rssFeed)
	if err != nil {
		log.Printf("Error while parsing RssFeed: %v", err)
		return model.RssFeed{}
	}

	return rssFeed
}
