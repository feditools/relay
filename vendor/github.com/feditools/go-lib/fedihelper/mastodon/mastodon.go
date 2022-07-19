package mastodon

import (
	"context"

	"github.com/feditools/go-lib/fedihelper"
	mastodon "github.com/mattn/go-mastodon"
	"golang.org/x/sync/singleflight"
)

// Helper is a mastodon helper.
type Helper struct {
	fedi *fedihelper.FediHelper
	kv   fedihelper.KV

	appClientName string
	appWebsite    string
	externalURL   string

	registerAppGroup singleflight.Group
}

// New returns a new mastodon helper.
func New(k fedihelper.KV, appClientName, appWebsite, externalURL string) (*Helper, error) {
	return &Helper{
		kv: k,

		appClientName: appClientName,
		appWebsite:    appWebsite,
		externalURL:   externalURL,
	}, nil
}

// newClient return new mastodon API client.
func (h *Helper) newClient(ctx context.Context, instance fedihelper.Instance, accessToken string) (*mastodon.Client, error) {
	l := logger.WithField("func", "newClient")

	// get oauth config
	clientID, clientSecret, err := h.kv.GetInstanceOAuth(ctx, instance.GetID())
	if err != nil {
		l.Errorf("kv get: %s", err.Error())

		return nil, err
	}

	// create client
	client := mastodon.NewClient(&mastodon.Config{
		Server:       "https://" + instance.GetDomain(),
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AccessToken:  accessToken,
	})

	// apply custom transport
	client.Transport = h.fedi.HTTP().Transport()

	return client, nil
}

// GetSoftware returns the software type of this module.
func (*Helper) GetSoftware() fedihelper.Software { return fedihelper.SoftwareMastodon }

// SetFedi adds the fedi module to a helper.
func (h *Helper) SetFedi(f *fedihelper.FediHelper) {
	h.fedi = f
}
