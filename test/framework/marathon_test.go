package test

import (
	"testing"
	"fmt"
	"io/ioutil"
	"github.com/jglobant/yale/model"
	"github.com/jglobant/yale/framework"
	"net/http"
	"net/http/httptest"
)

func TestListServices(t *testing.T) {
	
	fmt.Println("TestListServices Starting")
	content, _ := ioutil.ReadFile("../resources/marathon_tasks_response.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(content))
	}))
	defer ts.Close()

	helper, error := framework.NewMarathonHelper(ts.URL)
	
	fmt.Println("NewMarathonHelper CALLED")
	
	if error != nil {
		
		fmt.Println("ERROR !!!!!: " + error.Error())
		t.Errorf("Error: " + error.Error())
	}
	
	services := helper.ListServices();
	fmt.Printf("Services Found: %d \n", len(services))
}

func TestDeployService(t *testing.T) {
fmt.Println("yeah - 0")
	content, _ := ioutil.ReadFile("../resources/marathon_new_app_response.json")
fmt.Println("yeah")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, string(content))
        }))
        defer ts.Close()
	helper, error := framework.NewMarathonHelper(ts.URL)

	fmt.Println("Marathon created")
	if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	
	helper.DeployService(model.ServiceConfig{})
}
