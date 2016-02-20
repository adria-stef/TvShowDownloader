package downloader

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func extarctShow(title string) (string, string) {

	title = strings.ToLower(title)

	regex := `(.*)(s[0-9]*e[0-9]*).*`
	re := regexp.MustCompile(regex)

	show := re.FindAllStringSubmatch(string(title[:]), -1)

	showName := show[0][1]
	se := show[0][2]

	return showName, se
}

func extractSeasonEpisode(se string) (int, int) {

	seRegex := `s([0-9]*)e([0-9]*)`
	re := regexp.MustCompile(seRegex)

	result := re.FindAllStringSubmatch(string(se[:]), -1)

	seasonString := result[0][1]
	episodeString := result[0][2]

	seasonInt, err := strconv.Atoi(seasonString)
	if err != nil {
		log.Println("Could not get season number")
	}

	episodeInt, err := strconv.Atoi(episodeString)
	if err != nil {
		log.Println("Could not get episode number")
	}

	return seasonInt, episodeInt
}
