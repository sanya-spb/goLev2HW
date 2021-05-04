package config

import (
	"flag"
	"os"
	"strconv"
)

// const (
// 	_ = iota
// 	ERROR
// 	WARN
// 	NOTICE
// )

type Config struct {
	Debug   bool
	Dirs    []string
	Hash    string
	MaxSize uint64
	Size    bool
	// LogLevel int
}

// Init config
func NewConfig() *Config {
	var result *Config = new(Config)

	// Determine the initial directories.
	result.Dirs = os.Args[1:]
	if len(result.Dirs) == 0 {
		result.Dirs = []string{"."}
	}

	flag.BoolVar(&result.Debug, "debug", GetEnvBool("DEBUG", false), "Output verbose debug information")
	// flag.StringVar(&result.Dirs, "dirs", GetEnv("DIR", "./"), "Path to scan files to doubles, default current directory")
	flag.StringVar(&result.Hash, "hash", GetEnv("HASH", "md5"), "HASH method: [md5, crc32], default md5")
	flag.Uint64Var(&result.MaxSize, "max-size", GetEnvUInt("MAX_SIZE", 0), "limit maximum file size for checking, 0 - disable")
	flag.BoolVar(&result.Size, "size", GetEnvBool("SIZE", true), "Compare files by size")
	flag.Parse()

	return result
}

// Get uint value from ENV
func GetEnvUInt(key string, defaultVal uint64) uint64 {
	if envVal, ok := os.LookupEnv(key); ok {
		if envBool, err := strconv.ParseUint(envVal, 10, 64); err == nil {
			return envBool
		}
	}
	return defaultVal
}

// Get string value from ENV
func GetEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

// Get bool value from ENV
func GetEnvBool(key string, defaultVal bool) bool {
	if envVal, ok := os.LookupEnv(key); ok {
		if envBool, err := strconv.ParseBool(envVal); err == nil {
			return envBool
		}
	}
	return defaultVal
}
