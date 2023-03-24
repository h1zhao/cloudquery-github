package installations

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client/mocks"
)

func buildInstallations(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs github.Installation
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	total := 1
	mock.EXPECT().ListInstallations(gomock.Any(), "testorg", gomock.Any()).Return(
		&github.OrganizationInstallations{TotalCount: &total, Installations: []*github.Installation{&cs}}, &github.Response{}, nil)

	return client.GithubServices{Organizations: mock}
}

func TestInstallations(t *testing.T) {
	client.GithubMockTestHelper(t, Installations(), buildInstallations, client.TestOptions{})
}
