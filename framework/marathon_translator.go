package framework

import (
        "github.com/gambol99/go-marathon"
	"github.com/jglobant/yale/model"
	"github.com/jglobant/yale/util"
        "strconv"
)

type MarathonAppTranslator struct {
}


func (translator *MarathonAppTranslator) TranslateServiceConfig(config model.ServiceConfig) *marathon.Application {
        application := marathon.NewDockerApplication()
        imageWithTag := config.ImageName + ":" + config.Tag
        labels := map[string]string{
                "image_name": config.ImageName,
                "image_tag":  config.Tag,
        }

        application.ID = config.ServiceId
        application.Name(imageWithTag)
        application.CPU(0.1) // how to map this ?
        application.Memory(float64(config.Memory))
        application.Count(config.Instances)
        //application.Arg("/usr/sbin/apache2ctl", "-D", "FOREGROUND")
        application.Env = util.StringSlice2Map(config.Envs)
        application.Labels = labels
        // add the docker container
        application.Container.Docker.Container(imageWithTag)
        //application.Container.Docker.Expose(80, 443)
        application.Container.Docker.PortMappings = createPorMappings(config.Publish)
        //application.CheckHTTP("/health", 10, 5)
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

