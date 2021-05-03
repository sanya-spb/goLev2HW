package main

func main() {
	type confDatabase struct {
		Host string `toml:"host" yaml:"host" json:"host"`
		Port int    `toml:"port" yaml:"port" json:"port"`
		User string `toml:"user" yaml:"user" json:"user"`
		Pass string `toml:"pass" yaml:"pass" json:"pass"`
		Ssl  bool   `toml:"ssl" yaml:"ssl" json:"ssl"`
	}

	type confServer struct {
		Bind     []string `toml:"bind" yaml:"bind" json:"bind"`
		Port     int      `toml:"port" yaml:"port" json:"port"`
		LogLevel int      `toml:"log_level" yaml:"log_level" json:"log_level"`
	}

	type Config struct {
		Debug    bool         `toml:"debug" yaml:"debug" json:"debug"`
		MyUrl    string       `toml:"my_url" yaml:"my_url" json:"my_url"`
		Database confDatabase `toml:"database" yaml:"database" json:"database"`
		Server   confServer   `toml:"server" yaml:"server" json:"server"`
	}

}

func fillStruct(in *interface{}, filler map[string]interface{}) {

}
