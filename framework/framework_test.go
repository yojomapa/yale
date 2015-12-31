package framework

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestCreateFramework(t *testing.T) {
	
	fmt.Println("TestCreateFramework Starting")

	helper, error := NewFrameworkHelper("http://localhost:8081")
	
	if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	assert.Equal(t, "*framework.MarathonHelper", fmt.Sprintf("%T", helper), "Type of helper should be *framework.MarathonHelper")
}
