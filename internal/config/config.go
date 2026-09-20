package config

import (
	"encoding/json"

	"github.com/vitorlfaria/go-observer/internal/types"
)

func Parse(data []byte) (types.Config, error) {
	var configFile types.Config
	if err := json.Unmarshal(data, &configFile); err != nil {
		return configFile, err
	}
	return configFile, nil
}
