package downloader

import (
	"regexp"
	"strings"

	"github.com/adria-stef/TvShowDownloader/config"
	"github.com/adria-stef/TvShowDownloader/database"
	"github.com/adria-stef/TvShowDownloader/model"
)

//getItemsForDownload returns []model.Items in rssfeed that are both in the list of tv show and have note yet been watched
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

//getNewItems returns []model.Item containg all new model.Item-
func getNewItems() []model.Item {
	rssFile := getRssFile()
	rssFeed := configuration.GetRssFeed(rssFile)

	return rssFeed.Channel.Items
}

//isInMyList returns true if title is in list of tv show followed
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

//haveNotWatchedIt returns true if this episode hase not yet been downloaded
func haveNotWatchedIt(title, dbFilePath string) bool {
	db := database.GetDB(dbFilePath)
	defer db.Close()

	title = strings.ToLower(title)
	properFormat, _ := regexp.MatchString(`(.*)(s[0-9]*e[0-9]*).*`, title)

	if !properFormat {
		return false
	}

	showName, se := ExtarctShow(title)
	currentSeason, currentEpisode := ExtractSeasonEpisode(se)

	lastSE := database.GetValue(db, []byte(strings.ToLower(showName)))
	if lastSE == nil {

		return true
	}

	lastSeason, lastEpisode := ExtractSeasonEpisode(string(lastSE))

	if lastSeason < currentSeason {
		return true
	} else if lastEpisode < currentEpisode {
		return true
	}

	return false
}
