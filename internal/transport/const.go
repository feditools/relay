package transport

import "github.com/go-fed/httpsig"

const (
	rfc1123WithoutZone = "Mon, 02 Jan 2006 15:04:05"
)

var (
	digestAlgo = httpsig.DigestSha256
	algoPrefs  = []httpsig.Algorithm{httpsig.RSA_SHA256}

	getHeaders  = []string{httpsig.RequestTarget, "host", "date"}
	postHeaders = []string{httpsig.RequestTarget, "host", "date", "digest", "content-type"}
)
