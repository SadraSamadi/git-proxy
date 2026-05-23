package store

import "regexp"

type Data struct {
	Proxies map[string]string `json:"proxies"`
}

func getDefault() *Data {
	data := &Data{
		Proxies: map[string]string{},
	}
	return data
}

func Validate(value string) bool {
	pattern := regexp.MustCompile("^(http|socks5h?)://[a-zA-Z0-9._-]+(:\\d{1,5})?$")
	return pattern.MatchString(value)
}
