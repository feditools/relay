package lib

import (
	"strings"
)

// SplitAccount splits a federated account into a username and domain.
func SplitAccount(act string) (username string, domain string, err error) {
	actFragments := strings.Split(strings.ToLower(act), "@")

	//revive:disable:add-constant
	switch len(actFragments) {
	case 2:
		return actFragments[0], actFragments[1], nil
	case 3:
		return actFragments[1], actFragments[2], nil
	default:
		return "", "", ErrInvalidAccountFormat
	} //revive:enable:add-constant
}
