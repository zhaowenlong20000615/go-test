package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

type TomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

var Config *TomlConfig

func init() {
	Config = new(TomlConfig)
	var err error
	Config.System.CurrentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	Config.System.AppName = "go-blog"
	Config.System.Version = 1.0
	_, err = toml.DecodeFile("config/config.toml", &Config)
	if err != nil {
		panic(err)
	}
}
