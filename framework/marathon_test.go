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
	_, err := NewMarathon("malformed url")
	assert.True(t, err != nil, "Malformed Url should throw error")
}

func TestListServices(t *testing.T) {
	
	ts := setup("../test/resources/marathon_tasks_response.json")
	defer ts.Close()

	m, error := NewMarathon(ts.URL)
	
	if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	
	services := m.ListServices("nginx");
	assert.Equal(t, 2, len(services), "Should have found two services")
}

func TestDeployService(t *testing.T) {
	ts := setup("../test/resources/marathon_new_app_response.json")
        defer ts.Close()
	helper, error := NewMarathon(ts.URL)

	if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	
	err := helper.DeployService(model.ServiceConfig{})
	assert.True(t, err == nil, "Deploy should work")
}

func TestErrorDeployService(t *testing.T) {
	ts := setupFailure()
        defer ts.Close()
        helper, error := NewMarathon(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }
	err := helper.DeployService(model.ServiceConfig{})
	assert.True(t, err != nil, "Deploy should throw error")
}

func TestScaleService(t *testing.T) {
	ts := setup("../test/resources/marathon_update_instances_response.json")
        defer ts.Close()
        helper, error := NewMarathon(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }

        err := helper.ScaleService("nginx", 2)
        assert.True(t, err == nil, "Scale should work")
}

func TestErrorScaleService(t *testing.T) {
        ts := setupFailure()
	defer ts.Close()
        helper, error := NewMarathon(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }
        err := helper.ScaleService("error-id", -1)
        assert.True(t, err != nil, "Scale should throw error")
}

func TestDeleteService(t *testing.T) {
        ts := setup("../test/resources/marathon_delete_app_response.json")
	defer ts.Close()
	helper, error := NewMarathon(ts.URL)

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }

        err := helper.DeleteService("nginx")
        assert.True(t, err == nil, "Delete should work")
}

func TestErrorDeleteService(t *testing.T) {
	ts := setupFailure()
	defer ts.Close()
        helper, error := NewMarathon(ts.URL)
        
        if error != nil {
                t.Errorf("Error: " + error.Error())
        }       
        err := helper.DeleteService("error-id")
        assert.True(t, err != nil, "Delete should throw error")
}

func setupFailure() (*httptest.Server){
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, "failed response")
        }))
	return ts
}

func setup(url string) (*httptest.Server) {
        content, _ := ioutil.ReadFile(url)
        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, string(content))
        }))
	return ts
}

func TestUndeployInstance(t *testing.T) {
        m, error := NewMarathon("http://localhost:8081")

        if error != nil {
                t.Errorf("Error: " + error.Error())
        }
	i := model.Instance{}
        err := m.UndeployInstance(&i)
        assert.True(t, err == nil, "Delete should work")
}
