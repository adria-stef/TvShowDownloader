package cmd

import (
	"log"
	"os/exec"
	"runtime"
)

func DownloadTvShow(torrentName, downloadPath string) {

	script := getProperScript()
	tracker := "udp://tracker.opentrackr.org:1337/announce"

	cmd := exec.Command(script, downloadPath, torrentName, tracker)

	err := cmd.Run()
	if err != nil {
		log.Fatal("Error while executing download script.", err)
	}
}

func getProperScript() string {
	var script string
	switch runtime.GOOS {
	case "windows":
		script = "./scripts/download_windows.sh"
	case "darwin":
		script = "./scripts/download_darwin.sh"
	default:
		script = "./scripts/download.sh"
	}

	log.Printf("Downloading using the %s script", script)
	return script
}
