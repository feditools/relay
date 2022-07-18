package models

// NodeInfo is a federated nodeinfo object.
type NodeInfo struct {
	Links []Link `json:"links"`
}

// NodeInfo2 is a federated nodeinfo 2.0 object.
type NodeInfo2 struct {
	Software NodeInfo2Software `json:"software"`
}

// NodeInfo2Software is the software section of a nodeinfo 2.0 object.
type NodeInfo2Software struct {
	Name string `json:"name"`
}
