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

// GetProcessorInformation 获取处理器信息
func GetProcessorInformation() (*ProcessorInformation, error) {
	p := &ProcessorInformation{}
	// 调用dmidecode命令，获取处理器信息
	output, err := exec.Command("sh", "-c", "dmidecode -t 4").Output()
	if err != nil {
		return nil, err
	}
	// 获取处理器信息
	result := unsafe.String(unsafe.SliceData(output), len(output))
	// 匹配Processor ProcessorInformation
	pattern := `Processor ProcessorInformation\n([\s\S]+)`
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
