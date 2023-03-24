package traffic

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client/mocks"
)

func buildViews(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	repositoriesMock := mocks.NewMockRepositoriesService(ctrl)

	var views *github.TrafficViews
	if err := faker.FakeObject(&views); err != nil {
		t.Fatal(err)
	}

	opts := github.TrafficBreakdownOptions{}

	repositoriesMock.EXPECT().ListTrafficViews(gomock.Any(), "test string", "test string", &opts).Return(views, nil, nil)
	return client.GithubServices{Repositories: repositoriesMock}
}

func TestViews(t *testing.T) {
	client.GithubMockTestHelper(t, Views(), buildViews, client.TestOptions{})
}
