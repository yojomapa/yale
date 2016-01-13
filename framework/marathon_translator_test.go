package framework

import (
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestTranslateApp(t *testing.T) {
	cfg := ServiceConfig{}
	app := translateServiceConfig(cfg, 1)
        assert.True(t, app.Container.Docker.PortMappings == nil, "App should not have any ports")

	cfg = ServiceConfig {
		Publish: []string{},
	}
	app = translateServiceConfig(cfg, 1)
	assert.True(t, app.Container.Docker.PortMappings == nil, "App should not have any ports")

        cfg = ServiceConfig {
                Publish: []string{"80","443"},
        }
        app = translateServiceConfig(cfg, 1)
        assert.Equal(t, 2, len(app.Container.Docker.PortMappings), "App should not have 2 ports")
	assert.Equal(t, 80, app.Container.Docker.PortMappings[0].ContainerPort, "Port should be 80")
	assert.Equal(t, 443, app.Container.Docker.PortMappings[1].ContainerPort, "Port should be 443")
}
