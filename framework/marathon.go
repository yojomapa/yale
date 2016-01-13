package framework 

import (
	"errors"
	"time"
	"github.com/jglobant/yale/util"
	"github.com/gambol99/go-marathon"
	"strings"
	"fmt"
)

const frameworkID = "marathon"

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

func (helper *Marathon) FindServiceInformation(criteria ServiceInformationCriteria) ([]*ServiceInformation, error) {
	//by default this method does not return task-info
	apps, err := helper.client.Applications(nil)
        if err != nil {
                return nil, err
        } else {
		if apps == nil {
			return make([]*ServiceInformation, 0), nil
		} else {
			allServices := make([]*ServiceInformation, len(apps.Apps))
			for i, app := range apps.Apps {
				allServices[i] = helper.getServiceInformationFromApp(&app)
			}
			filteredServices := criteria.MeetCriteria(allServices)
			if filteredServices == nil || len(filteredServices) == 0 {
				return nil, errors.New("No services found")
			}
			// request task info of app, that is why we do the loop again
			services := make([]*ServiceInformation, len(filteredServices))
			for i, service := range filteredServices {
				app, err := helper.client.Application(service.ID)
				if err != nil {
					return nil, errors.New("Error listing filtered services")
				}
				services[i] = helper.getServiceInformationFromApp(app)
			}
			return services, nil
		}
        }
}

func (m *Marathon) getServiceInformationFromApp(app *marathon.Application) *ServiceInformation {
	service := ServiceInformation{ImageTag : "latest",}
	service.ID = app.ID
	imageInfo := strings.Split(app.Container.Docker.Image, ":")
	service.ImageName = imageInfo[0]
	if len(imageInfo) > 1 {
		service.ImageTag = imageInfo[1]
	}
	service.Instances = m.getInstancesFromTasks(app.Tasks, app.Container.Docker.PortMappings)
	return &service
}

func (helper *Marathon) createService(config ServiceConfig, instances int) (*ServiceInformation, error) {
        app := translateServiceConfig(config, instances)
        _, err := helper.client.CreateApplication(app)
        if err != nil {
                return nil, err
        } else {
		helper.client.WaitOnApplication(config.ServiceID, 10*time.Second)
                app, err := helper.client.Application(app.ID)
                if err != nil {
                        return nil, err
                } else {
                        return helper.getServiceInformationFromApp(app), nil
                }       
        }
}
func (helper *Marathon) DeployService(config ServiceConfig, instances int) (*ServiceInformation, error) {
	apps, err := helper.client.ListApplications(nil)
	if err != nil {
		return nil, err
	}
	if (!helper.containsApp(apps, config.ServiceID)) {
		return helper.createService(config, instances)
	} else {
		return helper.scaleService(config.ServiceID, instances)
	}
}

func (helper *Marathon) scaleService(id string, instances int) (*ServiceInformation, error){
	deploymentId, err := helper.client.ScaleApplicationInstances(id, instances, true)
	if err != nil {
    		util.Log.Errorf("Failed to Scale the application: %s, error: %s", id, err)
		return nil, err
	} else {
		app, err := helper.client.Application(id)
		if err != nil {
                	return nil, err
	        } else {
			serviceInformation := helper.getServiceInformationFromApp(app) 
			serviceInformation.Instances = helper.getInstancesByVersion(app.Tasks, app.Container.Docker.PortMappings, deploymentId.Version)
			return serviceInformation, nil
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

func (scheduler *Marathon) UndeployInstance(id string) error {
	return errors.New("Not implemented yet")
}

func (helper *Marathon) getInstancesFromTasks(tasks []*marathon.Task, dockerPortMappings []*marathon.PortMapping) []*Instance {
        instances := make([]*Instance, len(tasks))
	if tasks == nil || len(tasks) == 0 {
		return nil
	}
        for i, task := range tasks {
                instance := Instance{}
                instance.ID = task.ID
		instance.Host = task.Host
		//instance.ContainerName
                instance.Ports = helper.buildInstancePorts(dockerPortMappings, task.Ports)
                instances[i] = &instance
        }
        return instances
}

func (helper *Marathon) buildInstancePorts(dockerPortMappings []*marathon.PortMapping, taskPorts []int) map[string]InstancePort{
	if dockerPortMappings == nil || len(dockerPortMappings) == 0 || taskPorts == nil || len(taskPorts) == 0 {
		return nil
	}
	fmt.Println("task ports len", len(taskPorts))
	fmt.Println("docker port len", dockerPortMappings)
	
	ports := make(map[string]InstancePort, len(taskPorts))
	
	for i, port := range taskPorts {
		instancePort := InstancePort{}
		instancePort.Type = NewInstancePortType(dockerPortMappings[i].Protocol)
		instancePort.Internal = int64(port)
		instancePort.Publics = []int64{int64(dockerPortMappings[i].ContainerPort)}
		//instancePort.Advertise = 
		ports[string(instancePort.Publics[0]) + "/" + string(instancePort.Type)] = instancePort
	}
	return ports
}

func (helper *Marathon) getInstancesByVersion(tasks []*marathon.Task, dockerPortMappings []*marathon.PortMapping, version string) []*Instance {
	for i, task := range tasks {
		if task.Version != version {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	return helper.getInstancesFromTasks(tasks, dockerPortMappings)
}

func (m *Marathon) containsApp(apps []string, search string) bool{
	for _, a := range apps {
		if a == search {
			return true
		}
	}
	return false
}
func (s *Marathon) ID() string {
	return frameworkID
}
