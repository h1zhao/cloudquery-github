package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/actions"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/billing"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/external"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/hooks"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/installations"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/issues"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/organizations"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/repositories"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/teams"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/resources/services/traffic"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		actions.Workflows(),
		billing.Action(),
		billing.Storage(),
		billing.Package(),
		external.Groups(),
		issues.Issues(),
		hooks.Hooks(),
		installations.Installations(),
		organizations.Organizations(),
		repositories.Repositories(),
		teams.Teams(),
		traffic.Clones(),
		traffic.Paths(),
		traffic.Views(),
		traffic.Referrers(),
	}
}
