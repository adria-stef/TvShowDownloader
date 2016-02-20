//Package main is where the magic happens
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adria-stef/TvShowDownloader/downloader"

	"github.com/jasonlvhit/gocron"
)

func main() {

	var minutes = flag.Uint64("minites", 60, "how often would you like to check for new tv shows (in minutes)")
	flag.Parse()

	file, err := os.OpenFile("./testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening logging file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	s := gocron.NewScheduler()
	s.Every(*minutes).Seconds().Do(task)
	<-s.Start()

}

func task() {
	log.Printf("Task triggered at: [%v]", time.Now())
	downloader.Download()
}
