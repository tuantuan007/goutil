package processor

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
