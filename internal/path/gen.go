package path

import (
	"fmt"
	"net/url"
)

// GenActor returns a url for an actor
func GenActor(d string) string {
	return fmt.Sprintf("https://%s/%s", d, PartActor)
}

// GenFollowers returns a url for an actor's followers
func GenFollowers(d string) string {
	return fmt.Sprintf("https://%s/%s", d, PartFollowers)
}

// GenFollowing returns a url for actors following an actor
func GenFollowing(d string) string {
	return fmt.Sprintf("https://%s/%s", d, PartFollowing)
}

// GenInbox returns a url for an actor's inbox
func GenInbox(d string) string {
	return fmt.Sprintf("https://%s/%s", d, PartInbox)
}

// GenNodeinfo20 returns a url for an nodeinfo 2.0 url
func GenNodeinfo20(d string) string {
	return fmt.Sprintf("https://%s/%s/2.0", d, PartNodeinfo)
}

// GenPublicKey returns a url for an actor's public key
func GenPublicKey(d string) string {
	return fmt.Sprintf("%s#%s", GenActor(d), PartPublicKey)
}

// GenWellKnownNodeInfoURL returns a url for well known node info url
func GenWellKnownNodeInfoURL(d string) *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   d,
		Path:   APWellKnownNodeInfo,
	}
}
