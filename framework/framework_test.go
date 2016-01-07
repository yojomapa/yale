package framework

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestCreateFramework(t *testing.T) {
	cfg := FrameworkConfig{
		EndpointUrl : "http://localhost:8081",
		Type : MARATHON,
	}
	helper, error := NewFramework(cfg)
	assert.True(t, error == nil, "Helper should not return error")
	assert.Equal(t, "*framework.Marathon", fmt.Sprintf("%T", helper), "Type of helper should be *framework.Marathon")
}

func TestCreateNotSupportedFramework(t *testing.T) {
        cfg := FrameworkConfig{
                EndpointUrl : "http://localhost:8081",
                Type : SWARM,
        }
        _, error := NewFramework(cfg)
        assert.NotNil(t, error, "Should throw error because Framework is not supported yet")
        _, error = NewFrameworkTlsHelper(cfg)
        assert.NotNil(t, error, "Should throw error because Framework is not supported yet")
        _, error = NewFrameworkTlsVerifyHelper(cfg)
        assert.NotNil(t, error, "Should throw error because Framework is not supported yet")
}


func TestCreateFrameworkTls(t *testing.T) {
        cfg := FrameworkConfig{
                EndpointUrl : "http://localhost:8081",
                Type : MARATHON,
		Cert : "cert-path",
		Key : "key-path",
        }
        _, error := NewFrameworkTlsHelper(cfg)
	assert.True(t, error != nil, "Helper is not implemented yet")
}

func TestCreateFrameworkTlsVerify(t *testing.T) {
        cfg := FrameworkConfig{
                EndpointUrl : "http://localhost:8081",
                Type : MARATHON,
                Cert : "cert-path",
                Key : "key-path",
		Ca : "ca-path",
        }
        _, error := NewFrameworkTlsVerifyHelper(cfg)
	assert.True(t, error != nil, "Helper is not implemented yet")
}
