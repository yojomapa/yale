package framework

import (
        "github.com/gambol99/go-marathon"
	"github.com/jglobant/yale/model"
        "strconv"
	"fmt"
	"encoding/json"
)

func translateServiceConfig(config model.ServiceConfig) *marathon.Application {
        application := marathon.NewDockerApplication()
        imageWithTag := config.ImageName + ":" + config.Tag
        labels := map[string]string{
                "image_name": config.ImageName,
                "image_tag":  config.Tag,
        }

        application.Name(config.ServiceId)
        application.CPU(0.25) // how to map this ?
        //application.Memory(float64(config.Memory))
        application.Memory(64)
        application.Count(config.Instances)
        //application.Env = util.StringSlice2Map(config.Envs)
        application.Labels = labels
	//application.RequirePorts = true

        // add the docker container
        application.Container.Docker.Container(imageWithTag)
        application.Container.Docker.Expose(80, 443)
        //application.Container.Docker.PortMappings = createPorMappings(config.Publish)
	s := []string{"80", "443"}
        application.Container.Docker.PortMappings = createPorMappings(s)
        //application.CheckHTTP("/health", 10, 5)

	b, _ := json.Marshal(application)
	fmt.Println(string(b))
        return application
}

func createPorMappings(ports []string) []*marathon.PortMapping {
        if ports == nil || len(ports) == 0 {
                return nil
        }

        portMappings := make([]*marathon.PortMapping, len(ports))
        for i, val := range ports {
                iPort, _ := strconv.Atoi(val)
                portMappings[i] = createPortMapping(iPort, "tcp")
        }

        return portMappings
}

func createPortMapping(containerPort int, protocol string) *marathon.PortMapping {
        return &marathon.PortMapping{
                ContainerPort: containerPort,
                HostPort:      0,
                ServicePort:   0,
                Protocol:      protocol,
        }
}

