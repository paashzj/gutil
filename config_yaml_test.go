package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadYamlSuccess(t *testing.T) {
	_, err := ConfigYamlFromFile("config_yaml_test.yaml")
	assert.Nil(t, err)
}

func TestWriteYamlSuccess(t *testing.T) {
	configYaml, err := ConfigYamlFromFile("config_yaml_test.yaml")
	assert.Nil(t, err)
	configYaml.SetBool("female", false)
	configYaml.SetBool("male", true)
	configYaml.SetInt("age", 36)
	skillMap := make(map[string]interface{})
	skillMap["java"] = "master"
	skillMap["php"] = "master"
	configYaml.SetMap("skill", skillMap)
	err = configYaml.Write("config_yaml_test_gen.yaml")
	assert.Nil(t, err)
}
