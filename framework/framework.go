package framework

import (
	"errors"
	"github.com/jglobant/yale/model"
)

type FrameworkHelper interface {
	ListServices(serviceName string) []model.Container
	DeployService(config model.ServiceConfig) (error)
	ScaleService(id string, instances int) (error)
	DeleteService(id string) (error)
}

type FrameworkType int

const ( 
	MARATHON FrameworkType = 1 << iota
	CHRONOS
	SWARM
)

type FrameworkConfig struct {
	Type		FrameworkType
	EndpointUrl	string
	Cert		string
	Key		string
	Ca		string	
}

func NewFrameworkHelper(cfg FrameworkConfig) (FrameworkHelper, error) {
	switch cfg.Type {
		case MARATHON:
			return NewMarathonHelper(cfg.EndpointUrl)
	}
	return nil, errors.New("Not implemented yet")
}

func NewFrameworkTlsHelper(cfg FrameworkConfig) (FrameworkHelper, error) {
        switch cfg.Type {
                case MARATHON:
                        return NewMarathonTlsHelper(cfg.EndpointUrl, cfg.Cert,  cfg.Key)
        }
        return nil, errors.New("Not implemented yet")
}

func NewFrameworkTlsVerifyHelper(cfg FrameworkConfig) (FrameworkHelper, error) {
        switch cfg.Type {
                case MARATHON:
		return NewMarathonTlsVerifyHelper(cfg.EndpointUrl, cfg.Cert,  cfg.Key, cfg.Ca)                    
        }
        return nil, errors.New("Not implemented yet")
}
