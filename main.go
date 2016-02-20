package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adria-stef/TvShowDownloader/downloader"

	"github.com/adria-stef/TvShowDownloader/Godeps/_workspace/src/github.com/jasonlvhit/gocron"
)

func main() {
	file, err := os.OpenFile("./testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening logging file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	s := gocron.NewScheduler()
	s.Every(30).Seconds().Do(task)
	<-s.Start()

}

func task() {
	log.Printf("Task triggered at: [%v]", time.Now())
	downloader.Download()
}
