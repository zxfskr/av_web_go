package config

import (
	"encoding/json"
	"io/ioutil"
)

var webConfig WebConfig

// WebConfig 程序的基本配置信息
type WebConfig struct {
	Web map[string]string `json:"web"`
}

// NewConfig 构造函数
func NewConfig() (config *WebConfig, err error) {
	config = &webConfig
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &config)
	if err != nil {
		return
	}

	return
}

// GetWebConfig 返回配置信息
func GetWebConfig() WebConfig {
	return webConfig
}
