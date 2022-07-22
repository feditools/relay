package template

import (
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/relay/internal/models"
)

// AdminBlockListName is the name of the admin block list template.
const AdminBlockListName = "admin_block_list"

// AdminBlockList contains the variables for the admin block list template.
type AdminBlockList struct {
	Common

	FormError                 *libtemplate.Alert
	FormDomainValue           libtemplate.FormInput
	FormObfuscatedDomainValue libtemplate.FormInput
	FormBlockSubdomainsValue  libtemplate.FormInput

	Blocks []*models.Block
}
