package gutil

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type ConfigYaml struct {
	data map[string]interface{}
}

func ConfigYamlFromFile(file string) (*ConfigYaml, error) {
	configYaml := ConfigYaml{}
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(fileBytes, &configYaml.data)
	if err != nil {
		return nil, err
	}
	return &configYaml, nil
}

func (c *ConfigYaml) Set(key, val string) {
	c.data[key] = val
}

func (c *ConfigYaml) SetBool(key string, val bool) {
	c.data[key] = val
}

func (c *ConfigYaml) SetInt(key string, val int) {
	c.data[key] = val
}

func (c *ConfigYaml) SetInt32(key string, val int32) {
	c.data[key] = val
}

func (c *ConfigYaml) SetInt64(key string, val int64) {
	c.data[key] = val
}

func (c *ConfigYaml) SetMap(key string, val map[string]interface{}) {
	c.data[key] = val
}

func (c *ConfigYaml) Write(dstFile string) error {
	file, err := os.OpenFile(dstFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0640))
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := yaml.Marshal(&c.data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dstFile, data, os.FileMode(0640))
	return err
}

func (c *ConfigYaml) Init() {
	c.data = make(map[string]interface{})
}
