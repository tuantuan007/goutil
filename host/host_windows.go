//go:build windows

package host

import "goutil/common/windows"

type Win32_OperatingSystem struct {
	Version              string
	OSArchitecture       string
	Organization         string
	Name                 string
	BootDevice           string
	BuildNumber          string
	BuildType            string
	Caption              string
	CodeSet              string
	CountryCode          string
	CreationClassName    string
	CSCreationClassName  string
	CSDVersion           string
	CSName               string
	Description          string
	Locale               string
	Manufacturer         string
	MUILanguages         []string
	OtherTypeDescription string
	PlusProductID        string
	PlusVersionNumber    string
	RegisteredUser       string
	SerialNumber         string
	Status               string
	SystemDevice         string
	SystemDirectory      string
	SystemDrive          string
}

func GetHostInformation() (*HostInformation, error) {
	results, err := windows.Query[Win32_OperatingSystem]()
	if err != nil {
		return nil, err
	}
	result := results[0]
	return &HostInformation{
		StaticHostname:  result.CSName,
		IconName:        result.Name,
		Chassis:         result.BuildType,
		MachineID:       "",
		BootID:          "",
		Virtualization:  "",
		OperatingSystem: result.Caption,
		CPEOSName:       result.Name,
		Kernel:          result.Version,
		Architecture:    result.OSArchitecture,
	}, nil
}
