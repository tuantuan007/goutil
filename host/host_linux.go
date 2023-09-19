//go:build linux

package host

import (
	"bytes"
	"os/exec"
)

type Information struct {
	StaticHostname  string `yaml:"Static hostname"`
	IconName        string `yaml:"Icon name"`
	Chassis         string `yaml:"Chassis"`
	MachineID       string `yaml:"Machine ID"`
	BootID          string `yaml:"Boot ID"`
	Virtualization  string `yaml:"Virtualization"`
	OperatingSystem string `yaml:"Operating System"`
	CPEOSName       string `yaml:"CPE OS Name"`
	Kernel          string `yaml:"Kernel"`
	Architecture    string `yaml:"Architecture"`
}

func GetHostInformation() (*Information, error) {
	info := &Information{}
	output, err := exec.Command("hostnamectl", "status").Output()
	if err != nil {
		return nil, err
	}
	buffer := bytes.Buffer{}
	buffer.Grow(len(output))
	lines := bytes.Split(output, []byte("\n"))
	for _, line := range lines {
		buffer.Write(bytes.TrimSpace(line))
		buffer.Write([]byte("\n"))
	}
	if err := yaml.Unmarshal(buffer.Bytes(), info); err != nil {
		return nil, err
	}
	return info, nil
}
