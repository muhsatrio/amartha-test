package yaml

import (
	"golang-boilerplate/platform/jwt"
	"golang-boilerplate/platform/mysql"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Auth       jwt.Config `yaml:"auth"`
	DataSource DataSource `yaml:"data_source"`
}

type DataSource struct {
	MySQL mysql.ConfigDB `yaml:"mysql"`
}

func Open(fileName string) (config Config, err error) {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}

	return
}
