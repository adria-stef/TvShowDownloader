package configuration

import (
	"io/ioutil"

	"github.com/adria-stef/TvShowDownloader/model"

	"gopkg.in/yaml.v2"
)

//GetConfig returns Cofig struct containing all personal cofiguration
func GetConfig() model.Config {
	var config model.Config
	data, err := ioutil.ReadFile("./files/list.yml")
	if err != nil {
		panic("Could not read list.yml. Please verify that your configuration file is readable.")
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("Could not unmarshal list.yml. Please verify that your configuration file is properly populated.")
	}

	return config
}
