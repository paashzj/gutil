package gutil

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type ConfigProperties struct {
	properties map[string]string
}

func ConfigPropFromFile(file string) (*ConfigProperties, error) {
	configProp := ConfigProperties{}
	configProp.Init()
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(fileBytes), "\n")
	for _, line := range split {
		trimLine := strings.TrimSpace(line)
		if len(trimLine) == 0 {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		array := strings.Split(line, "=")
		if len(array) != 2 {
			return nil, errors.New(fmt.Sprintf("line %s has multi = ", line))
		}
		configProp.Set(array[0], array[1])
	}
	return &configProp, nil
}

func (c *ConfigProperties) Init() {
	c.properties = make(map[string]string)
}

func (c *ConfigProperties) Set(key, val string) {
	c.properties[key] = val
}

func (c *ConfigProperties) SetBool(key string, val bool) {
	if val {
		c.properties[key] = "true"
	} else {
		c.properties[key] = "false"
	}
}

func (c *ConfigProperties) SetInt(key string, val int) {
	c.properties[key] = strconv.Itoa(val)
}

func (c *ConfigProperties) SetInt32(key string, val int32) {
	c.properties[key] = strconv.Itoa(int(val))
}

func (c *ConfigProperties) SetInt64(key string, val int64) {
	c.properties[key] = strconv.FormatInt(val, 10)
}

func (c *ConfigProperties) Write(dstFile string) error {
	file, err := os.OpenFile(dstFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0640))
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for k, v := range c.properties {
		_, err = writer.WriteString(fmt.Sprintf(k + "=" + v + "\n"))
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	return nil
}
