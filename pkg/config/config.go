package config

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	Target interface{}
	Path   string
	Module string
	Env    string
}

func ReadModuleConfig(cfg *Cfg) error {

	getFormatFile := filePath(cfg.Path)

	switch getFormatFile {
	case ".json":
		fname := cfg.Path + "/" + cfg.Module + "." + cfg.Env + getFormatFile

		jsonFile, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}

		return json.Unmarshal(jsonFile, cfg.Target)
	default:
		fname := cfg.Path + "/" + cfg.Module + "." + cfg.Env + getFormatFile
		yamlFile, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(yamlFile, cfg)
	}

}

func filePath(root string) string {
	var file string

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		file = filepath.Ext(info.Name())
		return nil
	})

	return file
}
