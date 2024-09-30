package file

import (
	"gopkg.in/yaml.v3"
	"os"
)

// GenEmptyYamlFile 用于生成空的Yaml配置文件
func GenEmptyYamlFile[T comparable](destFilePath string) error {
	cnf := new(T)
	data, err := yaml.Marshal(&cnf)
	if err != nil {
		return err
	}
	err = os.WriteFile(destFilePath, data, 0644)
	return err
}
