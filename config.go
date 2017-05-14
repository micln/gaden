package main

type Config struct {
	Ignores []string
	Files   []string
}


func NewConfigFromFile() *Config {
	cfg := &Config{}
	return cfg
}
