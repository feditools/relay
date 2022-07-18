package models

type HostMeta struct {
	XMLNS string `xml:"xmlns,attr"`
	Links []Link `xml:"Link"`
}

func (h *HostMeta) WebfingerURI() WebfingerURI {
	for _, link := range h.Links {
		if link.Rel == HostMetaWebFingerTemplateRel {
			return WebfingerURI(link.Template)
		}
	}
	return ""
}
