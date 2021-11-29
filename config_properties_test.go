package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadPropertiesSuccess(t *testing.T) {
	_, err := ConfigPropFromFile("config_properties_test.conf")
	assert.Nil(t, err)
}
