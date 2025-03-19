package main

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"

	"github.com/mach-composer/mach-composer-plugin-hive/internal"
)

func main() {
	p := internal.NewHivePlugin()
	plugin.ServePlugin(p)
}
