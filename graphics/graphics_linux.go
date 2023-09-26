//go:build linux

package graphics

import (
	"os/exec"
	"regexp"
	"strconv"
	"unsafe"
)

// GetGraphicsCardInformation 获取显卡信息
func GetGraphicsCardInformation() (*GraphicsInformation, error) {
	output, err := exec.Command("sh", "-c", "lspci -vnn | grep VGA -A 12").Output()
	if err != nil {
		return nil, err
	}
	manufactureRegex := regexp.MustCompile(`.*: (.*?) Device \[(.*)] \(rev (.*?)\)`)
	nonPrefetchableMemoryRegex := regexp.MustCompile(`.*Memory at (.*)non-prefetchable.* \[size=(.+)]`)
	prefetchableMemoryRegex := regexp.MustCompile(`.*Memory at .* prefetchable.* \[size=(.+)]`)
	ioPortRegex := regexp.MustCompile(`.*I/O ports at (\d+) \[size=(\d+)]`)
	result := unsafe.String(unsafe.SliceData(output), len(output))
	info := &GraphicsInformation{}
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
