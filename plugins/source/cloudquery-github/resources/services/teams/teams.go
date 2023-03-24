package teams

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:      "github_teams",
		Resolver:  fetchTeams,
		Multiplex: client.OrgMultiplex,
		Transform: client.TransformWithStruct(&github.Team{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{members(), repositories()},
	}
}

func fetchTeams(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{PerPage: 100}
	for {
		repos, resp, err := c.Github.Teams.ListTeams(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- repos
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}

var teamIDColumn = schema.Column{
	Name:            "team_id",
	Type:            schema.TypeInt,
	Resolver:        schema.ParentColumnResolver("id"),
	CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
}
