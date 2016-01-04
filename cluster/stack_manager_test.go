package cluster

import (
	"github.com/jglobant/yale/framework"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestConstructor(t *testing.T) {
	sm:= NewStackManager()
	assert.True(t, sm != nil, "Instance should be healthy")
	helper, _ := framework.NewFrameworkHelper("http://localhost:8081")
	sm.AppendStack(&helper)
}
