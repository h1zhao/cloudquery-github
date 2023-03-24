package traffic

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client/mocks"
)

func buildClones(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	repositoriesMock := mocks.NewMockRepositoriesService(ctrl)

	var clones *github.TrafficClones
	if err := faker.FakeObject(&clones); err != nil {
		t.Fatal(err)
	}

	opts := github.TrafficBreakdownOptions{}

	repositoriesMock.EXPECT().ListTrafficClones(gomock.Any(), "test string", "test string", &opts).Return(clones, nil, nil)
	return client.GithubServices{Repositories: repositoriesMock}
}

func TestClones(t *testing.T) {
	client.GithubMockTestHelper(t, Clones(), buildClones, client.TestOptions{})
}
