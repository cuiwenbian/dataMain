package conf

const CONFIG = "config"

// Config Config struct
type Config struct {
	Mongodb Mongodb `yaml:"mongodb"`
}

type Mongodb struct {
	Server   string `yaml:"server"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
