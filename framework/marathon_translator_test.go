package framework

import (
        "testing"
        "github.com/jglobant/yale/model"
        "github.com/stretchr/testify/assert"
)

func TestTranslateApp(t *testing.T) {
	cfg := model.ServiceConfig{}
	app := translateServiceConfig(cfg)
        assert.True(t, app.Container.Docker.PortMappings == nil, "App should not have any ports")

	cfg = model.ServiceConfig {
		Publish: []string{},
	}
	app = translateServiceConfig(cfg)
	assert.True(t, app.Container.Docker.PortMappings == nil, "App should not have any ports")

        cfg = model.ServiceConfig {
                Publish: []string{"80","443"},
        }
        app = translateServiceConfig(cfg)
        assert.Equal(t, 2, len(app.Container.Docker.PortMappings), "App should not have 2 ports")
	assert.Equal(t, 80, app.Container.Docker.PortMappings[0].ContainerPort, "Port should be 80")
	assert.Equal(t, 443, app.Container.Docker.PortMappings[1].ContainerPort, "Port should be 443")
}
