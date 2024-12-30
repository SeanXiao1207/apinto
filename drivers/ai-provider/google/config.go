package google

import (
	"fmt"
	"net/url"

	"github.com/eolinker/eosc"
)

type Config struct {
	APIKey string `json:"google_api_key"`
	Base   string `json:"google_api_base"`
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
	return conf, nil
}
