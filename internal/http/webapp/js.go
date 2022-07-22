package webapp

import "fmt"

const jsOpenModal = `
var autoOpenModal = new bootstrap.Modal(document.getElementById('%s'), {})
autoOpenModal.toggle()`

func JSOpenModal(selector string) string {
	return fmt.Sprintf(jsOpenModal, selector)
}
