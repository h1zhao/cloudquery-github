package traffic

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client/mocks"
)

func buildReferrers(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	repositoriesMock := mocks.NewMockRepositoriesService(ctrl)

	var referrers []*github.TrafficReferrer
	if err := faker.FakeObject(&referrers); err != nil {
		t.Fatal(err)
	}

	repositoriesMock.EXPECT().ListTrafficReferrers(gomock.Any(), "test string", "test string").Return(referrers, nil, nil)
	return client.GithubServices{Repositories: repositoriesMock}
}

func TestReferrers(t *testing.T) {
	client.GithubMockTestHelper(t, Referrers(), buildReferrers, client.TestOptions{})
}
