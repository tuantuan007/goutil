//go:build linux

package processor

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strings"
	"unsafe"
)

type Information struct {
	// SocketDesignation 处理器的插槽名称
	SocketDesignation string `yaml:"Socket Designation" json:"SocketDesignation,omitempty"`
	// Type 处理器的类型
	Type string `yaml:"Type" json:"Type,omitempty"`
	// Family 处理器的系列
	Family string `yaml:"Family" json:"Family,omitempty"`
	// Manufacturer 处理器的制造商
	Manufacturer string `yaml:"Manufacturer" json:"Manufacturer,omitempty"`
	// ID 处理器的唯一标识符
	ID string `yaml:"ID" json:"ID,omitempty"`
	// Signature  处理器的签名，包括类型、系列、型号和步进信息
	Signature string `yaml:"Signature" json:"Signature,omitempty"`
	// Flags 处理器的特性标志，如支持的指令集及其他功能
	Flags []string `yaml:"Flags" json:"Flags,omitempty"`
	// Version 处理器的版本信息
	Version string `yaml:"Version" json:"Version,omitempty"`
	// Voltage 处理器的电压
	Voltage string `yaml:"Voltage" json:"Voltage,omitempty"`
	// ExternalClock 外部时钟频率
	ExternalClock string `yaml:"External Clock" json:"ExternalClock,omitempty"`
	// MaxSpeed 处理器的最大速度
	MaxSpeed string `yaml:"Max Speed" json:"MaxSpeed,omitempty"`
	// CurrentSpeed 当前处理器的运行速度
	CurrentSpeed string `yaml:"Current Speed" json:"CurrentSpeed,omitempty"`
	// Status 处理器的状态，如是否启用
	Status string `yaml:"Status" json:"Status,omitempty"`
	// Upgrade 处理器的升级选项
	Upgrade string `yaml:"Upgrade" json:"Upgrade,omitempty"`
	// L1CacheHandle L1缓存的句柄
	L1CacheHandle string `yaml:"L1 Cache Handle" json:"L1CacheHandle,omitempty"`
	// L2CacheHandle L2缓存的句柄
	L2CacheHandle string `yaml:"L2 Cache Handle" json:"L2CacheHandle,omitempty"`
	// L3CacheHandle L3缓存的句柄
	L3CacheHandle string `yaml:"L3 Cache Handle" json:"L3CacheHandle,omitempty"`
	// SerialNumber 处理器的序列号
	SerialNumber string `yaml:"Serial Number" json:"SerialNumber,omitempty"`
	// AssetTag 处理器的资产标记
	AssetTag string `yaml:"Asset Tag" json:"AssetTag,omitempty"`
	// PartNumber 处理器的零件号
	PartNumber string `yaml:"Part Number" json:"PartNumber,omitempty"`
	// CoreCount  处理器的核心数量
	CoreCount int `yaml:"Core Count" json:"CoreCount,omitempty"`
	// CoreEnabled 处理器的启用核心数量
	CoreEnabled int `yaml:"Core Enabled" json:"CoreEnabled,omitempty"`
	// ThreadCount 处理器的线程数量
	ThreadCount int `yaml:"Thread Count" json:"ThreadCount,omitempty"`
	// Characteristics 处理器的特性标志，如64位支持、多核心、硬件线程、执行保护、增强虚拟化、功耗/性能控制等
	Characteristics []string `yaml:"Characteristics" json:"Characteristics,omitempty"`
}

// GetProcessorInformation 获取处理器信息
func GetProcessorInformation() (*Information, error) {
	p := &Information{}
	// 调用dmidecode命令，获取处理器信息
	output, err := exec.Command("sh", "-c", "dmidecode -t 4").Output()
	if err != nil {
		return nil, err
	}
	// 获取处理器信息
	result := unsafe.String(unsafe.SliceData(output), len(output))
	// 匹配Processor Information
	pattern := `Processor Information\n([\s\S]+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(result)
	if len(matches) < 2 {
		return nil, errors.New("no processor information found")
	}
	// 创建一个缓冲区，用于存放处理器信息
	builder := bytes.Buffer{}
	builder.Grow(len(matches[1]))
	// 将处理器信息拆分成行
	lines := strings.Split(matches[1], "\n")
	// 遍历行，将每一行添加到缓冲区中
	for _, line := range lines {
		if strings.HasPrefix(line, "\t\t") {
			line = strings.Replace(line, "\t\t", " - ", 1)
		}
		if strings.HasPrefix(line, "\t") {
			line = strings.Replace(line, "\t", "", 1)
		}
		if strings.Contains(line, "Characteristics") && strings.Contains(line, "None") {
			line = strings.Replace(line, "None", "", 1)
		}
		builder.WriteString(line + "\n")
	}
	// 将缓冲区中的内容转换为YAML格式
	if err := yaml.Unmarshal(builder.Bytes(), p); err != nil {
		return nil, err
	}
	return p, nil
}
