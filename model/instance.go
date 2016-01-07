package model

type State int

const (
	RUNNING State = 1 + iota
	UNDEPLOYED
)

type Step int

const (
	STEP_CREATED Step = 1 + iota
	STEP_SMOKE_READY
	STEP_WARM_READY
	STEP_FAILED
)

type Instance struct {
	Id	string
	Type	string
	Name	string
	Ports	[]int
	Node	string
	state	State
	Created	string
	Image 	string
	loaded	bool
	step	Step
}

func (i *Instance) SetLoaded(newLoaded bool) {
	i.loaded = newLoaded
}

func (i *Instance) IsLoaded() bool {
	return i.loaded
}

func (i *Instance) SetStep(status Step) {
	i.step = status
}

func (i *Instance) GetStep() Step {
	return i.step
}

func (i *Instance) CheckState(state State) bool {
	if i.Id == "" {
		return false
	}

	if state == RUNNING && i.state == RUNNING {
		return true
	}

	if state == UNDEPLOYED && i.state == UNDEPLOYED {
		return true
	}

	return false
}

func (i *Instance) setState(state State) {
	i.state = state
}

func (i *Instance) RegistratorId() string {
	return i.Node + ":" + i.Name + ":8080"
}
