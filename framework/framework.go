package framework

import (
	"errors"
	"strings"
)

// Framework es una interfaz que debe implementar para la comunicacion con los Schedulers de Docker
// Para un ejemplo ir a swarm.Framework
type Framework interface {
	ID() string
	FindServiceInformation(ServiceInformationCriteria) ([]*ServiceInformation, error)
	UndeployInstance(string) error
	DeployService(ServiceConfig, int) (*ServiceInformation, error)
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
