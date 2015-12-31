package framework

import (
	"testing"
	"fmt"
	"io/ioutil"
	"github.com/jglobant/yale/model"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestConstructorError(t *testing.T) {
	_, err := NewMarathonHelper("malformed url")
	assert.True(t, err != nil, "Malformed Url should throw error")
}

func TestListServices(t *testing.T) {
	
	content, _ := ioutil.ReadFile("../test/resources/marathon_tasks_response.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(content))
	}))
	defer ts.Close()

	helper, error := NewMarathonHelper(ts.URL)
	
	if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	
	services := helper.ListServices();
	assert.Equal(t, 1, len(services), "Should have found one service")
}

func TestDeployService(t *testing.T) {
	content, _ := ioutil.ReadFile("../test/resources/marathon_new_app_response.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, string(content))
        }))
        defer ts.Close()
	helper, error := NewMarathonHelper(ts.URL)

	if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	
	err := helper.DeployService(model.ServiceConfig{})
	//assert.Equal(t, "nginx", app.ID, "ID should be nginx")
	//assert.Equal(t, 1, app.Instances, "App should have 1 instance")
	assert.True(t, err == nil, "Deploy should work")
}

func TestErrorDeployService(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, "failed response")
        }))
        defer ts.Close()
        helper, error := NewMarathonHelper(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }
	err := helper.DeployService(model.ServiceConfig{})
	assert.True(t, err != nil, "Deploy should throw error")
}

func TestScaleService(t *testing.T) {
        content, _ := ioutil.ReadFile("../test/resources/marathon_update_instances_response.json")
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, string(content))
        }))
        defer ts.Close()
        helper, error := NewMarathonHelper(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }

        err := helper.ScaleService("nginx", 2)
        assert.True(t, err == nil, "Scale should work")
}

func TestErrorScaleService(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, "failed response")
        }))
        defer ts.Close()
        helper, error := NewMarathonHelper(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }
        err := helper.ScaleService("error-id", -1)
        assert.True(t, err != nil, "Scale should throw error")
}

func TestDeleteService(t *testing.T) {
        content, _ := ioutil.ReadFile("../test/resources/marathon_delete_app_response.json")
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, string(content))
        }))
        defer ts.Close()
        helper, error := NewMarathonHelper(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }

        err := helper.DeleteService("nginx")
        assert.True(t, err == nil, "Delete should work")
}

func TestErrorDeleteService(t *testing.T) {
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, "failed response")
        }))     
        defer ts.Close()
        helper, error := NewMarathonHelper(ts.URL)
        
        if error != nil {
                t.Errorf("Error: " + error.Error())
        }       
        err := helper.DeleteService("error-id")
        assert.True(t, err != nil, "Delete should throw error")
}
