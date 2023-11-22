//go:build windows

package graphics

import (
	"github.com/wxlbd/goutil/common/windows"
	"time"
)

type Win32_VideoController struct {
	AcceleratorCapabilities      []uint16  `json:"AcceleratorCapabilities"`
	AdapterCompatibility         string    `json:"AdapterCompatibility"`
	AdapterDACType               string    `json:"AdapterDACType"`
	AdapterRAM                   uint32    `json:"AdapterRAM"`
	Availability                 uint16    `json:"Availability"`
	CapabilityDescriptions       []string  `json:"CapabilityDescriptions"`
	Caption                      string    `json:"Caption"`
	ColorTableEntries            uint32    `json:"ColorTableEntries"`
	ConfigManagerErrorCode       uint32    `json:"ConfigManagerErrorCode"`
	ConfigManagerUserConfig      bool      `json:"ConfigManagerUserConfig"`
	CreationClassName            string    `json:"CreationClassName"`
	CurrentBitsPerPixel          uint32    `json:"CurrentBitsPerPixel"`
	CurrentHorizontalResolution  uint32    `json:"CurrentHorizontalResolution"`
	CurrentNumberOfColors        uint64    `json:"CurrentNumberOfColors"`
	CurrentNumberOfColumns       uint32    `json:"CurrentNumberOfColumns"`
	CurrentNumberOfRows          uint32    `json:"CurrentNumberOfRows"`
	CurrentRefreshRate           uint32    `json:"CurrentRefreshRate"`
	CurrentScanMode              uint16    `json:"CurrentScanMode"`
	CurrentVerticalResolution    uint32    `json:"CurrentVerticalResolution"`
	Description                  string    `json:"Description"`
	DeviceID                     string    `json:"DeviceID"`
	DeviceSpecificPens           uint32    `json:"DeviceSpecificPens"`
	DitherType                   uint32    `json:"DitherType"`
	DriverDate                   time.Time `json:"DriverDate"`
	DriverVersion                string    `json:"DriverVersion"`
	ErrorCleared                 bool      `json:"ErrorCleared"`
	ErrorDescription             string    `json:"ErrorDescription"`
	ICMIntent                    uint32    `json:"ICMIntent"`
	ICMMethod                    uint32    `json:"ICMMethod"`
	InfFilename                  string    `json:"InfFilename"`
	InfSection                   string    `json:"InfSection"`
	InstallDate                  time.Time `json:"InstallDate"`
	InstalledDisplayDrivers      string    `json:"InstalledDisplayDrivers"`
	LastErrorCode                uint32    `json:"LastErrorCode"`
	MaxMemorySupported           uint32    `json:"MaxMemorySupported"`
	MaxNumberControlled          uint32    `json:"MaxNumberControlled"`
	MaxRefreshRate               uint32    `json:"MaxRefreshRate"`
	MinRefreshRate               uint32    `json:"MinRefreshRate"`
	Monochrome                   bool      `json:"Monochrome"`
	Name                         string    `json:"Name"`
	NumberOfColorPlanes          uint16    `json:"NumberOfColorPlanes"`
	NumberOfVideoPages           uint32    `json:"NumberOfVideoPages"`
	PNPDeviceID                  string    `json:"PNPDeviceID"`
	PowerManagementCapabilities  []uint16  `json:"PowerManagementCapabilities"`
	PowerManagementSupported     bool      `json:"PowerManagementSupported"`
	ProtocolSupported            uint16    `json:"ProtocolSupported"`
	ReservedSystemPaletteEntries uint32    `json:"ReservedSystemPaletteEntries"`
	SpecificationVersion         uint32    `json:"SpecificationVersion"`
	Status                       string    `json:"Status"`
	StatusInfo                   uint16    `json:"StatusInfo"`
	SystemCreationClassName      string    `json:"SystemCreationClassName"`
	SystemName                   string    `json:"SystemName"`
	SystemPaletteEntries         uint32    `json:"SystemPaletteEntries"`
	TimeOfLastReset              time.Time `json:"TimeOfLastReset"`
	VideoArchitecture            uint16    `json:"VideoArchitecture"`
	VideoMemoryType              uint16    `json:"VideoMemoryType"`
	VideoMode                    uint16    `json:"VideoMode"`
	VideoModeDescription         string    `json:"VideoModeDescription"`
	VideoProcessor               string    `json:"VideoProcessor"`
}

func GetGraphicsCardInformation() (*GraphicsInformation, error) {
	dst, err := windows.Query[Win32_VideoController]()
	if err != nil {
		return nil, err
	}
	result := dst[0]
	return &GraphicsInformation{
		Manufacturer: result.AdapterCompatibility,
		DeviceID:     result.DeviceID,
		Name:         result.Name,
	}, nil
}
