package downloader

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/adria-stef/TvShowDownloader/cmd"
	"github.com/adria-stef/TvShowDownloader/config"
	"github.com/adria-stef/TvShowDownloader/database"
	"github.com/adria-stef/TvShowDownloader/model"
)

func Download() {
	config := configuration.GetConfig()
	downloadPath := initDownloadPath(config.DownloadPath)
	fmt.Printf("Download path: %s", downloadPath)

	itemsForDownload := getItemsForDownload()

	for _, item := range itemsForDownload {
		if notInQueue(item.Title, downloadPath) {
			fmt.Printf("Downloading %s file...\n", changeSpaces(item.Title))
			torrentName := getTorrentFile(item, downloadPath)
			//TODO check if  item's torrent is downloaded successfully
			cmd.DownloadTvShow(torrentName, downloadPath)
			//TODO check if file downloaded successfully
			addToDB(item.Title)
		}
	}
}

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

func notInQueue(title, path string) bool {
	files, _ := filepath.Glob(fmt.Sprintf("%s*", path))

	showName, se := extarctShow(title)

	for _, value := range files {
		valueToLowerCase := strings.ToLower(value)

		if strings.Contains(valueToLowerCase, changeSpaces(strings.Trim(showName, " "))) && strings.Contains(valueToLowerCase, strings.Trim(se, " ")) {
			return false
		}
	}
	return true
}

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

func changeSpaces(title string) string {
	return strings.Replace(title, " ", ".", len(title))
}

func addToDB(itemTitle string) {
	db := database.GetDB()
	defer db.Close()

	titleToLowerCase := strings.ToLower(itemTitle)

	showName, se := extarctShow(titleToLowerCase)
	database.StoreData(db, []byte(showName), []byte(se))
}
