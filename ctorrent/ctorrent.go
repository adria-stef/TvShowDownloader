//Package ctorrent provides primitives for downloading torrent files with the ctorrent tool via the command line
package ctorrent

import (
	"log"
	"os/exec"
	"runtime"
)

//DownloadTvShow downloads torrent file to a specified download path by executing command
func DownloadTvShow(torrentName, downloadPath string) {

	script := getProperScript()
	tracker := "udp://tracker.opentrackr.org:1337/announce"

	cmd := exec.Command(script, downloadPath, torrentName, tracker)

	err := cmd.Run()
	if err != nil {
		log.Fatal("Error while executing download script.", err)
	}
}

//getProperScript return proper script for download depending on the platform
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
