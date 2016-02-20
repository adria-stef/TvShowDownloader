package model

import "encoding/xml"

//RssFeed contains Rss and Channel
type RssFeed struct {
	Rss     string  `xml:"rss"`
	Channel Channel `xml:"channel"`
}

//Rss contains version
type Rss struct {
	XMLName xml.Name `xml:"rss"`
}

//Channel contains title, link, description, items
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Items       []Item   `xml:"item"`
}

//Item contains title, category, author, link, guid, pubdate, enclosure
type Item struct {
	XMLName   xml.Name  `xml:"item"`
	Title     string    `xml:"title"`
	Category  string    `xml:"category"`
	Author    string    `xml:"author"`
	Link      string    `xml:"link"`
	GUID      string    `xml:"guid"`
	PubDate   string    `xml:"pubdate"`
	Enclosure Enclosure `xml:"enclosure"`
}

//Enclosure contains torrenturl
type Enclosure struct {
	XMLName    xml.Name `xml:"enclosure"`
	TorrentUrl string   `xml:"url,attr"`
}
