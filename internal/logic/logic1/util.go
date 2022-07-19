package logic1

import "fmt"

func genActorSelf(domain string) string {
	return fmt.Sprintf("https://%s/actor", domain)
}
