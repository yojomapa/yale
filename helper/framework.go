package helper

import "github.com/yojomapa/yale/model"

// The Framework Helper interface that all future framework implementations should fulfill
type FrameworkHelper interface {
	ListServices() []string
	DeployService(config model.ServiceConfig)
	ScaleService(id string, instances int)
	DeleteService(id string)
	
}

// This is the 'Factory of Helper' kinda
func NewFrameworkHelper(endpointUrl string) (FrameworkHelper, error) {
	
	//just for now we only have MarathonHelper
	return NewMarathonHelper(endpointUrl)
}