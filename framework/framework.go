package framework

import "github.com/jglobant/yale/model"

type FrameworkHelper interface {
	ListServices() []string
	DeployService(config model.ServiceConfig) (error)
	ScaleService(id string, instances int) (error)
	DeleteService(id string) (error)
}

func NewFrameworkHelper(endpointUrl string) (FrameworkHelper, error) {
	return NewMarathonHelper(endpointUrl)
}

func NewFrameworkTlsHelper(endpointUrl, cert,  key string) (FrameworkHelper, error) {
        return NewMarathonTlsHelper(endpointUrl, cert,  key)
}

func NewFrameworkTlsVerifyHelper(endpointUrl, cert,  key, ca string) (FrameworkHelper, error) {
        return NewMarathonTlsVerifyHelper(endpointUrl, cert,  key, ca)
}
