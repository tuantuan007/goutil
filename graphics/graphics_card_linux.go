//go:build linux

package graphics

import (
	"os/exec"
	"regexp"
	"strconv"
	"unsafe"
)

// GraphicsInformation 显卡信息
type Information struct {
	Manufacturer    string `json:"Manufacturer,omitempty"`
	DeviceID        string `json:"DeviceID,omitempty"`
	Version         string `json:"Version,omitempty"`
	NonPrefetchable string `json:"NonPrefetchable,omitempty"`
	Prefetchable    string `json:"Prefetchable,omitempty"` // 预取内存大小
	IOPort          int    `json:"IOPort,omitempty"`
	IOSize          int    `json:"IOSize,omitempty"`
}

// GetGraphicsCardInformation  获取显卡信息
func GetGraphicsCardInformation() (*Information, error) {
	output, err := exec.Command("sh", "-c", "lspci -vnn | grep VGA -A 12").Output()
	if err != nil {
		return nil, err
	}
	manufactureRegex := regexp.MustCompile(`.*: (.*?) Device \[(.*)] \(rev (.*?)\)`)
	nonPrefetchableMemoryRegex := regexp.MustCompile(`.*Memory at (.*)non-prefetchable.* \[size=(.+)]`)
	prefetchableMemoryRegex := regexp.MustCompile(`.*Memory at .* prefetchable.* \[size=(.+)]`)
	ioPortRegex := regexp.MustCompile(`.*I/O ports at (\d+) \[size=(\d+)]`)
	result := unsafe.String(unsafe.SliceData(output), len(output))
	info := &Information{}
	matches := manufactureRegex.FindStringSubmatch(result)
	info.Manufacturer = matches[1]
	info.DeviceID = matches[2]
	info.Version = matches[3]
	matches = nonPrefetchableMemoryRegex.FindStringSubmatch(result)
	info.NonPrefetchable = matches[2]
	matches = prefetchableMemoryRegex.FindStringSubmatch(result)
	info.Prefetchable = matches[1]
	matches = ioPortRegex.FindStringSubmatch(result)
	info.IOPort, _ = strconv.Atoi(matches[1])
	info.IOSize, _ = strconv.Atoi(matches[2])
	return info, nil
}
