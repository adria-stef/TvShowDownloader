package downloader

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getRssFile() []byte {
	url := "https://eztv.ag/ezrss.xml"
	cli := http.Client{}
	resp, err := cli.Get(url)
	if err != nil {
		log.Fatalf("Error while getting rss feed with %vw", err)
		return nil
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("./files/rss.xml", []byte(contents), 0644)

	if err != nil {
		log.Fatalf("Error while reading rss feed response body with %v", err)
		return nil
	}
	return contents
}
