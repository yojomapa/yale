package cluster

import (
	"github.com/jglobant/yale/framework"
	"github.com/stretchr/testify/assert"
	"testing"
	"reflect"
)

func TestConstructor(t *testing.T) {
	sm := NewStackManager()
	assert.True(t, sm != nil, "Instance should be healthy")
	cfg := framework.FrameworkConfig{
		EndpointUrl: "http://localhost:8081",
		Type:        framework.MARATHON,
	}
	helper, _ := framework.NewFramework(cfg)
	sm.AppendStack(helper)
	v := reflect.ValueOf(sm).Elem()
	stacks := v.FieldByName("stacks")
	assert.Equal(t, 1, stacks.Len(), "Cli should instantiate at least one stack")
}
