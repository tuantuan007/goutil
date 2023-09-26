//go:build windows

package memory

import (
	"goutil/common/windows"
	"strconv"
	"time"
)

type Win32_PhysicalMemoryArray struct {
	Caption               string
	CreationClassName     string
	Depth                 float32
	Description           string
	Height                float32
	HotSwappable          bool
	InstallDate           time.Time
	Location              uint16
	Manufacturer          string
	MaxCapacity           uint32
	MaxCapacityEx         uint64
	MemoryDevices         uint16
	MemoryErrorCorrection uint16
	Model                 string
	Name                  string
	OtherIdentifyingInfo  string
	PartNumber            string
	PoweredOn             bool
	Removable             bool
	Replaceable           bool
	SerialNumber          string
	Sku                   string
	Status                string
	Tag                   string
	Use                   uint16
	Version               string
	Weight                float32
	Width                 float32
}

var (
	PhysicalMemoryArrayLocationMap = map[uint16]string{
		0:  "Reserved",
		1:  "Other",
		2:  "Unknown",
		3:  "System board or motherboard",
		4:  "ISA add-on card",
		5:  "EISA add-on card",
		6:  "PCI add-on card",
		7:  "MCA add-on card",
		8:  "PCMCIA add-on card",
		9:  "Proprietary add-on card",
		10: "NuBus",
		11: "PC-98/C20 add-on card",
		12: "PC-98/C24 add-on card",
		13: "PC-98/E add-on card",
		14: "PC-98/Local bus add-on card",
	}
	PhysicalMemoryArrayMemoryTypesMap = map[uint16]string{
		0: "Reserved",
		1: "Other",
		2: "Unknown",
		3: "System memory",
		4: "Video memory",
		5: "Flash memory",
		6: "Non-volatile RAM",
		7: "Cache memory",
	}
	PhysicalMemoryArrayMemoryErrorCorrectionMap = map[uint16]string{
		0: "Reserved",
		1: "Other",
		2: "Unknown",
		3: "None",
		4: "Parity",
		5: "Single-bit ECC",
		6: "Multi-bit ECC",
		7: "CRC",
	}
)

func GetPhysicalMemoryArray() (*PhysicalMemoryArray, error) {
	dst, err := windows.Query[Win32_PhysicalMemoryArray]()
	if err != nil {
		return nil, err
	}
	result := dst[0]
	output := &PhysicalMemoryArray{
		Location:               PhysicalMemoryArrayLocationMap[result.Location],
		Use:                    PhysicalMemoryArrayMemoryTypesMap[result.Use],
		ErrorCorrectionType:    PhysicalMemoryArrayMemoryErrorCorrectionMap[result.MemoryErrorCorrection],
		MaximumCapacity:        strconv.Itoa(int(result.MaxCapacity >> 20)),
		ErrorInformationHandle: "",
		NumberOfDevices:        strconv.Itoa(int(result.MemoryDevices)),
	}
	return output, nil
}

// Win32_PhysicalMemory https://learn.microsoft.com/zh-cn/windows/win32/cimwin32prov/win32-physicalmemory
type Win32_PhysicalMemory struct {
	Attributes           uint32
	BankLabel            string
	Capacity             uint64
	Caption              string
	ConfiguredClockSpeed uint32
	ConfiguredVoltage    uint32
	CreationClassName    string
	DataWidth            uint16
	Description          string
	DeviceLocator        string
	FormFactor           uint16
	HotSwappable         bool
	InstallDate          time.Time
	InterleaveDataDepth  uint16
	InterleavePosition   uint32
	Manufacturer         string
	MaxVoltage           uint32
	MemoryType           uint16
	MinVoltage           uint32
	Model                string
	Name                 string
	OtherIdentifyingInfo string
	PartNumber           string
	PositionInRow        uint32
	PoweredOn            bool
	Removable            bool
	Replaceable          bool
	SerialNumber         string
	SKU                  string
	SMBIOSMemoryType     uint32
	Speed                uint32
	Status               string
	Tag                  string
	TotalWidth           uint16
	TypeDetail           uint16
	Version              string
}

