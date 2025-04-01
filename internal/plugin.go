package internal

import (
	"fmt"

	"github.com/mach-composer/mach-composer-plugin-helpers/helpers"
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

type Plugin struct {
	environment  string
	provider     string
	globalConfig *HiveConfig
	siteConfigs  map[string]*HiveConfig
	enabled      bool
}

func NewHivePlugin() schema.MachComposerPlugin {
	state := &Plugin{
		provider:    "0.2.0",
		siteConfigs: map[string]*HiveConfig{},
	}
	return plugin.NewPlugin(&schema.PluginSchema{
		Identifier: "hive",

		Configure: state.Configure,
		IsEnabled: state.IsEnabled,

		GetValidationSchema: state.GetValidationSchema,

		// Config
		SetGlobalConfig: state.SetGlobalConfig,
		SetSiteConfig:   state.SetSiteConfig,

		// Renders
		RenderTerraformProviders: state.TerraformRenderProviders,
		RenderTerraformResources: state.TerraformRenderResources,
		RenderTerraformComponent: state.RenderTerraformComponent,
	})
}

func (p *Plugin) Configure(environment string, provider string) error {
	p.environment = environment
	if provider != "" {
		p.provider = provider
	}
	return nil
}

func (p *Plugin) IsEnabled() bool {
	return p.enabled
}

func (p *Plugin) GetValidationSchema() (*schema.ValidationSchema, error) {
	result := getSchema()
	return result, nil
}

func (p *Plugin) SetGlobalConfig(data map[string]any) error {
	cfg := HiveConfig{}
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.globalConfig = &cfg
	p.enabled = true
	return nil
}

func (p *Plugin) SetSiteConfig(site string, data map[string]any) error {
	if data == nil {
		return nil
	}

	cfg := HiveConfig{}
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.siteConfigs[site] = &cfg
	p.enabled = true
	return nil
}

func (p *Plugin) TerraformRenderProviders(site string) (string, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return "", nil
	}

	result := fmt.Sprintf(`
		hive = {
			source = "labd/hive"
			version = "%s"
		}`, helpers.VersionConstraint(p.provider))
	return result, nil
}

func (p *Plugin) TerraformRenderResources(site string) (string, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return "", nil
	}

	template := `
		provider "hive" {
	       {{ renderProperty "token" .Token }}
	       {{ renderOptionalProperty "endpoint" .Endpoint }}
	    }
	`
	return helpers.RenderGoTemplate(template, cfg)
}

func (p *Plugin) getSiteConfig(site string) *HiveConfig {
	if p.globalConfig == nil {
		return nil
	}
	cfg, ok := p.siteConfigs[site]
	if !ok {
		cfg = &HiveConfig{}
	}
	return cfg.extendConfig(p.globalConfig)
}

func (p *Plugin) RenderTerraformComponent(site string, _ string) (*schema.ComponentSchema, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return nil, nil
	}

	result := &schema.ComponentSchema{
		Providers: []string{
			"hive = hive",
		},
		DependsOn: []string{},
	}

	return result, nil
}
