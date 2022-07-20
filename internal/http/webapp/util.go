package webapp

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
	"github.com/feditools/relay/web"
	"github.com/gorilla/sessions"
	"io/ioutil"
	nethttp "net/http"
)

// auth helpers

func (m *Module) authRequireLoggedIn(w nethttp.ResponseWriter, r *nethttp.Request) (*models.Account, bool) {
	us := r.Context().Value(ContextKeySession).(*sessions.Session) // nolint

	account, ok := r.Context().Value(ContextKeyAccount).(*models.Account)
	if !ok {
		// Save current page
		if r.URL.Query().Encode() == "" {
			us.Values[SessionKeyLoginRedirect] = r.URL.Path
		} else {
			us.Values[SessionKeyLoginRedirect] = r.URL.Path + "?" + r.URL.Query().Encode()
		}
		err := us.Save(r, w)
		if err != nil {
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return nil, true
		}

		// redirect to login
		nethttp.Redirect(w, r, path.AppLogin, nethttp.StatusFound)

		return nil, true
	}

	return account, false
}

// signature caching

func getSignature(filePath string) (string, error) {
	l := logger.WithField("func", "getSignature")

	file, err := web.Files.Open(filePath)
	if err != nil {
		l.Errorf("opening file: %s", err.Error())

		return "", err
	}

	// read it
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// hash it
	h := sha512.New384()
	_, err = h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("sha384-%s", signature), nil
}

func (m *Module) getSignatureCached(filePath string) (string, error) {
	if sig, ok := m.readCachedSignature(filePath); ok {
		return sig, nil
	}
	sig, err := getSignature(filePath)
	if err != nil {
		return "", err
	}
	m.writeCachedSignature(filePath, sig)

	return sig, nil
}

func (m *Module) readCachedSignature(filePath string) (string, bool) {
	m.sigCacheLock.RLock()
	val, ok := m.sigCache[filePath]
	m.sigCacheLock.RUnlock()

	return val, ok
}

func (m *Module) writeCachedSignature(filePath string, sig string) {
	m.sigCacheLock.Lock()
	m.sigCache[filePath] = sig
	m.sigCacheLock.Unlock()
}
