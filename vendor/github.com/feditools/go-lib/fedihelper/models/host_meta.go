package models

type HostMeta struct {
	XMLNS string         `xml:"xmlns,attr"`
	Links []HostMetaLink `xml:"Link"`
}

type HostMetaLink struct {
	Rel      string `xml:"rel,attr"`
	Template string `xml:"template,attr"`
}
