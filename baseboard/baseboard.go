package baseboard

type Information struct {
	Manufacturer string   `yaml:"Manufacturer"`
	ProductName  string   `yaml:"Product Name"`
	Version      string   `yaml:"Version"`
	SerialNumber string   `yaml:"Serial Number"`
	AssetTag     string   `yaml:"Asset Tag"`
	Features     []string `yaml:"Features"`
	Type         string   `yaml:"Type"`
}
