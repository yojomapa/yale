package framework

import (
	"errors"
	"github.com/jglobant/yale/model"
	"strings"
)

type Framework interface {
	FindServiceInformation(serviceName string) ([]*model.Instance, error)
	DeployService(config model.ServiceConfig) ([]*model.Instance, error)
	DeleteService(id string) (error)
	UndeployInstance(instance *model.Instance) (error)
}

type FrameworkType int

const ( 
	MARATHON FrameworkType = 1 << iota
	CHRONOS
	SWARM
	NOT_VALID
)

type FrameworkConfig struct {
	Type		FrameworkType
	EndpointUrl	string
	Cert		string
	Key		string
	Ca		string	
}

func NewFramework(cfg FrameworkConfig) (Framework, error) {
	switch cfg.Type {
		case MARATHON:
			return NewMarathon(cfg.EndpointUrl)
	}
	return nil, errors.New("Not implemented yet")
}

func NewFrameworkTlsHelper(cfg FrameworkConfig) (Framework, error) {
        switch cfg.Type {
                case MARATHON:
                        return NewMarathonTls(cfg.EndpointUrl, cfg.Cert,  cfg.Key)
        }
        return nil, errors.New("Not implemented yet")
}

func NewFrameworkTlsVerifyHelper(cfg FrameworkConfig) (Framework, error) {
        switch cfg.Type {
                case MARATHON:
		return NewMarathonTlsVerify(cfg.EndpointUrl, cfg.Cert,  cfg.Key, cfg.Ca)                    
        }
        return nil, errors.New("Not implemented yet")
}

func GetFrameworkType(aType string) FrameworkType {
	switch strings.ToUpper(aType) {
		case "MARATHON":
			return MARATHON		
	}
	return NOT_VALID
}
