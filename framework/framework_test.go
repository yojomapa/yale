package framework

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestCreateFramework(t *testing.T) {
	helper, error := NewFrameworkHelper("http://localhost:8081")
	assert.True(t, error == nil, "Helper should not return error")
	assert.Equal(t, "*framework.MarathonHelper", fmt.Sprintf("%T", helper), "Type of helper should be *framework.MarathonHelper")
}

func TestCreateFrameworkTls(t *testing.T) {
        _, error := NewFrameworkTlsHelper("http://localhost:8081", "cert-path", "key-path")
	assert.True(t, error != nil, "Helper is not implemented yet")
}

func TestCreateFrameworkTlsVerify(t *testing.T) {
        _, error := NewFrameworkTlsVerifyHelper("http://localhost:8081", "cert-path", "key-path", "ca-path")
	assert.True(t, error != nil, "Helper is not implemented yet")
}
