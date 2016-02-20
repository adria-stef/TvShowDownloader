//Package downloader provides primitives for downloading
package downloader

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/adria-stef/TvShowDownloader/config"
	"github.com/adria-stef/TvShowDownloader/ctorrent"
	"github.com/adria-stef/TvShowDownloader/database"
	"github.com/adria-stef/TvShowDownloader/model"
)

//Download starts the Download task.
//Download task consits of the following steps:
//
//1. Get the torrent file
//
//2. Download the actual file from torent
//
//3. Add download info to db for future references
func Download() {
	config := configuration.GetConfig()
	dbFilePath := "./files/bolt.db"
	downloadPath := initDownloadPath(config.DownloadPath)
	fmt.Printf("Check start: %s\n", time.Now())

	itemsForDownload := getItemsForDownload(dbFilePath)

	for _, item := range itemsForDownload {
		if notInQueue(item.Title, downloadPath) {
			fmt.Printf("Downloading %s file...\n", changeSpaces(item.Title))

			torrentName := getTorrentFile(item, downloadPath)
			ctorrent.DownloadTvShow(torrentName, downloadPath)
			addToDB(item.Title, dbFilePath)
		}
	}
}

//initDownloadPath returns path that is going to be used for downloads
//
//if one is not specified in the configuration (files/list.yml) default location is: /torrents directory
func initDownloadPath(downloadPath string) string {
	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		dir := os.Getenv("PWD")
		downloadPath = fmt.Sprintf("%s%s", dir, "/torrents/")
	}

	if i := strings.LastIndex(downloadPath, "/"); i != len(downloadPath) {
		downloadPath = fmt.Sprintf("%s/", downloadPath)
	}

	return downloadPath
}

//notInQueue returns true if the same episode has not yet been downloaded
func notInQueue(title, path string) bool {
	files, _ := filepath.Glob(fmt.Sprintf("%s*", path))

	showName, se := ExtarctShow(title)

	for _, value := range files {
		valueToLowerCase := strings.ToLower(value)

		if strings.Contains(valueToLowerCase, changeSpaces(strings.Trim(showName, " "))) && strings.Contains(valueToLowerCase, strings.Trim(se, " ")) {
			return false
		}
	}
	return true
}

//getTrrentFile downloads torrent file to specified path and returns the name of the file
func getTorrentFile(item model.Item, downloadPath string) string {

	cli := http.Client{}
	resp, err := cli.Get(item.Enclosure.TorrentUrl)
	if err != nil {
		log.Printf("Error while getting torrent file [%s] with %v", item.Title, err)
		return ""
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	fileName := fmt.Sprintf("%s.torrent", changeSpaces(item.Title))
	ioutil.WriteFile(fmt.Sprintf("%s%s", downloadPath, fileName), []byte(contents), 0644)

	return fileName
}

//changeSpaces turns spaces to dots
func changeSpaces(title string) string {
	return strings.Replace(title, " ", ".", len(title))
}

//addToDB adds a key-value pair title-last season and last episode to DB
func addToDB(itemTitle, dbFilePath string) {
	db := database.GetDB(dbFilePath)
	defer db.Close()

	titleToLowerCase := strings.ToLower(itemTitle)

	showName, se := ExtarctShow(titleToLowerCase)
	database.StoreData(db, []byte(showName), []byte(se))
}
