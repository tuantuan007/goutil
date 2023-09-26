package graphics

// GraphicsInformation 显卡信息
type GraphicsInformation struct {
	Manufacturer    string `json:"Manufacturer,omitempty"`
	Name            string
	DeviceID        string `json:"DeviceID,omitempty"`
	Version         string `json:"Version,omitempty"`
	NonPrefetchable string `json:"NonPrefetchable,omitempty"`
	Prefetchable    string `json:"Prefetchable,omitempty"` // 预取内存大小
	IOPort          int    `json:"IOPort,omitempty"`
	IOSize          int    `json:"IOSize,omitempty"`
}
