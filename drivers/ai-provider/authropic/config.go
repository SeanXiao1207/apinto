package anthropic

import (
	"fmt"
	"net/url"

	"github.com/eolinker/eosc"
)

const defaultVersion = "2023-06-01"

type Config struct {
	APIKey  string `json:"anthropic_api_key"`
	Base    string `json:"anthropic_api_url"`
	Version string `json:"anthropic_api_version"`
}

func checkConfig(v interface{}) (*Config, error) {
	conf, ok := v.(*Config)
	if !ok {
		return nil, eosc.ErrorConfigType
	}
	if conf.APIKey == "" {
		return nil, fmt.Errorf("api_key is required")
	}
	if conf.Base != "" {
		u, err := url.Parse(conf.Base)
		if err != nil {
			return nil, fmt.Errorf("base url is invalid")
		}
		if u.Scheme == "" || u.Host == "" {
			return nil, fmt.Errorf("base url is invalid")
		}
	}
	if conf.Version == "" {
		conf.Version = defaultVersion
	}
	return conf, nil
}
