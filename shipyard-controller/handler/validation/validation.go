package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/keptn/keptn/shipyard-controller/config"
)

// ConfigurableProjectNameValidator can be used
// validate a Keptn project name
type ConfigurableProjectNameValidator struct {
	maxProjectNameLength int
}

// NewProjectValidator creates a new ConfigurableProjectNameValidator
func NewProjectValidator(env config.EnvConfig) *ConfigurableProjectNameValidator {
	return &ConfigurableProjectNameValidator{maxProjectNameLength: env.ProjectNameMaxLength}
}

// Validate performs the actual validation on field level
func (p ConfigurableProjectNameValidator) Validate(fl validator.FieldLevel) bool {
	if projectName, ok := fl.Field().Interface().(string); ok {
		return len(projectName) <= p.maxProjectNameLength
	}
	return true
}

// Tag returns the go tag the validator is bound to
func (p ConfigurableProjectNameValidator) Tag() string {
	return "projectname"
}
