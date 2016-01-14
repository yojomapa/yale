package cluster

import (
	"github.com/Pallinder/go-randomdata"
	log "github.com/Sirupsen/logrus"
	"github.com/yojomapa/yale/framework"
	"github.com/yojomapa/yale/util"
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
	instances             []*framework.ServiceInformation
	serviceIdNotification chan string
	stackNofitication     chan<- StackStatus
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
			if srv.ID == key {
				exist = true
			}
		}

		if !exist {
			return key
		}
	}
}


func (s *Stack) DeployCheckAndNotify(serviceConfig framework.ServiceConfig, instances int, tolerance float64) {
	_, err := s.frameworkApiHelper.DeployService(serviceConfig, instances)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Stack) setStatus(status StackStatus) {
	s.stackNofitication <- status
}

func (s *Stack) undeployInstance(instance string) {
	s.frameworkApiHelper.UndeployInstance(instance)
}

func (s *Stack) Rollback() {
	s.log.Infof("Comenzando Rollback en el Stack")
	for _, srv := range s.instances {
		//if !srv.IsLoaded() {
		s.undeployInstance(srv.ID)
		//}
	}
}

func (s *Stack) UndeployInstances(total int) {
	undeployed := 0
	for _, srv := range s.instances {
		if undeployed == total {
			return
		}
		s.undeployInstance(srv.ID)
		undeployed++
	}
}
