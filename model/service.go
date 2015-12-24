package model

import (
	"fmt"
	"regexp"
	"github.com/yojomapa/yale/util"
)

type ServiceConfig struct {
	ServiceId string
	CpuShares int
	Envs      []string
	ImageName string
	Memory    int64
	Publish   []string
	Tag       string
	Instances int
}

func (s *ServiceConfig) Version() string {
	rp := regexp.MustCompile("^([\\d\\.]+)-")
	result := rp.FindStringSubmatch(s.Tag)
	if result == nil {
		util.Log.Fatalln("Formato de TAG invalido")
	}
	return result[1]
}

func (s *ServiceConfig) String() string {
	return fmt.Sprintf("ImageName: %s - Tag: %s - CpuShares: %d - Memory: %s - Publish: %#v - Envs: %s", s.ImageName, s.Tag, s.CpuShares, s.Memory, s.Publish, util.MaskEnv(s.Envs))
}