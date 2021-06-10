package config

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

func Init(c string) error {
	bytesContent, err := Asset(fmt.Sprintf("config/%s.toml", c))
	if err != nil {

		return err
	}
	viper.SetConfigType("toml") // 设置配置文件格式为YAML
	if err := viper.ReadConfig(bytes.NewBuffer(bytesContent)); err != nil {
		return err
	}
	return nil
}
