package store

import "regexp"

type Config struct {
	Proxies map[string]string `json:"proxies"`
}

func DefaultConfig() *Config {
	config := &Config{
		Proxies: map[string]string{},
	}
	return config
}

func Validate(value string) bool {
	pattern := regexp.MustCompile("^(http|socks5h?)://[a-zA-Z0-9._-]+(:\\d{1,5})?$")
	return pattern.MatchString(value)
}
