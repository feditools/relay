package kv

import "strconv"

const (
	keyBase = "relay:"

	keyAccount            = keyBase + "acct:"
	keyAccountAccessToken = keyAccount + "at:"

	keyFedi         = keyBase + "fedi:"
	keyFediActor    = keyFedi + "actor:"
	keyFediHostMeta = keyFedi + "hm:"
	keyFediNodeInfo = keyFedi + "ni:"

	keyInstance      = keyBase + "instance:"
	keyInstanceOAuth = keyInstance + "oauth:"

	keySession = keyBase + "session:"
)

// KeyAccountAccessToken returns the kv key which holds a user's access token.
func KeyAccountAccessToken(i int64) string { return keyAccountAccessToken + strconv.FormatInt(i, 10) }

// KeyFediActor returns the kv key which holds cached actor.
func KeyFediActor(a string) string { return keyFediActor + a }

// KeyFediNodeInfo returns the kv key which holds cached nodeinfo.
func KeyFediNodeInfo(d string) string { return keyFediNodeInfo + d }

// KeyFediHostMeta returns the kv key which holds cached host meta.
func KeyFediHostMeta(d string) string { return keyFediHostMeta + d }

// KeyInstanceOAuth returns the kv key which holds an instance's oauth tokens.
func KeyInstanceOAuth(i int64) string { return keyInstanceOAuth + strconv.FormatInt(i, 10) }

// KeySession returns the base kv key prefix.
func KeySession() string { return keySession }
