package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/client"
)

func alerts() *schema.Table {
	return &schema.Table{
		Name:      "github_repository_dependabot_alerts",
		Resolver:  fetchAlerts,
		Transform: client.TransformWithStruct(&github.DependabotAlert{}, transformers.WithPrimaryKeys("Number")),
		Columns:   []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)
	opts := &github.ListAlertsOptions{ListCursorOptions: github.ListCursorOptions{PerPage: 99}}

	for {
		alerts, resp, err := c.Github.Dependabot.ListRepoAlerts(ctx, c.Org, *repo.Name, opts)
		if err != nil {
			return err
		}
		res <- alerts
		opts.Page = strconv.FormatInt(int64(resp.NextPage), 10)
		fmt.Printf("Repo %s continue page %d\n", *repo.Name, resp.NextPage)
		if resp.NextPage == 0 {
			fmt.Printf("Repo %s next page is %d. Stop\n", *repo.Name, resp.NextPage)
			break
		}
	}

	return nil
}
