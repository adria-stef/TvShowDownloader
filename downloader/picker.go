package downloader

import (
	"regexp"
	"strings"

	"github.com/adria-stef/TvShowDownloader/config"
	"github.com/adria-stef/TvShowDownloader/database"
	"github.com/adria-stef/TvShowDownloader/model"
)

func getItemsForDownload(dbFilePath string) []model.Item {
	items := getNewItems()

	var itemsForDownload []model.Item

	for _, currentItem := range items {

		title := currentItem.Title

		if isInMyList(title) && haveNotWatchedIt(title, dbFilePath) {
			itemsForDownload = append(itemsForDownload, currentItem)
		}
	}

	return itemsForDownload
}

func getNewItems() []model.Item {
	rssFile := getRssFile()
	rssFeed := configuration.GetRssFeed(rssFile)

	return rssFeed.Channel.Items
}

func isInMyList(title string) bool {
	config := configuration.GetConfig()

	myShows := config.List

	for _, showName := range myShows {
		if strings.Contains(strings.ToLower(title), strings.ToLower(showName)) {
			return true
		}
	}

	return false
}

func haveNotWatchedIt(title, dbFilePath string) bool {
	db := database.GetDB(dbFilePath)
	defer db.Close()

	title = strings.ToLower(title)
	properFormat, _ := regexp.MatchString(`(.*)(s[0-9]*e[0-9]*).*`, title)

	if !properFormat {
		return false
	}

	showName, se := extarctShow(title)
	currentSeason, currentEpisode := extractSeasonEpisode(se)

	lastSE := database.GetValue(db, []byte(strings.ToLower(showName)))
	if lastSE == nil {

		return true
	}

	lastSeason, lastEpisode := extractSeasonEpisode(string(lastSE))

	if lastSeason < currentSeason {
		return true
	} else if lastEpisode < currentEpisode {
		return true
	}

	return false
}
