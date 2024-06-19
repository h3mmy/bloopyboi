package config

type ArrClientParam string

const ApiKey = "apiKey"
const AppURL = "appUrl"

type ArrClientConfig struct {
	Name string
	Type string
	URL string
	ApiKey string
}

func (cfg *ArrClientConfig) ToParams() map[ArrClientParam]string {
	params := make(map[ArrClientParam]string, 2)
	params[ApiKey] = cfg.ApiKey
	params[AppURL] = cfg.URL
	return params
}
