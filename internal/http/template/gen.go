package template

import "github.com/feditools/relay/internal/path"

func genAppAdminBlockView(domain string) interface{} {
	return func(token string) string {
		return path.GenAppAdminBlockView(domain, token).String()
	}
}
func genAppAdminInstanceView(domain string) interface{} {
	return func(token string) string {
		return path.GenAppAdminInstanceView(domain, token).String()
	}
}
