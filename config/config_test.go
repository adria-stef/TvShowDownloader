package configuration_test

import (
	. "github.com/adria-stef/TvShowDownloader/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Describe("Getting rssFeed", func() {

		Context("from valid rss.xml", func() {
			validRssXML := `
			<?xml version="1.0" encoding="UTF-8"?>
			<rss version="2.0" xmlns:torrent="http://xmlns.ezrss.it/0.1/">
			<channel>
					<title>TV Torrents RSS feed - EZTV</title>
					<link>https://eztv.ag/</link>
					<description>V Torrents RSS feed - EZTV</description>
					<item>
							<title>Last Man Standing US S05E16 720p HDTV x264-AVS</title>
							<category>TV</category>
							<author>https://eztv.ag/</author>
							<link>https://eztv.ag/ep/141856/last-man-standing-us-s05e16-720p-hdtv-x264-avs/</link>
							<guid>https://eztv.ag/ep/141856/last-man-standing-us-s05e16-720p-hdtv-x264-avs/</guid>
							<pubDate>Fri, 19 Feb 2016 19:44:23 -0600</pubDate>
							<torrent:contentLength>515962450</torrent:contentLength>
							<torrent:infoHash>8B2A7E455D06141529255BF9A1CEC740029CF192</torrent:infoHash>
							<torrent:magnetURI><![CDATA[magnet:?xt=urn:btih:8B2A7E455D06141529255BF9A1CEC740029CF192&dn=Last.Man.Standing.US.S05E16.720p.HDTV.x264-AVS%5Beztv%5D.mkv&tr=udp%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=udp%3A%2F%2Fopen.demonii.com%3A1337&tr=http%3A%2F%2Ftracker.trackerfix.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Fexodus.desync.com%3A6969]]></torrent:magnetURI>
							<torrent:seeds>0</torrent:seeds>
							<torrent:peers>0</torrent:peers>
							<torrent:verified>0</torrent:verified>
							<torrent:fileName>Last.Man.Standing.US.S05E16.720p.HDTV.x264-AVS[eztv].mkv</torrent:fileName>
							<enclosure url="https://zoink.ch/torrent/Last.Man.Standing.US.S05E16.720p.HDTV.x264-AVS[eztv].mkv.torrent" length="515962450" type="application/x-bittorrent" />
					</item>
					</channel>
			</rss>
		`
			It("should be succeed", func() {
				rssFeed := GetRssFeed([]byte(validRssXML))
				Expect(rssFeed.Channel.Items[0].Category).To(Equal("TV"))
			})
		})

		Context("from invalid rss.xml", func() {
			invalidRssXML := `
		<?xml version="1.0" encoding="UTF-8"?>
		<rss version="2.0" xmlns:torrent="http://xmlns.ezrss.it/0.1/">
		<channel>
				<title>TV Torrents RSS feed - EZTV</title>
				<link>https://eztv.ag/</link>
				<description>V Torrents RSS feed - EZTV</description>
				<item>
						<title>Last Man Standing US S05E16 720p HDTV x264-AVS</title>
						<category>TV</category>
						<author>https://eztv.ag/</author>
						<link>https://eztv.ag/ep/141856/last-man-standing-us-s05e16-720p-hdtv-x264-avs/</link>
						<guid>https://eztv.ag/ep/141856/last-man-standing-us-s05e16-720p-hdtv-x264-avs/</guid>
						<pubDate>Fri, 19 Feb 2016 19:44:23 -0600</pubDate>
						<torrent:contentLength>515962450</torrent:contentLength>
						<torrent:infoHash>8B2A7E455D06141529255BF9A1CEC740029CF192</torrent:infoHash>
						<torrent:magnetURI><![CDATA[magnet:?xt=urn:btih:8B2A7E455D06141529255BF9A1CEC740029CF192&dn=Last.Man.Standing.US.S05E16.720p.HDTV.x264-AVS%5Beztv%5D.mkv&tr=udp%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=udp%3A%2F%2Fopen.demonii.com%3A1337&tr=http%3A%2F%2Ftracker.trackerfix.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Fexodus.desync.com%3A6969]]></torrent:magnetURI>
						<torrent:seeds>0</torrent:seeds>
						<torrent:peers>0</torrent:peers>
						<torrent:verified>0</torrent:verified>
						<torrent:fileName>Last.Man.Standing.US.S05E16.720p.HDTV.x264-AVS[eztv].mkv</torrent:fileName>
						<enclosure url="https://zoink.ch/torrent/Last.Man.Standing.US.S05E16.720p.HDTV.x264-AVS[eztv].mkv.torrent" length="515962450" type="application/x-bittorrent" />
		</rss>
		`
			It("should error", func() {
				rssFeed := GetRssFeed([]byte(invalidRssXML))
				Expect(rssFeed.Channel.Items).To(BeNil())
			})
		})
	})
})
