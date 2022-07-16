package models

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// Actor represents an activity pub actor
type Actor struct {
	Context           interface{} `json:"@context"`
	Endpoints         Endpoints   `json:"endpoints"`
	Followers         string      `json:"followers"`
	Following         string      `json:"following"`
	Inbox             string      `json:"inbox"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
	ID                string      `json:"id"`
	PublicKey         PublicKey   `json:"publicKey"`
	Summary           string      `json:"summary"`
	PreferredUsername string      `json:"preferredUsername"`
	URL               string      `json:"url"`
}

// Endpoints represents known activity pub endpoints
type Endpoints struct {
	SharedInbox string `json:"sharedInbox"`
}

// PublicKey represents an actor's public key
type PublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPEM string `json:"publicKeyPem"`
}

func (a *Actor) RSAPublicKey() (*rsa.PublicKey, error) {
	l := logger.WithField("func", "RSAPublicKey")

	actorPem, _ := pem.Decode([]byte(a.PublicKey.PublicKeyPEM))
	if actorPem == nil {
		msg := fmt.Sprintf("actor %s has invalid public key", a.URL)
		l.Debugf(msg)

		return nil, errors.New(msg)
	}

	l.Debugf("public key type: %s", actorPem.Type)
	if actorPem.Type != "PUBLIC KEY" {
		msg := fmt.Sprintf("actor %s has wrong key type", a.URL)
		l.Debugf(msg)

		return nil, errors.New(msg)
	}

	parsedKey, err := x509.ParsePKIXPublicKey(actorPem.Bytes)
	if err != nil {
		msg := fmt.Sprintf("can't parse public key for %s", a.URL)
		l.Debugf(msg)

		return nil, errors.New(msg)
	}
	pubKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		msg := fmt.Sprintf("can't cast public key for %s", a.URL)
		l.Debugf(msg)

		return nil, errors.New(msg)
	}

	return pubKey, nil
}
