package internal

type HiveConfig struct {
	Token        string `mapstructure:"token"`
	Endpoint     string `mapstructure:"endpoint"`
	Organization string `mapstructure:"organization"`
}

func (c *HiveConfig) extendConfig(o *HiveConfig) *HiveConfig {
	cfg := &HiveConfig{
		Token:        o.Token,
		Endpoint:     o.Endpoint,
		Organization: o.Organization,
	}
	if c.Token != "" {
		cfg.Token = c.Token
	}
	if c.Endpoint != "" {
		cfg.Endpoint = c.Endpoint
	}
	if c.Organization != "" {
		cfg.Organization = c.Organization
	}

	return cfg
}
