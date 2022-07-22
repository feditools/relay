package webapp

import (
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/util"
	"net/http"
)

// AdminBlockListGetHandler serves the home page.
func (m *Module) AdminBlockListGetHandler(w http.ResponseWriter, r *http.Request) {
	m.displayAdminBlockList(w, r, displayAdminBlockListConfig{})
}

// AdminBlockListPostHandler serves the home page.
func (m *Module) AdminBlockListPostHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "AdminBlockListPostHandler")

	// get localizer
	localizer := r.Context().Value(ContextKeyLocalizer).(*language.Localizer)

	// get form data
	if err := r.ParseForm(); err != nil {
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

		return
	}

	domain := r.FormValue(FormDomain)
	obfuscatedDomain := r.FormValue(FormObfuscatedDomain)
	blockSubdomains := false
	if r.FormValue(FormSubdomain) != "" {
		blockSubdomains = true
	}

	l.Debugf("got domain: %s", domain)
	l.Debugf("got obfuscated domain: %s", obfuscatedDomain)
	l.Debugf("got block subdomains: %t", blockSubdomains)

	// validate forms
	formValidationFailed := false
	formDomainValidation := &libtemplate.FormValidation{
		Response: localizer.TextLooksGood().String(),
		Valid:    true,
	}
	if err := validate.Var(domain, "required,fqdn"); err != nil {
		l.Debugf("domain validation error: %s", err.Error())

		formDomainValidation = &libtemplate.FormValidation{
			Response: "Must be valid fqdn",
			Valid:    false,
		}

		formValidationFailed = true
	}

	var formObfuscatedDomainValidation *libtemplate.FormValidation
	if obfuscatedDomain != "" {
		formObfuscatedDomainValidation = &libtemplate.FormValidation{
			Response: localizer.TextLooksGood().String(),
			Valid:    true,
		}
	}
	if obfuscatedDomain != "" && !util.ValidateDomainObfuscation(domain, obfuscatedDomain) {
		formObfuscatedDomainValidation = &libtemplate.FormValidation{
			Response: "Obfuscated domain doesn't match domain",
			Valid:    false,
		}

		formValidationFailed = true
	}

	if formValidationFailed {
		m.displayAdminBlockList(w, r, displayAdminBlockListConfig{
			DisplayAddModal: true,

			FormDomainValue:                domain,
			FormDomainValidation:           formDomainValidation,
			FormObfuscatedDomainValue:      obfuscatedDomain,
			FormObfuscatedDomainValidation: formObfuscatedDomainValidation,
			FormBlockSubdomainsValue:       blockSubdomains,
		})

		return
	}

	// check for existing block
	blocked, err := m.logic.IsDomainBlocked(r.Context(), domain)
	if err != nil {
		l.Errorf("checking domain block: %s", err.Error())
		m.returnErrorPage(w, r, http.StatusInternalServerError, ErrorResponseDBError)

		return
	}
	if blocked {
		m.displayAdminBlockList(w, r, displayAdminBlockListConfig{
			DisplayAddModal: true,

			FormError: &libtemplate.Alert{
				Text: localizer.TextBlockExists(domain).String(),
			},

			FormDomainValue:                domain,
			FormDomainValidation:           formDomainValidation,
			FormObfuscatedDomainValue:      obfuscatedDomain,
			FormObfuscatedDomainValidation: formObfuscatedDomainValidation,
			FormBlockSubdomainsValue:       blockSubdomains,
		})
	}

	// create new block
	newBlock := &models.Block{
		Domain:           domain,
		ObfuscatedDomain: obfuscatedDomain,
		BlockSubdomains:  blockSubdomains,
	}
	err = m.db.CreateBlock(r.Context(), newBlock)
	if err != nil {
		l.Errorf("db create: %s", err.Error())
		m.returnErrorPage(w, r, http.StatusInternalServerError, ErrorResponseDBError)

		return
	}

	// enqueue block update
	err = m.runner.EnqueueProcessBlock(r.Context(), newBlock.ID)
	if err != nil {
		l.Errorf("enqueueing job: %s", err.Error())
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

		return
	}

	// display page
	m.displayAdminBlockList(w, r, displayAdminBlockListConfig{})
}

type displayAdminBlockListConfig struct {
	DisplayAddModal bool

	FormError *libtemplate.Alert

	FormDomainValue                string
	FormDomainValidation           *libtemplate.FormValidation
	FormObfuscatedDomainValue      string
	FormObfuscatedDomainValidation *libtemplate.FormValidation
	FormBlockSubdomainsValue       bool
}

func (m *Module) displayAdminBlockList(w http.ResponseWriter, r *http.Request, config displayAdminBlockListConfig) {
	l := logger.WithField("func", "displayAdminBlockList")

	// get localizer
	localizer := r.Context().Value(ContextKeyLocalizer).(*language.Localizer) //nolint

	// Init template variables
	tmplVars := &template.AdminBlockList{
		Common: template.Common{
			PageTitle: localizer.TextBlock(2).String(),
		},
	}
	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())

		return
	}

	// create form inputs
	tmplVars.FormDomainValue = libtemplate.FormInput{
		ID:          "formInputDomain",
		Type:        libtemplate.FormInputTypeText,
		Name:        FormDomain,
		Placeholder: "example.com",
		Label:       localizer.TextDomain(1),
		LabelClass:  "form-label",
		Value:       config.FormDomainValue,
		Disabled:    false,
		Required:    true,
		Validation:  config.FormDomainValidation,
	}
	tmplVars.FormObfuscatedDomainValue = libtemplate.FormInput{
		ID:          "formInputObfuscatedDomain",
		Type:        libtemplate.FormInputTypeText,
		Name:        FormObfuscatedDomain,
		Placeholder: "e*****e.com",
		Label:       localizer.TextObfuscatedDomain(1),
		LabelClass:  "form-label",
		Value:       config.FormObfuscatedDomainValue,
		Disabled:    false,
		Required:    false,
		Validation:  config.FormObfuscatedDomainValidation,
	}
	tmplVars.FormBlockSubdomainsValue = libtemplate.FormInput{
		ID:         "formInputSubdomain",
		Type:       libtemplate.FormInputTypeCheckbox,
		Name:       FormSubdomain,
		Label:      localizer.TextBlockSubdomain(2),
		LabelClass: "form-check-label",
		Checked:    config.FormBlockSubdomainsValue,
		Disabled:   false,
		Required:   false,
	}

	if config.DisplayAddModal {
		tmplVars.FooterExtraScript = JSOpenModal("addModal")
	}

	err = m.executeTemplate(w, template.AdminBlockListName, tmplVars)
	if err != nil {
		l.Errorf("could not render '%s' template: %s", template.AdminBlockListName, err.Error())
	}
}
