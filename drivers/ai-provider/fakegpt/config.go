package fakegpt

import (
	"fmt"

	"github.com/eolinker/eosc"
)

type Config struct {
	APIKey string `json:"apikey"`
}

func checkConfig(v interface{}) (*Config, error) {
	conf, ok := v.(*Config)
	if !ok {
		return nil, eosc.ErrorConfigType
	}
	if conf.APIKey == "" {
		return nil, fmt.Errorf("apikey is required")
	}
	return conf, nil
}
