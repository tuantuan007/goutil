//go:build windows

package baseboard

import (
	"goutil/common/windows"
	"time"
)

type Win32_BaseBoard struct {
	Caption                 string
	ConfigOptions           []string
	CreationClassName       string
	Depth                   int
	Description             string
	Height                  int
	HostingBoard            bool
	HotSwappable            bool
	InstallDate             time.Time
	Manufacturer            string
	Model                   string
	Name                    string
	OtherIdentifyingInfo    string
	PartNumber              string
	PoweredOn               bool
	Product                 string
	Removable               bool
	Replaceable             bool
	RequirementsDescription string
	RequiresDaughterBoard   bool
	SerialNumber            string
	SKU                     string
	SlotLayout              string
	SpecialRequirements     bool
	Status                  string
	Tag                     string
	Version                 string
	Weight                  int
	Width                   int
}

func GetBaseboardInformation() (*BaseboardInformation, error) {
	dst, err := windows.Query[Win32_BaseBoard]()
	if err != nil {
		return nil, err
	}
	var features []string
	if dst[0].HostingBoard {
		features = append(features, "HostingBoard")
	}
	if dst[0].HotSwappable {
		features = append(features, "HotSwappable")
	}
	if dst[0].Replaceable {
		features = append(features, "Replaceable")
	}
	if dst[0].Removable {
		features = append(features, "Removable")
	}
	info := &BaseboardInformation{
		Manufacturer: dst[0].Manufacturer,
		ProductName:  dst[0].Product,
		Version:      dst[0].Version,
		SerialNumber: dst[0].SerialNumber,
		AssetTag:     dst[0].Tag,
		Features:     features,
		Type:         dst[0].Description,
	}
	return info, nil
}
