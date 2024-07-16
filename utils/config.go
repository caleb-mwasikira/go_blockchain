package utils

import (
	"encoding/json"
	"errors"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type KeysConfig struct {
	PrivateKeyFile string `json:"private"`
	PublicKeyFile  string `json:"public"`
}

type ServerConfig struct {
	Host     string     `json:"host"`
	MinPort  int        `json:"min_port"`
	MaxPort  int        `json:"max_port"`
	LogFile  string     `json:"log_file"`
	CertFile string     `json:"cert_file"`
	Keys     KeysConfig `json:"keys"`
}

var (
	ErrInvalidConfigSection error = errors.New("invalid config section")
)

func LoadServerConfig(section_name string) (*ServerConfig, error) {
	// load configuration files
	config_fpath := filepath.Join(ProjectPath, "config.json")
	viper.SetConfigFile(config_fpath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file; %v", err)
	}

	var server_config ServerConfig
	section := viper.Get(section_name)
	if section == nil {
		return nil, ErrInvalidConfigSection
	}

	defer func() {
		if err := recover(); err != nil {
			// panic caught
			log.Fatal(ErrInvalidConfigSection)
		}
	}()

	data, err := json.Marshal(section.(map[string]interface{}))
	if err != nil {
		return nil, ErrInvalidConfigSection
	}

	err = json.Unmarshal(data, &server_config)
	if err != nil {
		return nil, err
	}
	return &server_config, nil
}
