package model

import (
	"errors"
	"fmt"
	"regexp"
	"github.com/jglobant/yale/util"
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

func (s *ServiceConfig) Version() (string, error) {
	rp := regexp.MustCompile("^([\\d\\.]+)-")
	result := rp.FindStringSubmatch(s.Tag)
	if result == nil {
		return "", errors.New("Formato de TAG invalido")
	}
	return result[1], nil
}

func (s *ServiceConfig) String() string {
	return fmt.Sprintf("ImageName: %s - Tag: %s - CpuShares: %d - Memory: %s - Publish: %#v - Envs: %s", s.ImageName, s.Tag, s.CpuShares, s.Memory, s.Publish, util.MaskEnv(s.Envs))
}
