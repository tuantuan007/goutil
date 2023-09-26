package memory

// PhysicalMemoryArray 物理内存阵列
type PhysicalMemoryArray struct {
	Location               string
	Use                    string
	ErrorCorrectionType    string
	MaximumCapacity        string
	ErrorInformationHandle string
	NumberOfDevices        string
}

// MemoryDeviceInformation 内存设备
type MemoryDeviceInformation struct {
	TotalWidth           string
	DataWidth            string
	Size                 string
	FormFactor           string
	Locator              string
	BankLocator          string
	Type                 string
	TypeDetail           string
	Speed                string
	Manufacturer         string
	SerialNumber         string
	AssetTag             string
	PartNumber           string
	Rank                 string
	ConfiguredClockSpeed string
	MinimumVoltage       string
	MaximumVoltage       string
	ConfiguredVoltage    string
}
