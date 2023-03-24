package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/h1zhao/cq-source-cloudquery-github/resources/plugin"
)

const sentryDSN = "https://99f66f1a627f48deb66e49a25d6028a6@o1396617.ingest.sentry.io/6747628"

func main() {
	serve.Source(plugin.Plugin(), serve.WithSourceSentryDSN(sentryDSN))
}
