package helper

import (
	"fmt"
	"github.com/yojomapa/yale/model"
	"github.com/yojomapa/yale/util"
	"github.com/gambol99/go-marathon"
)

type MarathonHelper struct {
	client 	marathon.Marathon
	endpointUrl string
}

func NewMarathonHelper(endpointUrl string) (*MarathonHelper, error) {
	helper := new(MarathonHelper)
	helper.endpointUrl = endpointUrl
	config := marathon.NewDefaultConfig()
	client, err := marathon.NewClient(config)
	
	if err != nil {
	    fmt.Println("Failed to create a client for marathon, error: %s", err)
		return nil, err
	}
	
	helper.client = client
	return helper, nil
}

func (helper *MarathonHelper) ListServices() []string {
	
	applications, _ := helper.client.Applications(nil)
	marathonApps := applications.Apps
	appList := make([]string, len(marathonApps))
	
	for i, app := range marathonApps {
		appList[i] = app.ID
	}
	return appList
}

func (helper *MarathonHelper) DeployService(config model.ServiceConfig) {
	
	application := translateServiceConfig(config)
	
	if _, err := helper.client.CreateApplication(application); err != nil {
    	util.Log.Fatalf("Failed to create application: %s, error: %s", application, err)
	} else {
    	util.Log.Printf("Created the application: %s", application)
	}
	
}

func (helper *MarathonHelper) ScaleService(id string, instances int) {
	
}

func (helper *MarathonHelper) DeleteService(id string) {
	
}

func translateServiceConfig(config model.ServiceConfig) *marathon.Application {
	application := marathon.NewDockerApplication()
	
	application.Name("/product/name/frontend")
	application.CPU(0.1).Memory(64).Storage(0.0).Count(2)
	application.Arg("/usr/sbin/apache2ctl", "-D", "FOREGROUND")
	application.AddEnv("NAME", "frontend_http")
	application.AddEnv("SERVICE_80_NAME", "test_http")
	application.AddLabel("environment", "staging")
	application.AddLabel("security", "none")
	// add the docker container
	application.Container.Docker.Container("quay.io/gambol99/apache-php:latest").Expose(80, 443)
	application.CheckHTTP("/health", 10, 5)
	
	return application
}

func RunMarathon() {
	marathonURL := "http://localhost:8081"
	config := marathon.NewDefaultConfig()
	config.URL = marathonURL
	client, err := marathon.NewClient(config)
	if err != nil {
	    fmt.Println("Failed to create a client for marathon, error: %s", err)
	}
	
	applications, _ := client.Applications(nil)
	fmt.Println(applications.Apps)
}