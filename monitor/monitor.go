package monitor

type MonitorType int

const (
	HTTP MonitorType = 1 + iota
	TCP
)

var monitorType = [...]string{
	"HTTP",
	"TCP",
}

func (s MonitorType) String() string {
	return monitorType[s-1]
}

func GetMonitor(t string) MonitorType {
	if t == TCP.String() {
		return TCP
	}

	return HTTP
}

type MonitorConfig struct {
	Type    MonitorType
	Retries int
	Ping    string
	Pong    string
}

type Monitor interface {
	Check(addr string) bool
	SetEndpoint(ep string)
	SetExpect(ex string)
	SetRetries(retries int)
	Configured() bool
}
