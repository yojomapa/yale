package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVersionError(t *testing.T) {
        cfg := ServiceConfig {
                Tag: "_bla",
        }

	res, err := cfg.Version()
	assert.Equal(t, "", res, "Version should be empty")
	assert.NotNil(t, err, "Should throw an error")
}

func TestVersion(t *testing.T) {
        cfg := ServiceConfig {
                Tag: "1.2-Service",
        }

        res, _ := cfg.Version()
        assert.Equal(t, "1.2", res, "Version should be 1.2")
}

func TestString(t *testing.T) {
        cfg := ServiceConfig {
                Tag: "bla",
        }
	res := cfg.String()
	assert.Contains(t, res, "Tag: bla", "Config should contain Tag: bla")
}
