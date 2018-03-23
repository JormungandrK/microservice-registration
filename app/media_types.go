// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "user": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-registration/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-registration
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// users media type (default view)
//
// Identifier: application/vnd.goa.user+json; view=default
type Users struct {
	// Status of user account
	Active bool `form:"active" json:"active" xml:"active"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
	// External id of user
	ExternalID string `form:"externalId" json:"externalId" xml:"externalId"`
	// Full name of user
	Fullname string `form:"fullname" json:"fullname" xml:"fullname"`
	// Unique user ID
	ID string `form:"id" json:"id" xml:"id"`
	// Roles of user
	Roles []string `form:"roles" json:"roles" xml:"roles"`
}

// Validate validates the Users media type instance.
func (mt *Users) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Fullname == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "fullname"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Roles == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "roles"))
	}
	if mt.ExternalID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "externalId"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`^([a-zA-Z0-9 ]{4,30})$`, mt.Fullname); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.fullname`, mt.Fullname, `^([a-zA-Z0-9 ]{4,30})$`))
	}
	return
}
