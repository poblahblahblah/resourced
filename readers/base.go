package readers

import (
	"errors"
)

func NewGoStruct(name string) (IReaderWriter, error) {
	var structInstance IReaderWriter

	if name == "CpuInfo" {
		structInstance = NewCpuInfo()
	}
	if name == "Df" {
		structInstance = NewDf()
	}
	if name == "Du" {
		structInstance = NewDu()
	}
	if name == "DockerContainersMemory" {
		structInstance = NewDockerContainersMemory()
	}
	if name == "DockerContainersCpu" {
		structInstance = NewDockerContainersCpu()
	}
	if name == "LoadAvg" {
		structInstance = NewLoadAvg()
	}
	if name == "Memory" {
		structInstance = NewMemory()
	}
	if name == "Ps" {
		structInstance = NewPs()
	}
	if name == "Meminfo" {
		structInstance = NewMeminfo()
	}
	if name == "NetworkInterfaces" {
		structInstance = NewNetworkInterfaces()
	}
	if name == "Uptime" {
		structInstance = NewUptime()
	}

	if structInstance == nil {
		return nil, errors.New("GoStruct is undefined.")
	}

	return structInstance, nil
}

type IReaderWriter interface {
	Run() error
	ToJson() ([]byte, error)
}
