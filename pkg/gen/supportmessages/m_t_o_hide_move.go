// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOHideMove describes the MTO ID and a description reason why the move was hidden.
//
// swagger:model MTOHideMove
type MTOHideMove struct {

	// Reason the move was selected to be hidden
	// Example: invalid name
	HideReason *string `json:"hideReason,omitempty"`

	// ID of the associated moveTaskOrder
	// Example: 1f2270c7-7166-40ae-981e-b200ebdf3054
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`
}

// Validate validates this m t o hide move
func (m *MTOHideMove) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOHideMove) validateMoveTaskOrderID(formats strfmt.Registry) error {
	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this m t o hide move based on context it is used
func (m *MTOHideMove) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MTOHideMove) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOHideMove) UnmarshalBinary(b []byte) error {
	var res MTOHideMove
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
