package cluster

import (
	"github.com/jglobant/yale/framework"
	"github.com/jglobant/yale/model"
	"github.com/jglobant/yale/monitor"
	"github.com/jglobant/yale/util"
)

type StackManager struct {
	stacks            map[string]*Stack
	stackNotification chan StackStatus
}

func NewStackManager() *StackManager {
	sm := new(StackManager)
	sm.stacks = make(map[string]*Stack)
	sm.stackNotification = make(chan StackStatus, 100)

	return sm
}

func (sm *StackManager) createId() string {
	i := 0
	for {
		key := util.Letter(i)
		exist := false

		for k := range sm.stacks {
			if k == key {
				exist = true
			}
		}

		if !exist {
			return key
		}
		i++
	}
}

func (sm *StackManager) AppendStack(fh framework.Framework) {
	key := sm.createId()
	util.Log.Infof("API configurada y mapeada a la llave %s", key)
	sm.stacks[key] = NewStack(key, sm.stackNotification, fh)
}

func (sm *StackManager) Deploy(serviceConfig model.ServiceConfig, smokeConfig monitor.MonitorConfig, warmConfig monitor.MonitorConfig, instances int, tolerance float64) bool {
	util.Log.Infoln("enter deploy stack manager %d", len(sm.stacks))
	
	for stackKey, _ := range sm.stacks {
		sm.stacks[stackKey].DeployCheckAndNotify(serviceConfig, smokeConfig, warmConfig, instances, tolerance)
	}
/*
	for i := 0; i < len(sm.stacks); i++ {
		stackStatus := <-sm.stackNotification
		util.Log.Infoln("Se recibió notificación del Stack con estado", stackStatus)
		if stackStatus == STACK_FAILED {
			util.Log.Errorln("Fallo el stack, se procederá a realizar Rollback")
			sm.Rollback()
			return false
		}
	}*/
	util.Log.Infoln("Proceso de deploy OK")
	return true
}

func (sm *StackManager) DeployedContainers() []*model.Instance {
	var containers []*model.Instance

	for stackKey, _ := range sm.stacks {
		containers = append(containers, sm.stacks[stackKey].ServicesWithStep(model.STEP_WARM_READY)...)
	}

	return containers
}

func (sm *StackManager) Rollback() {
	util.Log.Infoln("Iniciando el Rollback")
	for stack, _ := range sm.stacks {
		sm.stacks[stack].Rollback()
	}
}
