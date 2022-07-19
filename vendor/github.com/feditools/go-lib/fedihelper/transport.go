package fedihelper

import (
	"bytes"
	"context"
	"crypto"
	"fmt"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/go-fed/activity/pub"
	"github.com/go-fed/httpsig"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	rfc1123WithoutZone = "Mon, 02 Jan 2006 15:04:05"
)

var (
	digestAlgo = httpsig.DigestSha256
	algoPrefs  = []httpsig.Algorithm{httpsig.RSA_SHA256}

	getHeaders  = []string{httpsig.RequestTarget, "host", "date"}
	postHeaders = []string{httpsig.RequestTarget, "host", "date", "digest", "content-type"}
)

// NewTransport creates a new Transport module
func NewTransport(clock pub.Clock, client HttpClient, pubKeyID string, privkey crypto.PrivateKey) (*Transport, error) {
	return &Transport{
		client: client,
		clock:  clock,

		keyID:   pubKeyID,
		privKey: privkey,
	}, nil
}

// Transport handled signing outgoing requests to federated instances
type Transport struct {
	client HttpClient
	clock  pub.Clock

	keyID   string
	privKey crypto.PrivateKey

	signerExp  time.Time
	getSigner  httpsig.Signer
	postSigner httpsig.Signer
	signerMu   sync.Mutex
}

func (t *Transport) Client() HttpClient {
	return t.client
}

func (t *Transport) InstanceGet(ctx context.Context, uri *url.URL, accepts ...libhttp.Mime) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", uri.String(), nil)
	if err != nil {
		return nil, NewErrorf("creating http request: %s", err.Error())
	}

	for _, accept := range accepts {
		req.Header.Add("Accept", accept.String())
	}
	req.Header.Add("Date", t.clock.Now().UTC().Format(rfc1123WithoutZone)+" GMT")
	req.Header.Set("Host", uri.Host)

	t.doSign(func() {
		err = t.getSigner.SignRequest(t.privKey, t.keyID, req, nil)
	})
	if err != nil {
		return nil, NewErrorf("can't lock signer: %s", err.Error())
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, NewErrorf("http client do: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get %s: %d-%s", uri.String(), resp.StatusCode, resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

func (t *Transport) InstancePost(ctx context.Context, uri *url.URL, body []byte, contentType libhttp.Mime, accepts ...libhttp.Mime) ([]byte, error) {
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bodyReader)
	if err != nil {
		return nil, NewErrorf("creating http request: %s", err.Error())
	}

	for _, accept := range accepts {
		req.Header.Add("Accept", accept.String())
	}
	req.Header.Add("Date", t.clock.Now().UTC().Format(rfc1123WithoutZone)+" GMT")
	req.Header.Set("Host", uri.Host)
	req.Header.Set("Content-Type", contentType.String())

	t.doSign(func() {
		err = t.postSigner.SignRequest(t.privKey, t.keyID, req, body)
	})
	if err != nil {
		return nil, NewErrorf("can't lock signer: %s", err.Error())
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, NewErrorf("http client do: %s", err.Error())
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewErrorf("can't read body: %s", err.Error())
	}

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return nil, NewErrorf("http post %s: %s-%s", uri.String(), resp.Status, string(respBody))
	}
	return respBody, nil
}

func (t *Transport) doSign(do func()) {
	// Perform within mu safety
	t.signerMu.Lock()
	defer t.signerMu.Unlock()

	if now := t.clock.Now(); now.After(t.signerExp) {
		const expiry = 120

		// Signers have expired and require renewal
		t.getSigner, _ = genGetSigner(expiry)
		t.postSigner, _ = genPostSigner(expiry)
		t.signerExp = now.Add(time.Second * expiry)
	}

	// Perform signing
	do()
}

func genGetSigner(expiresIn int64) (httpsig.Signer, error) {
	sig, _, err := httpsig.NewSigner(algoPrefs, digestAlgo, getHeaders, httpsig.Signature, expiresIn)
	return sig, err
}

func genPostSigner(expiresIn int64) (httpsig.Signer, error) {
	sig, _, err := httpsig.NewSigner(algoPrefs, digestAlgo, postHeaders, httpsig.Signature, expiresIn)
	return sig, err
}
