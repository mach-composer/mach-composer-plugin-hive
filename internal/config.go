package internal

type HiveConfig struct {
	Token    string `mapstructure:"token"`
	Endpoint string `mapstructure:"endpoint"`
}

func (c *HiveConfig) extendConfig(o *HiveConfig) *HiveConfig {
	cfg := &HiveConfig{
		Token:    o.Token,
		Endpoint: o.Endpoint,
	}
	if c.Token != "" {
		cfg.Token = c.Token
	}
	if c.Endpoint != "" {
		cfg.Endpoint = c.Endpoint
	}
	return cfg
}
