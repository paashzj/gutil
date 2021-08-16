package gutil

import (
	"bufio"
	"fmt"
	"os"
)

type ConfigProperties struct {
	properties map[string]string
}

func (c *ConfigProperties) Init() {
	c.properties = make(map[string]string)
}

func (c *ConfigProperties) Set(key, val string) {
	c.properties[key] = val
}

func (c *ConfigProperties) Write(dstFile string) error {
	file, err := os.OpenFile(dstFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0666))
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
