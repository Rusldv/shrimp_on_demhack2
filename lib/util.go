package lib

import (
	"encoding/json"
	"github.com/rusldv/kit/fileutil"
)

// Config main node loaded in the struct from configurate file.
type Config struct {
	Port                string `json:"port"`
	SSLPort             string `json:"ssl_port"`
	RootDir             string `json:"root_dir"`
	IndexComponent      string `json:"index_component"`
	FileOptions         string `json:"file_options"`
	FileController      string `json:"file_controller"`
	FileView            string `json:"file_view"`
	OptionsExtension    string `json:"options_extension"`
	ControllerExtension string `json:"controller_extension"`
	ViewExtension       string `json:"view_extension"`
}

// ParseConfig читает файл конфига и возращает в виде объекта
func ParseConfig(path string) (*Config, error) {
	f, err := fileutil.ReadFileString(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal([]byte(f), &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
