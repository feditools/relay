package transport

import (
	"bytes"
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

	t.doSign(func() {
		err = t.getSigner.SignRequest(t.privKey, t.keyID, req, nil)
	})
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

func (t *Transport) InstancePost(ctx context.Context, uri *url.URL, body []byte, contentType string, accepts ...string) ([]byte, error) {
	l := logger.WithField("func", "InstancePost")

	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bodyReader)
	if err != nil {
		l.Errorf("creating http request: %s", err.Error())
		return nil, err
	}

	for _, accept := range accepts {
		req.Header.Add("Accept", accept)
	}
	req.Header.Add("Date", t.clock.Now().UTC().Format(rfc1123WithoutZone)+" GMT")
	req.Header.Set("Host", uri.Host)
	req.Header.Set("Content-Type", contentType)

	t.doSign(func() {
		err = t.postSigner.SignRequest(t.privKey, t.keyID, req, body)
	})
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
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Warnf("can't read body: %s", err.Error())
	}

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return nil, fmt.Errorf("http post %s: %s-%s", uri.String(), resp.Status, string(respBody))
	}
	return ioutil.ReadAll(resp.Body)
}
