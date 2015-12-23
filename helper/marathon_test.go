package helper

import (
	"testing"
	"fmt"
	"github.com/yojomapa/yale/model"
	)

func TestListServices(t *testing.T) {
	
	fmt.Println("TestListServices Starting")

	helper, error := NewMarathonHelper("http://localhost:8081")
	
	fmt.Println("NewMarathonHelper CALLED")
	
	if error != nil {
		
		fmt.Println("ERROR !!!!!: " + error.Error())
		t.Errorf("Error: " + error.Error())
	}
	
	services := helper.ListServices();
	fmt.Printf("Services Found: %d \n", len(services))
}

func TestDeployService(t *testing.T) {

	helper, error := NewMarathonHelper("http://localhost:8081")
	
		if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	
	helper.DeployService(model.ServiceConfig{})
}