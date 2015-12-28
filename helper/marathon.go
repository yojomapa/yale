package helper

import (
	"fmt"
	"github.com/jglobant/yale/model"
	"github.com/jglobant/yale/util"
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
	config.URL = endpointUrl
	client, err := marathon.NewClient(config)
	
	if err != nil {
	   util.Log.Fatalf("Failed to create a client for marathon, error: %s", err)
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
	if _, err := helper.client.ScaleApplicationInstances(id, instances, true); err != nil {
    	util.Log.Fatalf("Failed to Scale the application: %s, error: %s", id, err)
	}
}

func (helper *MarathonHelper) DeleteService(id string) {
	if _, err := helper.client.DeleteApplication(id); err != nil {
		util.Log.Fatalf("Failed to Delete the application: %s, error: %s", id, err)
	}
}

func translateServiceConfig(config model.ServiceConfig) *marathon.Application {
	application := marathon.NewDockerApplication()
	imageWithTag := config.ImageName + ":" + config.Tag
	labels := map[string]string{
		"image_name": config.ImageName,
		"image_tag":  config.Tag,
	}
	
	application.ID = config.ServiceId
	application.Name(imageWithTag)
	application.CPU(0.1) // how to map this ?
	application.Memory(float64(config.Memory))
	application.Count(config.Instances)
	//application.Arg("/usr/sbin/apache2ctl", "-D", "FOREGROUND")
	application.Env = util.StringSlice2Map(config.Envs)
	application.Labels = labels
	// add the docker container
	application.Container.Docker.Container(imageWithTag)
	//application.Container.Docker.Expose(80, 443)
	application.Container.Docker.PortMappings = createPorMappings(config.Publish) // Hard to map!!
	//application.CheckHTTP("/health", 10, 5)
	return application
}

func createPorMappings(ports []string) []*marathon.PortMapping {
	
	for _, val := range ports {
		fmt.Println(val)
		
//		var portMap marathon.PortMamarathon.PortMapping
//		portMap
		
	}
	
	return nil
}

//func RunMarathon() {
//	marathonURL := "http://localhost:8081"
//	config := marathon.NewDefaultConfig()
//	config.URL = marathonURL
//	client, err := marathon.NewClient(config)
//	if err != nil {
//	    fmt.Println("Failed to create a client for marathon, error: %s", err)
//	}
	
//	applications, _ := client.Applications(nil)
//	fmt.Println(applications.Apps)
//}