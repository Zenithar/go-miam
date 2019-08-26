package models

import (
	"fmt"

	"github.com/dchest/uniuri"
)

// Application describes a resource provider in term of OIDC.
type Application struct {
	ID     string `json:"id"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

// NewApplication returns an application instance.
func NewApplication(label string) *Application {
	return &Application{
		ID:     uniuri.NewLen(32),
		Label:  label,
		Active: false,
	}
}

// -----------------------------------------------------------------------------

// URN returns the entity URN
func (app *Application) URN() string {
	return fmt.Sprintf("miam:v1::application:%s", app.ID)
}
