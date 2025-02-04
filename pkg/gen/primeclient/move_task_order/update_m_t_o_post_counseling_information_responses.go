// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMTOPostCounselingInformationReader is a Reader for the UpdateMTOPostCounselingInformation structure.
type UpdateMTOPostCounselingInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMTOPostCounselingInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMTOPostCounselingInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateMTOPostCounselingInformationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMTOPostCounselingInformationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMTOPostCounselingInformationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateMTOPostCounselingInformationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdateMTOPostCounselingInformationPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateMTOPostCounselingInformationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMTOPostCounselingInformationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateMTOPostCounselingInformationOK creates a UpdateMTOPostCounselingInformationOK with default headers values
func NewUpdateMTOPostCounselingInformationOK() *UpdateMTOPostCounselingInformationOK {
	return &UpdateMTOPostCounselingInformationOK{}
}

/*
UpdateMTOPostCounselingInformationOK describes a response with status code 200, with default header values.

Successfully updated move task order with post counseling information.
*/
type UpdateMTOPostCounselingInformationOK struct {
	Payload *primemessages.MoveTaskOrder
}

// IsSuccess returns true when this update m t o post counseling information o k response has a 2xx status code
func (o *UpdateMTOPostCounselingInformationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update m t o post counseling information o k response has a 3xx status code
func (o *UpdateMTOPostCounselingInformationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m t o post counseling information o k response has a 4xx status code
func (o *UpdateMTOPostCounselingInformationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update m t o post counseling information o k response has a 5xx status code
func (o *UpdateMTOPostCounselingInformationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update m t o post counseling information o k response a status code equal to that given
func (o *UpdateMTOPostCounselingInformationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update m t o post counseling information o k response
func (o *UpdateMTOPostCounselingInformationOK) Code() int {
	return 200
}

func (o *UpdateMTOPostCounselingInformationOK) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationOK  %+v", 200, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationOK) String() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationOK  %+v", 200, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationOK) GetPayload() *primemessages.MoveTaskOrder {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.MoveTaskOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationUnauthorized creates a UpdateMTOPostCounselingInformationUnauthorized with default headers values
func NewUpdateMTOPostCounselingInformationUnauthorized() *UpdateMTOPostCounselingInformationUnauthorized {
	return &UpdateMTOPostCounselingInformationUnauthorized{}
}

/*
UpdateMTOPostCounselingInformationUnauthorized describes a response with status code 401, with default header values.

The request was denied.
*/
type UpdateMTOPostCounselingInformationUnauthorized struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this update m t o post counseling information unauthorized response has a 2xx status code
func (o *UpdateMTOPostCounselingInformationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m t o post counseling information unauthorized response has a 3xx status code
func (o *UpdateMTOPostCounselingInformationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m t o post counseling information unauthorized response has a 4xx status code
func (o *UpdateMTOPostCounselingInformationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update m t o post counseling information unauthorized response has a 5xx status code
func (o *UpdateMTOPostCounselingInformationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update m t o post counseling information unauthorized response a status code equal to that given
func (o *UpdateMTOPostCounselingInformationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update m t o post counseling information unauthorized response
func (o *UpdateMTOPostCounselingInformationUnauthorized) Code() int {
	return 401
}

func (o *UpdateMTOPostCounselingInformationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationUnauthorized) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationForbidden creates a UpdateMTOPostCounselingInformationForbidden with default headers values
func NewUpdateMTOPostCounselingInformationForbidden() *UpdateMTOPostCounselingInformationForbidden {
	return &UpdateMTOPostCounselingInformationForbidden{}
}

/*
UpdateMTOPostCounselingInformationForbidden describes a response with status code 403, with default header values.

The request was denied.
*/
type UpdateMTOPostCounselingInformationForbidden struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this update m t o post counseling information forbidden response has a 2xx status code
func (o *UpdateMTOPostCounselingInformationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m t o post counseling information forbidden response has a 3xx status code
func (o *UpdateMTOPostCounselingInformationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m t o post counseling information forbidden response has a 4xx status code
func (o *UpdateMTOPostCounselingInformationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update m t o post counseling information forbidden response has a 5xx status code
func (o *UpdateMTOPostCounselingInformationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update m t o post counseling information forbidden response a status code equal to that given
func (o *UpdateMTOPostCounselingInformationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update m t o post counseling information forbidden response
func (o *UpdateMTOPostCounselingInformationForbidden) Code() int {
	return 403
}

func (o *UpdateMTOPostCounselingInformationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationForbidden) String() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationForbidden) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationNotFound creates a UpdateMTOPostCounselingInformationNotFound with default headers values
func NewUpdateMTOPostCounselingInformationNotFound() *UpdateMTOPostCounselingInformationNotFound {
	return &UpdateMTOPostCounselingInformationNotFound{}
}

/*
UpdateMTOPostCounselingInformationNotFound describes a response with status code 404, with default header values.

The requested resource wasn't found.
*/
type UpdateMTOPostCounselingInformationNotFound struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this update m t o post counseling information not found response has a 2xx status code
func (o *UpdateMTOPostCounselingInformationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m t o post counseling information not found response has a 3xx status code
func (o *UpdateMTOPostCounselingInformationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m t o post counseling information not found response has a 4xx status code
func (o *UpdateMTOPostCounselingInformationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update m t o post counseling information not found response has a 5xx status code
func (o *UpdateMTOPostCounselingInformationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update m t o post counseling information not found response a status code equal to that given
func (o *UpdateMTOPostCounselingInformationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update m t o post counseling information not found response
func (o *UpdateMTOPostCounselingInformationNotFound) Code() int {
	return 404
}

func (o *UpdateMTOPostCounselingInformationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationNotFound) String() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationNotFound) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationConflict creates a UpdateMTOPostCounselingInformationConflict with default headers values
func NewUpdateMTOPostCounselingInformationConflict() *UpdateMTOPostCounselingInformationConflict {
	return &UpdateMTOPostCounselingInformationConflict{}
}

/*
UpdateMTOPostCounselingInformationConflict describes a response with status code 409, with default header values.

The request could not be processed because of conflict in the current state of the resource.
*/
type UpdateMTOPostCounselingInformationConflict struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this update m t o post counseling information conflict response has a 2xx status code
func (o *UpdateMTOPostCounselingInformationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m t o post counseling information conflict response has a 3xx status code
func (o *UpdateMTOPostCounselingInformationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m t o post counseling information conflict response has a 4xx status code
func (o *UpdateMTOPostCounselingInformationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update m t o post counseling information conflict response has a 5xx status code
func (o *UpdateMTOPostCounselingInformationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update m t o post counseling information conflict response a status code equal to that given
func (o *UpdateMTOPostCounselingInformationConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update m t o post counseling information conflict response
func (o *UpdateMTOPostCounselingInformationConflict) Code() int {
	return 409
}

func (o *UpdateMTOPostCounselingInformationConflict) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationConflict  %+v", 409, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationConflict) String() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationConflict  %+v", 409, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationConflict) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationPreconditionFailed creates a UpdateMTOPostCounselingInformationPreconditionFailed with default headers values
func NewUpdateMTOPostCounselingInformationPreconditionFailed() *UpdateMTOPostCounselingInformationPreconditionFailed {
	return &UpdateMTOPostCounselingInformationPreconditionFailed{}
}

/*
UpdateMTOPostCounselingInformationPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.
*/
type UpdateMTOPostCounselingInformationPreconditionFailed struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this update m t o post counseling information precondition failed response has a 2xx status code
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m t o post counseling information precondition failed response has a 3xx status code
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m t o post counseling information precondition failed response has a 4xx status code
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this update m t o post counseling information precondition failed response has a 5xx status code
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this update m t o post counseling information precondition failed response a status code equal to that given
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the update m t o post counseling information precondition failed response
func (o *UpdateMTOPostCounselingInformationPreconditionFailed) Code() int {
	return 412
}

func (o *UpdateMTOPostCounselingInformationPreconditionFailed) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationPreconditionFailed) String() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationPreconditionFailed) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationUnprocessableEntity creates a UpdateMTOPostCounselingInformationUnprocessableEntity with default headers values
func NewUpdateMTOPostCounselingInformationUnprocessableEntity() *UpdateMTOPostCounselingInformationUnprocessableEntity {
	return &UpdateMTOPostCounselingInformationUnprocessableEntity{}
}

/*
UpdateMTOPostCounselingInformationUnprocessableEntity describes a response with status code 422, with default header values.

The request was unprocessable, likely due to bad input from the requester.
*/
type UpdateMTOPostCounselingInformationUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

// IsSuccess returns true when this update m t o post counseling information unprocessable entity response has a 2xx status code
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m t o post counseling information unprocessable entity response has a 3xx status code
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m t o post counseling information unprocessable entity response has a 4xx status code
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update m t o post counseling information unprocessable entity response has a 5xx status code
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update m t o post counseling information unprocessable entity response a status code equal to that given
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update m t o post counseling information unprocessable entity response
func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) String() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOPostCounselingInformationInternalServerError creates a UpdateMTOPostCounselingInformationInternalServerError with default headers values
func NewUpdateMTOPostCounselingInformationInternalServerError() *UpdateMTOPostCounselingInformationInternalServerError {
	return &UpdateMTOPostCounselingInformationInternalServerError{}
}

/*
UpdateMTOPostCounselingInformationInternalServerError describes a response with status code 500, with default header values.

A server error occurred.
*/
type UpdateMTOPostCounselingInformationInternalServerError struct {
	Payload *primemessages.Error
}

// IsSuccess returns true when this update m t o post counseling information internal server error response has a 2xx status code
func (o *UpdateMTOPostCounselingInformationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m t o post counseling information internal server error response has a 3xx status code
func (o *UpdateMTOPostCounselingInformationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m t o post counseling information internal server error response has a 4xx status code
func (o *UpdateMTOPostCounselingInformationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update m t o post counseling information internal server error response has a 5xx status code
func (o *UpdateMTOPostCounselingInformationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update m t o post counseling information internal server error response a status code equal to that given
func (o *UpdateMTOPostCounselingInformationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update m t o post counseling information internal server error response
func (o *UpdateMTOPostCounselingInformationInternalServerError) Code() int {
	return 500
}

func (o *UpdateMTOPostCounselingInformationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info][%d] updateMTOPostCounselingInformationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMTOPostCounselingInformationInternalServerError) GetPayload() *primemessages.Error {
	return o.Payload
}

func (o *UpdateMTOPostCounselingInformationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