var (
	TypeMap = map[uint16]string{
		0:  "Unknown",
		1:  "Other",
		2:  "DRAM",
		3:  "Synchronous DRAM",
		4:  "Cache DRAM",
		5:  "EDO",
		6:  "EDRAM",
		7:  "VRAM",
		8:  "SRAM",
		9:  "RAM",
		10: "ROM",
		11: "Flash",
		12: "EEPROM",
		13: "FEPROM",
		14: "EPROM",
		15: "CDRAM",
		16: "3DRAM",
		17: "SDRAM",
		18: "SGRAM",
		19: "RDRAM",
		20: "DDR",
		21: "DDR2",
		22: "DDR2 FB-DIMM",
		24: "DDR3",
		25: "FBD2",
		26: "DDR4",
	}
	FormFactorMap = map[uint16]string{
		0:  "Unknown",
		1:  "Other",
		2:  "SIP",
		3:  "DIP",
		4:  "ZIP",
		5:  "SOJ",
		6:  "Proprietary",
		7:  "SIMM",
		8:  "DIMM",
		9:  "TSOP",
		10: "PGA",
		11: "RIMM",
		12: "SODIMM",
		13: "SRIMM",
		14: "SMD",
		15: "SSMP",
		16: "QFP",
		17: "TQFP",
		18: "SOIC",
		19: "LCC",
		20: "PLCC",
		21: "BGA",
		22: "FPBGA",
		23: "LGA",
	}
	TypeDetailMap = map[uint16]string{
		1:    "Reserved",
		2:    "Other",
		4:    "Unknown",
		8:    "Fast-paged",
		16:   "Static column",
		32:   "Pseudo-static",
		64:   "RAMBUS",
		128:  "Synchronous",
		256:  "CMOS",
		512:  "EDO",
		1024: "Window DRAM",
		2048: "Cache DRAM",
		4096: "Non-volatile",
	}
)

func GetMemoryDevices() ([]*MemoryDeviceInformation, error) {
	dst, err := windows.Query[Win32_PhysicalMemory]()
	if err != nil {
		return nil, err
	}
	devices := make([]*MemoryDeviceInformation, 0, len(dst))
	for _, memory := range dst {
		devices = append(devices, &MemoryDeviceInformation{
			TotalWidth:           strconv.Itoa(int(memory.TotalWidth)) + " bits",
			DataWidth:            strconv.Itoa(int(memory.DataWidth)) + " bits",
			Size:                 strconv.FormatUint(memory.Capacity>>20, 10) + " MB",
			FormFactor:           FormFactorMap[memory.FormFactor],
			Locator:              memory.DeviceLocator,
			BankLocator:          memory.BankLabel,
			Type:                 TypeMap[memory.MemoryType],
			TypeDetail:           TypeDetailMap[memory.TypeDetail],
			Speed:                strconv.Itoa(int(memory.Speed)) + " MHz",
			Manufacturer:         memory.Manufacturer,
			SerialNumber:         memory.SerialNumber,
			AssetTag:             memory.Tag,
			PartNumber:           memory.PartNumber,
			Rank:                 strconv.Itoa(int(memory.Attributes)),
			ConfiguredClockSpeed: strconv.Itoa(int(memory.ConfiguredClockSpeed)) + " MHz",
			MinimumVoltage:       strconv.FormatFloat(float64(memory.MinVoltage)/1000, 'g', 2, 64) + " V",
			MaximumVoltage:       strconv.FormatFloat(float64(memory.MaxVoltage)/1000, 'g', 2, 64) + " V",
			ConfiguredVoltage:    strconv.FormatFloat(float64(memory.ConfiguredVoltage)/1000, 'g', 2, 64) + " V",
		})
	}
	return devices, nil
}
