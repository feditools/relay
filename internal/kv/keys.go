package kv

import "strconv"

const (
	keyBase = "relay:"

	keyAccount            = keyBase + "acct:"
	keyAccountAccessToken = keyAccount + "at:"

	keyFedi         = keyBase + "fedi:"
	keyFediHostMeta = keyFedi + "hm:"
	keyFediNodeInfo = keyFedi + "ni:"

	keyInstance      = keyBase + "instance:"
	keyInstanceOAuth = keyInstance + "oauth:"

	keySession = keyBase + "session:"
)

// KeyAccountAccessToken returns the kv key which holds a user's access token.
func KeyAccountAccessToken(i int64) string { return keyAccountAccessToken + strconv.FormatInt(i, 10) }

// KeyFediNodeInfo returns the kv key which holds cached nodeinfo.
func KeyFediNodeInfo(d string) string { return keyFediNodeInfo + d }

// KeyFediHostMeta returns the kv key which holds cached host meta.
func KeyFediHostMeta(d string) string { return keyFediHostMeta + d }

// KeyInstanceOAuth returns the kv key which holds an instance's oauth tokens.
func KeyInstanceOAuth(i int64) string { return keyInstanceOAuth + strconv.FormatInt(i, 10) }

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }
