package mastodon

import (
	"context"
	"net/url"

	"github.com/feditools/go-lib/fedihelper"
)

// GetAccessToken gets an access token for a account from a returned code.
func (h *Helper) GetAccessToken(ctx context.Context, instance fedihelper.Instance, code string) (accessToken string, err error) {
	// decrypt secret
	c, err := h.newClient(ctx, instance, "")
	if err != nil {
		return "", err
	}

	// authenticate
	instanceToken := h.fedi.GetTokenHandler(ctx, instance)
	err = c.AuthenticateToken(ctx, code, h.externalURL+"/callback/oauth/"+instanceToken)
	if err != nil {
		return "", err
	}

	return c.Config.AccessToken, nil
}

// MakeLoginURI creates a login redirect url for mastodon.
func (h *Helper) MakeLoginURI(ctx context.Context, instance fedihelper.Instance) (*url.URL, error) {
	l := logger.WithField("func", "MakeLoginURI")

	clientID, _, err := h.kv.GetInstanceOAuth(ctx, instance.GetID())
	if err != nil {
		l.Errorf("kv get: %s", err.Error())

		return nil, err
	}

	instanceToken := h.fedi.GetTokenHandler(ctx, instance)
	u := &url.URL{
		Scheme: "https",
		Host:   instance.GetDomain(),
		Path:   "/oauth/authorize",
	}
	q := u.Query()
	q.Set("client_id", clientID)
	q.Set("redirect_uri", h.externalURL+"/callback/oauth/"+instanceToken)
	q.Set("response_type", "code")
	q.Set("scope", "read:accounts")
	u.RawQuery = q.Encode()

	return u, nil
}
