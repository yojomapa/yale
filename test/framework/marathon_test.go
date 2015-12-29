package helper

import (
	"testing"
	"fmt"
	"github.com/jglobant/yale/model"
	"net/http"
	"net/http/httptest"
)

func TestListServices(t *testing.T) {
	
	fmt.Println("TestListServices Starting")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"apps":[{"id":"/nginx","cmd":null,"args":null,"user":null,"env":{},"instances":1,"cpus":0.25,"mem":256.0,"disk":0.0,"executor":"","constraints":[],"uris":[],"storeUrls":[],"ports":[80],"requirePorts":false,"backoffSeconds":1,"backoffFactor":1.15,"maxLaunchDelaySeconds":3600,"container":{"type":"DOCKER","volumes":[],"docker":{"image":"nginx","network":"BRIDGE","portMappings":[{"containerPort":80,"hostPort":0,"servicePort":80,"protocol":"tcp"}],"privileged":false,"parameters":[],"forcePullImage":false}},"healthChecks":[],"dependencies":[],"upgradeStrategy":{"minimumHealthCapacity":1.0,"maximumOverCapacity":1.0},"labels":{},"acceptedResourceRoles":null,"version":"2015-12-28T18:55:35.598Z","versionInfo":{"lastScalingAt":"2015-12-28T18:55:35.598Z","lastConfigChangeAt":"2015-12-28T18:44:32.774Z"},"tasksStaged":0,"tasksRunning":1,"tasksHealthy":0,"tasksUnhealthy":0,"deployments":[]}]}`)
	}))
	defer ts.Close()

	helper, error := NewMarathonHelper(ts.URL)
	
	fmt.Println("NewMarathonHelper CALLED")
	
	if error != nil {
		
		fmt.Println("ERROR !!!!!: " + error.Error())
		t.Errorf("Error: " + error.Error())
	}
	
	services := helper.ListServices();
	fmt.Printf("Services Found: %d \n", len(services))
}

func TestDeployService(t *testing.T) {
ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintln(w, `{"id":"/nginx","cmd":null,"args":null,"user":null,"env":{},"instances":1,"cpus":0.25,"mem":256.0,"disk":0.0,"executor":"","constraints":[],"uris":[],"storeUrls":[],"ports":[80],"requirePorts":false,"backoffSeconds":1,"backoffFactor":1.15,"maxLaunchDelaySeconds":3600,"container":{"type":"DOCKER","volumes":[],"docker":{"image":"nginx","network":"BRIDGE","portMappings":[{"containerPort":80,"hostPort":0,"servicePort":80,"protocol":"tcp"}],"privileged":false,"parameters":[],"forcePullImage":false}},"healthChecks":[],"dependencies":[],"upgradeStrategy":{"minimumHealthCapacity":1.0,"maximumOverCapacity":1.0},"labels":{},"acceptedResourceRoles":null,"version":"2015-12-28T18:44:32.774Z","tasksStaged":0,"tasksRunning":0,"tasksHealthy":0,"tasksUnhealthy":0,"deployments":[{"id":"b37bbd63-6e92-43b6-9199-d1ca27642b55"}],"tasks":[]}`)
        }))
        defer ts.Close()
	helper, error := NewMarathonHelper(ts.URL)
	
		if error != nil {
		t.Errorf("Error: " + error.Error())
	}
	
	helper.DeployService(model.ServiceConfig{})
}
