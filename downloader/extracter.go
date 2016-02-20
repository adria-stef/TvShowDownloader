package downloader

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

//ExtarctShow returns name and se(season-episode) from title if title is in the proper format:
//
//<name> <se> <info>
//
//If title is not in proper format it will return empty strings
func ExtarctShow(title string) (string, string) {

	title = strings.ToLower(title)

	regex := `(.*)(s[0-9]*e[0-9]*).*`
	re := regexp.MustCompile(regex)

	show := re.FindAllStringSubmatch(string(title[:]), -1)

	if len(show) < 1 || len(show[0]) < 2 {
		return "", ""
	}

	showName := show[0][1]
	se := show[0][2]

	return showName, se
}

//ExtractSeasonEpisode returns season and episode from se(season-episode) if se is in proper format:
//
//s<season>e<episode>
//
//If se is not in proper format it will return 0, 0
func ExtractSeasonEpisode(se string) (int, int) {

	seRegex := `s([0-9]*)e([0-9]*)`
	re := regexp.MustCompile(seRegex)

	result := re.FindAllStringSubmatch(string(se[:]), -1)

	if len(result) < 1 || len(result[0]) < 2 {
		return 0, 0
	}

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
