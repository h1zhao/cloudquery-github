package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
)

var (
	Version = "Development"
)

func Plugin() *source.Plugin {
	allTables := Tables()
	// here you can append custom non-generated tables
	return source.NewPlugin(
		"github",
		Version,
		allTables,
		client.Configure,
	)
}
