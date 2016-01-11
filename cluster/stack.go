package cluster

import (
	"github.com/Pallinder/go-randomdata"
	log "github.com/Sirupsen/logrus"
	"github.com/jglobant/yale/framework"
	"github.com/jglobant/yale/model"
	"github.com/jglobant/yale/monitor"
	"github.com/jglobant/yale/util"
	"fmt"
)

type StackStatus int

const (
	STACK_READY StackStatus = 1 + iota
	STACK_FAILED
)

var stackStatus = [...]string{
	"STACK_READY",
	"STACK_FAILED",
}

func (s StackStatus) String() string {
	return stackStatus[s-1]
}

type Stack struct {
	id                    string
	frameworkApiHelper    framework.Framework
	instances             []*model.Instance // refactorizar a interfaz service
	serviceIdNotification chan string
	stackNofitication     chan<- StackStatus
	smokeTestMonitor      monitor.Monitor
	warmUpMonitor         monitor.Monitor
	log                   *log.Entry
}

func NewStack(stackKey string, stackNofitication chan<- StackStatus, fh framework.Framework) *Stack {
	s := new(Stack)
	s.id = stackKey
	s.stackNofitication = stackNofitication
	s.frameworkApiHelper = fh
	s.serviceIdNotification = make(chan string, 1000)

	s.log = util.Log.WithFields(log.Fields{
		"stack": stackKey,
	})

	return s
}

func (s *Stack) createId() string {
	for {
		key := s.id + "_" + randomdata.Adjective()
		exist := false

		for _, srv := range s.instances {
			if srv.Id == key {
				exist = true
			}
		}

		if !exist {
			return key
		}
	}
}

func (s *Stack) createMonitor(config monitor.MonitorConfig) monitor.Monitor {
	var mon monitor.Monitor

	s.log.Infof("Creando monitor con mode [%s] y request [%s]", config.Type, config.Request)
	if config.Type == monitor.TCP {
		mon = new(monitor.TcpMonitor)
	} else {
		mon = new(monitor.HttpMonitor)
	}

	mon.SetRetries(config.Retries)
	mon.SetRequest(config.Request)
	mon.SetExpected(config.Expected)

	return mon
}

func (s *Stack) DeployCheckAndNotify(serviceConfig model.ServiceConfig, smokeConfig monitor.MonitorConfig, warmConfig monitor.MonitorConfig, instances int, tolerance float64) {
	_, err := s.frameworkApiHelper.DeployService(serviceConfig)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Stack) setStatus(status StackStatus) {
	s.stackNofitication <- status
}

func (s *Stack) undeployInstance(instance *model.Instance) {
	s.frameworkApiHelper.UndeployInstance(instance)
}

func (s *Stack) Rollback() {
	s.log.Infof("Comenzando Rollback en el Stack")
	for _, srv := range s.instances {
		if !srv.IsLoaded() {
			s.undeployInstance(srv)
		}
	}
}

func (s *Stack) UndeployInstances(total int) {
	undeployed := 0
	for _, srv := range s.instances {
		if undeployed == total {
			return
		}
		s.undeployInstance(srv)
		undeployed++
	}
}

func (s *Stack) ServicesWithStep(step model.Step) []*model.Instance {
	var instances []*model.Instance
	for k, v := range s.instances {
		if v.GetStep() == step {
			instances = append(instances, s.instances[k])
		}
	}
	return instances
}

func (s *Stack) ServicesWithState(state model.State) []*model.Instance {
	var instances []*model.Instance
	for k, v := range s.instances {
		if v.CheckState(state) {
			instances = append(instances, s.instances[k])
		}
	}
	return instances
}

func (s *Stack) countServicesWithStep(step model.Step) int {
	return len(s.ServicesWithStep(step))
}

func (s *Stack) countServicesWithState(state model.State) int {
	return len(s.ServicesWithState(state))
}
