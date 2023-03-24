package organizations

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/client"
	"github.com/h1zhao/cq-source-cloudquery-github/client/mocks"
)

func buildDependabot(t *testing.T, ctrl *gomock.Controller) client.DependabotService {
	mock := mocks.NewMockDependabotService(ctrl)

	var alert github.DependabotAlert
	if err := faker.FakeObject(&alert); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListOrgAlerts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.DependabotAlert{&alert}, &github.Response{}, nil)

	var secret github.Secret
	if err := faker.FakeObject(&secret); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListOrgSecrets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&github.Secrets{TotalCount: 1, Secrets: []*github.Secret{&secret}}, &github.Response{}, nil)

	return mock
}
