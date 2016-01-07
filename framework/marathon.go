package framework 

import (
	"errors"
	"github.com/jglobant/yale/model"
	"github.com/jglobant/yale/util"
	"github.com/gambol99/go-marathon"
)

type Marathon struct {
	client 	marathon.Marathon
	endpointUrl string
}

func NewMarathon(endpointUrl string) (*Marathon, error) {
	helper := new(Marathon)
	helper.endpointUrl = endpointUrl
	config := marathon.NewDefaultConfig()
	config.URL = endpointUrl
	client, err := marathon.NewClient(config)
	
	if err != nil {
		return nil, err
	}
	
	helper.client = client
	return helper, nil
}

func NewMarathonTls(endpointUrl, cert,  key string) (*Marathon, error) {
	return nil, errors.New("Not implemented yet")
}

func NewMarathonTlsVerify(endpointUrl, cert,  key, ca string) (*Marathon, error) {
	return nil, errors.New("Not implemented yet")
}

func (helper *Marathon) ListServices(serviceName string) []*model.Instance {
	
	application, _ := helper.client.Application(serviceName)
	tasks := application.Tasks
	containers := make([]*model.Instance, len(tasks))
	
	for i, task := range tasks {
		containers[i].Id = task.ID
		containers[i].Type = application.Container.Type
		//containers[i].Name = task.Name
		containers[i].Ports = task.Ports
		containers[i].Node = task.Host
		//containers[i].State = task.
		containers[i].Created = task.StagedAt
	}
	return containers
}

func (helper *Marathon) DeployService(config model.ServiceConfig) (error) {
	application := translateServiceConfig(config)
	_, err := helper.client.CreateApplication(application)
	return err
}

func (helper *Marathon) ScaleService(id string, instances int) (error){
	_, err := helper.client.ScaleApplicationInstances(id, instances, true)
	if err != nil {
    		util.Log.Errorf("Failed to Scale the application: %s, error: %s", id, err)
	}
	return err
}

func (helper *Marathon) DeleteService(id string) (error) {
	_, err := helper.client.DeleteApplication(id)
	if err != nil {
		util.Log.Errorf("Failed to Delete the application: %s, error: %s", id, err)
	}
	return err
}

func (scheduler *Marathon) UndeployInstance(instance *model.Instance) error {
	return errors.New("Not implemented yet")
}
