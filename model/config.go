//Package model provides the model
package model

//Config dexcribes the configuration of the app
type Config struct {
	DownloadPath string   `yaml:"download_path"`
	List         []string `yaml:"list"`
}
