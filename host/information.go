package host

type HostInformation struct {
	StaticHostname  string `yaml:"Static hostname"`
	IconName        string `yaml:"Icon name"`
	Chassis         string `yaml:"Chassis"`
	MachineID       string `yaml:"Machine ID"`
	BootID          string `yaml:"Boot ID"`
	Virtualization  string `yaml:"Virtualization"`
	OperatingSystem string `yaml:"Operating System"`
	CPEOSName       string `yaml:"CPE OS Name"`
	Kernel          string `yaml:"Kernel"`
	Architecture    string `yaml:"Architecture"`
}
