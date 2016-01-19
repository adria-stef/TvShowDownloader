package parser

import (
	"encoding/xml"
)

type RssFeed struct {
	Rss     string  `xml:"rss"`
	Channel Channel `xml:"channel"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Ttl         int      `xml:"ttl"`
	Items       []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name  `xml:"item"`
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Category    string    `xml:"category"`
	Link        string    `xml:"link"`
	Guid        string    `xml:"guid"`
	PubDate     string    `xml:"pubdate"`
	Enclosure   Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	XMLName    xml.Name `xml:"enclosure"`
	Content    string   `xml:"enclosure"`
	TorrentUrl string   `xml:"url,attr"`
}
