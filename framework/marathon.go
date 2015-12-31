package framework 

import (
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

func (helper *MarathonHelper) DeployService(config model.ServiceConfig) (error) {
	translator := MarathonAppTranslator{}
	application := translator.TranslateServiceConfig(config)
	_, err := helper.client.CreateApplication(application)
	return err
}

func (helper *MarathonHelper) ScaleService(id string, instances int) (error){
	_, err := helper.client.ScaleApplicationInstances(id, instances, true)
	if err != nil {
    		util.Log.Errorf("Failed to Scale the application: %s, error: %s", id, err)
	}
	return err
}

func (helper *MarathonHelper) DeleteService(id string) (error) {
	_, err := helper.client.DeleteApplication(id)
	if err != nil {
		util.Log.Errorf("Failed to Delete the application: %s, error: %s", id, err)
	}
	return err
}
