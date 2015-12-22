package helper

import (
	"testing"
	"fmt"
	//"github.com/yojomapa/yale/model"
	)

func TestListServices(t *testing.T) {

	helper, error := NewMarathonHelper("http://localhost:8081")
	services := helper.ListServices();
	
	if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	
	fmt.Println("Services Found")
	fmt.Println(len(services))

//	RunMarathon()

}

//func TestDeployService(t *testing.T) {

//	helper, error := NewMarathonHelper("http://localhost:8081")
	
//		if error != nil {
//		t.Errorf("Error: " + error.Error())
//	}
	
//	helper.DeployService(model.ServiceConfig{})
//}