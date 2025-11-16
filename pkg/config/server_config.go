package config

type HttpServerConfig struct {
	Hostname   string `mapstructure:"hostname"`
	BaseUrl    string `mapstructure:"baseUrl"`
	TlsEnabled bool   `mapstructure:"tlsEnabled"`
	CertFile   string `mapstructure:"certFile"`
	KeyFile    string `mapstructure:"keyFile"`
	Port       int    `mapstructure:"port"`
}

type GrpcServerConfig struct {
	Port int `mapstructure:"port"`
}

type GatewayConfig struct {
	HttpServerConfig *HttpServerConfig
	GrpcServerConfig *GrpcServerConfig
}
