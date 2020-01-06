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

func (f *FileConfig) ParseConfig(parser config.Parser, cfg interface{}) error {
	data, err := GetBytes(f.path)
	if err != nil {
		return err
	}

	return parser.Parse(data, cfg)
}
