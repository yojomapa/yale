package framework

import "github.com/jglobant/yale/model"

// The Framework Helper interface that all future framework implementations should fulfill
type FrameworkHelper interface {
	ListServices() []string
	DeployService(config model.ServiceConfig) (error)
	ScaleService(id string, instances int) (error)
	DeleteService(id string) (error)
}

// This is the 'Factory of Helper' kinda
func NewFrameworkHelper(endpointUrl string) (FrameworkHelper, error) {
	
	//just for now we only have MarathonHelper
	return NewMarathonHelper(endpointUrl)
}
