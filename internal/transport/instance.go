package transport

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (t *Transport) InstanceGet(ctx context.Context, uri *url.URL, accepts ...string) ([]byte, error) {
	l := logger.WithField("func", "InstanceGet")

	req, err := http.NewRequestWithContext(ctx, "GET", uri.String(), nil)
	if err != nil {
		l.Errorf("creating http request: %s", err.Error())
		return nil, err
	}

	for _, accept := range accepts {
		req.Header.Add("Accept", accept)
	}
	req.Header.Add("Date", t.clock.Now().UTC().Format(rfc1123WithoutZone)+" GMT")
	req.Header.Set("Host", uri.Host)
	t.getSignerLock.Lock()
	err = t.getSigner.SignRequest(t.privKey, t.keyID, req, nil)
	t.getSignerLock.Unlock()
	if err != nil {
		l.Errorf("can't lock signer: %s", err.Error())
		return nil, err
	}
	resp, err := t.client.Do(req)
	if err != nil {
		l.Errorf("http client do: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get %s: %d-%s", uri.String(), resp.StatusCode, resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}
