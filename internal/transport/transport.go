package transport

import (
	"crypto"
	"fmt"
	"github.com/feditools/relay/internal/http"
	"github.com/go-fed/activity/pub"
	"github.com/go-fed/httpsig"
	"github.com/sirupsen/logrus"
	nethttp "net/http"
	"sync"
)

var (
	digestAlgo = httpsig.DigestSha256
	algoPrefs  = []httpsig.Algorithm{httpsig.RSA_SHA256}

	getHeaders  = []string{httpsig.RequestTarget, "host", "date"}
	postHeaders = []string{httpsig.RequestTarget, "host", "date", "digest"}
)

// Transport handled signing outgoing requests to federated instances
type Transport struct {
	client        pub.HttpClient
	clock         pub.Clock
	sigTransport  *pub.HttpSigTransport
	getSigner     httpsig.Signer
	getSignerLock *sync.Mutex

	keyID   string
	privKey crypto.PrivateKey
}

// New creates a new Transport module
func New(clock pub.Clock, pubKeyID string, privkey crypto.PrivateKey) (*Transport, error) {
	l := logrus.WithField("func", "New")

	getSigner, _, err := httpsig.NewSigner(algoPrefs, digestAlgo, getHeaders, httpsig.Signature, 120)
	if err != nil {
		l.Debugf("can't make get signer: %s", err.Error())
		return nil, fmt.Errorf("can't make get signer: %s", err)
	}

	postSigner, _, err := httpsig.NewSigner(algoPrefs, digestAlgo, postHeaders, httpsig.Signature, 120)
	if err != nil {
		l.Debugf("can't make post signer: %s", err.Error())
		return nil, fmt.Errorf("can't make post signer: %s", err)
	}

	httpClient := &nethttp.Client{
		Transport: &http.Transport{},
	}
	sigTransport := pub.NewHttpSigTransport(httpClient, http.GetUserAgent(), clock, getSigner, postSigner, pubKeyID, privkey)

	return &Transport{
		client:        httpClient,
		clock:         clock,
		sigTransport:  sigTransport,
		getSigner:     getSigner,
		getSignerLock: &sync.Mutex{},

		keyID:   pubKeyID,
		privKey: privkey,
	}, nil
}
