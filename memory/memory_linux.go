//go:build linux

package memory

import (
	"os/exec"
	"regexp"
	"strings"
)

// GetPhysicalMemoryArray 获取linux物理内存阵列
func GetPhysicalMemoryArray() (*PhysicalMemoryArray, error) {
	cmd := exec.Command("sh", "-c", "dmidecode -t 16 | grep -e 'Location' -e 'Use' -e 'Error Correction Type' -e 'Maximum Capacity' -e 'Error Information Handle' -e 'Number Of Devices'")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(output), "\n")
	return &PhysicalMemoryArray{
		Location:               strings.Split(lines[0], ":")[1],
		Use:                    strings.Split(lines[1], ":")[1],
		ErrorCorrectionType:    strings.Split(lines[2], ":")[1],
		MaximumCapacity:        strings.Split(lines[3], ":")[1],
		ErrorInformationHandle: strings.Split(lines[4], ":")[1],
		NumberOfDevices:        strings.Split(lines[5], ":")[1],
	}, nil
}

func GetMemoryDevices() ([]*MemoryDeviceInformation, error) {
	output, err := exec.Command("sh", "-c", "dmidecode -t 17").Output()
	if err != nil {
		return nil, err
	}

	result := string(output)
	// 使用正则表达式提取字段值
	re := regexp.MustCompile(`(?s)Memory\sDevice(.*?)Memory\sDevice`)
	matches := re.FindStringSubmatch(result)
	devices := make([]*MemoryDeviceInformation, 0, len(matches))
	for _, match := range matches {
		device := &MemoryDeviceInformation{}
		lines := strings.Split(match, "\n")
		for _, line := range lines {
			keyAndValue := strings.SplitN(line, ":", 2)
			if len(keyAndValue) != 2 {
				continue
			}
			key := strings.TrimSpace(keyAndValue[0])
			value := strings.TrimSpace(keyAndValue[1])
			// 使用 switch 语句提取字段值
			switch key {
			case "Total Width":
				device.TotalWidth = value
			case "Data Width":
				device.DataWidth = value
			case "Size":
				device.Size = value
			case "Form Factor":
				device.FormFactor = value
			case "Locator":
				device.Locator = value
			case "Bank Locator":
				device.BankLocator = value
			case "Type":
				device.Type = value
			case "Type Detail":
				device.TypeDetail = value
			case "Speed":
				device.Speed = value
			case "Manufacturer":
				device.Manufacturer = value
			case "Serial Number":
				device.SerialNumber = value
			case "Asset Tag":
				device.AssetTag = value
			case "Part Number":
				device.PartNumber = value
			case "Rank":
				device.Rank = value
			case "Configured Clock Speed":
				device.ConfiguredClockSpeed = value
			case "Configured Voltage":
				device.ConfiguredVoltage = value
			case "Minimum Voltage":
				device.MinimumVoltage = value
			case "Maximum Voltage":
				device.MaximumVoltage = value
			}
		}
		devices = append(devices, device)
	}
	return devices, nil
}
