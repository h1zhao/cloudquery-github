package billing

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client/mocks"
)

func buildPackage(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockBillingService(ctrl)

	var cs *github.PackageBilling
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetPackagesBillingOrg(gomock.Any(), "testorg").Return(cs, &github.Response{}, nil)
	return client.GithubServices{Billing: mock}
}

func TestPackageBillings(t *testing.T) {
	client.GithubMockTestHelper(t, Package(), buildPackage, client.TestOptions{})
}
