package source

import (
	"config"
	"io/ioutil"
	"os"
)

type FileConfig struct {
	path string
}

func NewFileConfig(path string) *FileConfig {
	return &FileConfig{path: path}
}

func GetBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath) // For read access.
	defer file.Close()

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

func (f *FileConfig) ParseConfig(t config.ParserType, cfg interface{}) error {
	data, err := GetBytes(f.path)
	if err != nil {
		return err
	}

	return config.GetParser(t).Parse(data, cfg)
}
