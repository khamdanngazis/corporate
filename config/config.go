package config

import (
	"os"

	logger "../logging"
	"../parser"
)

// Configuration stores global configuration loaded from json file
type Configuration struct {
	ListenPort           string `yaml:"listenPort"`
	CertificateFile      string `yaml:"certificateFile"`
	KeyFile              string `yaml:"keyFile"`
	Query                string `yaml:"query"`
	DBUrl                string `yaml:"dbUrl"`
	DBType               string `yaml:"dbType"`
	ZSmartPath           string `yaml:"zsmartPath"`
	ZSmartUserName       string `yaml:"zsmartUserName"`
	ZSmartPassword       string `yaml:"zsmartPassword"`
	ZSmartZtePrivilegeId string `yaml:"zsmartztePrivilegeId"`
	ZSmartRequestID      string `yaml:"zsmartRequestID"`
	TokenAuth            string `yaml:"tokenAuth"`

	Log struct {
		FileName string `yaml:"filename"`
		Level    string `yaml:"level"`
	} `yaml:"log"`
}

// Param use as global variable for configuration
var Param Configuration

// LoadConfigFromFile use to load global configuration
func LoadConfigFromFile(fn *string) {
	if err := parser.LoadYAML(fn, &Param); err != nil {
		logger.Errorf("LoadConfigFromFile() - Failed opening config file %s\n%s", &fn, err)
		os.Exit(1)
	}
	//logger.Logf("Loaded configs: %v", Param)
	logger.Logf("Config %s", "Loaded")
}
