package Repositories

import (
	models "ds_nostr_go/Models"
	"encoding/json"
	"os"
)

const configPath = "Config/keys.json"

type Tools struct{}

func NewTools() *Tools {
	return &Tools{}
}

func (rep *Tools) ReadKeysJson() (configKey map[string]string, err error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config models.KeysConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	configKey = make(map[string]string)
	configKey["sk"] = config.Sk
	configKey["pk"] = config.Pk
	configKey["nsec"] = config.Nsec
	configKey["npub"] = config.Npub

	return configKey, nil
}
