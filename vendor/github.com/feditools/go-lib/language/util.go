package language

func isEmptyYaml(b []byte) bool {
	switch string(b) {
	case "":
		return true
	case "---":
		return true
	case "---\n":
		return true
	default:
		return false
	}
}
