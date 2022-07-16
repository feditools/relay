package logic

import "fmt"

func genActorSelf(domain string) string {
	return fmt.Sprintf("https://%s/actor", domain)
}
