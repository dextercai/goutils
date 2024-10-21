package file

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
)

// GenEmptyYamlFile 用于生成空的Yaml配置文件
func GenEmptyYamlFile[T any](destFilePath string) error {
	cnf := new(T)
	data, err := yaml.Marshal(&cnf)
	if err != nil {
		return err
	}
	err = os.WriteFile(destFilePath, data, 0644)
	return err
}

// GenEmptyJsonFile 用于生成空的json配置文件
func GenEmptyJsonFile[T any](destFilePath string) error {
	cnf := new(T)
	data, err := json.Marshal(&cnf)
	if err != nil {
		return err
	}
	err = os.WriteFile(destFilePath, data, 0644)
	return err
}
