package domain

import (
	"net/http"
	"net/mail"
)

type Customer struct {
	ID     string
	Email  string
	Name   string
	Active bool
}

func (c *Customer) Validate() error {

	fieldErrors := make(map[string]string)

	_, errInvalidEmail := mail.ParseAddress(c.Email)

	if errInvalidEmail != nil {
		fieldErrors["email"] = "invalid e-mail address"
	}

	if len(fieldErrors) > 0 {
		return NewCoreErrorWithFields(
			http.StatusBadRequest,
			"bad request",
			fieldErrors,
		)
	}

	return nil
}
