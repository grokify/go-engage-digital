package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/oauth2more"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"

	"github.com/grokify/go-ringcentral-engage/engagedigital"
)

const (
	ApiUrlFormat = `https://%s.api.engagement.dimelo.com/1.0`
)

type options struct {
	Site  string `short:"s" long:"site" description:"A site" required:"true"`
	Token string `short:"t" long:"token" description:"A token" required:"true"`
}

func NewApiClient(site, token string) *engagedigital.APIClient {
	cfg := engagedigital.NewConfiguration()
	cfg.HTTPClient = oauth2more.NewClientBearerTokenSimple(token)
	cfg.BasePath = fmt.Sprintf(ApiUrlFormat, site)
	return engagedigital.NewAPIClient(cfg)
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := NewApiClient(opts.Site, opts.Token)

	info, resp, err := client.SourcesApi.GetAllSources(context.Background(), nil)
	if err != nil {
		err = errors.Wrap(err, "E_GET_ALL_SOURCES")
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)

	for _, source := range info.Records {
		fmt.Println(">>> INITIAL <<<")
		fmtutil.PrintJSON(source)
		newActive := false
		if !source.Active {
			newActive = true
		}
		opts := &engagedigital.UpdateSourceOpts{
			Active: optional.NewBool(newActive),
		}
		info, resp, err := client.SourcesApi.UpdateSource(
			context.Background(), source.Id, opts)
		if err != nil {
			err = errors.Wrap(err, "E_UPDATE_SOURCE")
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmt.Println(">>> UPDATED <<<")
		fmtutil.PrintJSON(info)
		break
	}

	fmt.Println("DONE")
}
