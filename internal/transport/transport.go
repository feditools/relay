package transport

import (
	"crypto"
	"github.com/feditools/relay/internal/http"
	"github.com/go-fed/activity/pub"
	"github.com/go-fed/httpsig"
	nethttp "net/http"
	"sync"
	"time"
)

// Transport handled signing outgoing requests to federated instances
type Transport struct {
	client pub.HttpClient
	clock  pub.Clock

	keyID   string
	privKey crypto.PrivateKey

	signerExp  time.Time
	getSigner  httpsig.Signer
	postSigner httpsig.Signer
	signerMu   sync.Mutex
}

// New creates a new Transport module
func New(clock pub.Clock, pubKeyID string, privkey crypto.PrivateKey) (*Transport, error) {
	//l := logger.WithField("func", "New")

	httpClient := &nethttp.Client{
		Transport: &http.Transport{},
	}

	return &Transport{
		client: httpClient,
		clock:  clock,

		keyID:   pubKeyID,
		privKey: privkey,
	}, nil
}

func (t *Transport) doSign(do func()) {
	// Perform within mu safety
	t.signerMu.Lock()
	defer t.signerMu.Unlock()

	if now := time.Now(); now.After(t.signerExp) {
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
