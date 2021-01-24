package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var Auth *Config

type Config struct {
	AuthConfig *AuthConfig `toml:"mysql" json:"mysql"`
	HttpConfig   *HttpConfig   `toml:"http"`
}

type AuthConfig struct {
	Name   string `toml:"name" json:"sky_addr"`
	Passwd string `toml:"passwd" json:"sky_pprof_addr "`
	Ip     string `toml:"ip" json:"alarm_url"`
	Port   string `toml:"port" json:"is_debug"`
	Store  string `toml:"store" json:"log_level"`
}

type HttpConfig struct {
	Ip   string   `toml:"ip"`
	Port string   `toml:"port"`
}

func Init(path string) error {
	Auth = &Config{AuthConfig: &AuthConfig{},HttpConfig: &HttpConfig{}}
	if _, err := toml.DecodeFile(path, Auth); err != nil {
		fmt.Printf("toml decode err %v", err)
		return err
	}
	return nil
}
