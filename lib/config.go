package lib

import (
	"encoding/json"

	"github.com/ghodss/yaml"
	"github.com/thedevsaddam/gojsonq"
)

type yamlDecoder struct{}

func (i *yamlDecoder) Decode(data []byte, v interface{}) error {
	bb, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bb, &v)
}

// 读取 yaml 配置文件，
// 根据输入的 key 值，
// 返回一个interface
func ReadConfig(key string) interface{} {
	value := gojsonq.New(gojsonq.SetDecoder(&yamlDecoder{})).File("./config/config.yaml").Find(key)
	return value
}

func ReadConfigArray(key string) ([]string, error) {
	value, err := gojsonq.New(gojsonq.SetDecoder(&yamlDecoder{})).File("./config/config.yaml").FindR(key)
	if err != nil {
		return []string{}, nil
	}

	valueStringSlice, err := value.StringSlice()
	if err != nil {
		return valueStringSlice, err
	}
	return valueStringSlice, nil
}

func ReadJsonConfig(key string) interface{} {
	value := gojsonq.New().File("./config/batch.json").Find(key)
	return value
}
