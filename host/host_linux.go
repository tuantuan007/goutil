//go:build linux

package host

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"os/exec"
)

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
