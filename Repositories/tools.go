package Repositories

import (
	models "ds_nostr_go/Models"
	"encoding/json"
	"os"
)

const configPath = "Config/configs.json"

type Tools struct{}

func NewTools() *Tools {
	return &Tools{}
}

func (rep *Tools) ReadKeysJson() (configs map[string]string, err error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config models.KeysConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	configs = make(map[string]string)
	configs["sk"] = config.Sk
	configs["pk"] = config.Pk
	configs["nsec"] = config.Nsec
	configs["npub"] = config.Npub
	configs["postgres"] = config.Postgres
	return configs, nil
}
