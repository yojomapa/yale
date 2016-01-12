package framework 

import (
	"errors"
	"time"
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

func (helper *Marathon) SetClient(client marathon.Marathon) {
	helper.client = client
}

func (helper *Marathon) FindServiceInformation(serviceName string) ([]*model.Instance, error) {
	app, err := helper.client.Application(serviceName)
        if err != nil {
                return nil, err
        } else {
		if app == nil {
			return make([]*model.Instance, 0), nil
		} else {
			return helper.getInstancesFromTasks(app.Tasks), nil
		}
        }
}
func (helper *Marathon) createService(config model.ServiceConfig) ([]*model.Instance, error) {
        app := translateServiceConfig(config)
        _, err := helper.client.CreateApplication(app)
        if err != nil {
                return nil, err
        } else {
		helper.client.WaitOnApplication(config.ServiceId, 10*time.Second)
                app, err := helper.client.Application(app.ID)
                if err != nil {
                        return nil, err
                } else {
                        return helper.getInstancesFromTasks(app.Tasks), nil
                }       
        }
}
func (helper *Marathon) DeployService(config model.ServiceConfig) ([]*model.Instance, error) {
	apps, err := helper.client.ListApplications(nil)
	if err != nil {
		return nil, err
	}
	if (!helper.containsApp(apps, config.ServiceId)) {
		return helper.createService(config)
	} else {
		return helper.scaleService(config.ServiceId, config.Instances)
	}
}

func (helper *Marathon) scaleService(id string, instances int) ([]*model.Instance, error){
	// TODO define if it is up or downscale
	// todo identify current instances, ver si actualizacion es necesario
	deploymentId, err := helper.client.ScaleApplicationInstances(id, instances, true)
	if err != nil {
    		util.Log.Errorf("Failed to Scale the application: %s, error: %s", id, err)
		return nil, err
	} else {
		app, err := helper.client.Application(id)
		if err != nil {
                	return nil, err
	        } else {
			return helper.getInstancesByVersion(app.Tasks, deploymentId.Version), nil
        	}
	}
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

func (helper *Marathon) getInstancesFromTasks(tasks []*marathon.Task) []*model.Instance {
        instances := make([]*model.Instance, len(tasks))

        for i, task := range tasks {
                instance := model.Instance{}
                instance.Id = task.ID
                instance.Type = "DOCKER"
                //instance.Name = task.Name
                instance.Ports = task.Ports
                instance.Node = task.Host
                //instance.State = task.
                instance.Created = task.StagedAt
                instances[i] = &instance
        }
        return instances
}

func (helper *Marathon) getInstancesByVersion(tasks []*marathon.Task, version string) []*model.Instance {
	for i, task := range tasks {
		if task.Version != version {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	return helper.getInstancesFromTasks(tasks)
}

func (m *Marathon) containsApp(apps []string, search string) bool{
	for _, a := range apps {
		if a == search {
			return true
		}
	}
	return false
}
