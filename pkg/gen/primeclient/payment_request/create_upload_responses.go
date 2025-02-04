// Code generated by go-swagger; DO NOT EDIT.

package payment_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreateUploadReader is a Reader for the CreateUpload structure.
type CreateUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUploadCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUploadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUploadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUploadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUploadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateUploadUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUploadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUploadCreated creates a CreateUploadCreated with default headers values
func NewCreateUploadCreated() *CreateUploadCreated {
	return &CreateUploadCreated{}
}

/*
CreateUploadCreated describes a response with status code 201, with default header values.

Successfully created upload of digital file.
*/
type CreateUploadCreated struct {
	Payload *primemessages.UploadWithOmissions
}

// IsSuccess returns true when this create upload created response has a 2xx status code
func (o *CreateUploadCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create upload created response has a 3xx status code
func (o *CreateUploadCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload created response has a 4xx status code
func (o *CreateUploadCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create upload created response has a 5xx status code
func (o *CreateUploadCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload created response a status code equal to that given
func (o *CreateUploadCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create upload created response
func (o *CreateUploadCreated) Code() int {
	return 201
}

func (o *CreateUploadCreated) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadCreated  %+v", 201, o.Payload)
}

func (o *CreateUploadCreated) String() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadCreated  %+v", 201, o.Payload)
}

func (o *CreateUploadCreated) GetPayload() *primemessages.UploadWithOmissions {
	return o.Payload
}

func (o *CreateUploadCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.UploadWithOmissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadBadRequest creates a CreateUploadBadRequest with default headers values
func NewCreateUploadBadRequest() *CreateUploadBadRequest {
	return &CreateUploadBadRequest{}
}

/*
CreateUploadBadRequest describes a response with status code 400, with default header values.

The request payload is invalid.
*/
type CreateUploadBadRequest struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this create upload bad request response has a 2xx status code
func (o *CreateUploadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload bad request response has a 3xx status code
func (o *CreateUploadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload bad request response has a 4xx status code
func (o *CreateUploadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload bad request response has a 5xx status code
func (o *CreateUploadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload bad request response a status code equal to that given
func (o *CreateUploadBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create upload bad request response
func (o *CreateUploadBadRequest) Code() int {
	return 400
}

func (o *CreateUploadBadRequest) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUploadBadRequest) String() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUploadBadRequest) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateUploadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadUnauthorized creates a CreateUploadUnauthorized with default headers values
func NewCreateUploadUnauthorized() *CreateUploadUnauthorized {
	return &CreateUploadUnauthorized{}
}

/*
CreateUploadUnauthorized describes a response with status code 401, with default header values.

The request was denied.
*/
type CreateUploadUnauthorized struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this create upload unauthorized response has a 2xx status code
func (o *CreateUploadUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload unauthorized response has a 3xx status code
func (o *CreateUploadUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload unauthorized response has a 4xx status code
func (o *CreateUploadUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload unauthorized response has a 5xx status code
func (o *CreateUploadUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload unauthorized response a status code equal to that given
func (o *CreateUploadUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create upload unauthorized response
func (o *CreateUploadUnauthorized) Code() int {
	return 401
}

func (o *CreateUploadUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUploadUnauthorized) String() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUploadUnauthorized) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateUploadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadForbidden creates a CreateUploadForbidden with default headers values
func NewCreateUploadForbidden() *CreateUploadForbidden {
	return &CreateUploadForbidden{}
}

/*
CreateUploadForbidden describes a response with status code 403, with default header values.

The request was denied.
*/
type CreateUploadForbidden struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this create upload forbidden response has a 2xx status code
func (o *CreateUploadForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload forbidden response has a 3xx status code
func (o *CreateUploadForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload forbidden response has a 4xx status code
func (o *CreateUploadForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload forbidden response has a 5xx status code
func (o *CreateUploadForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload forbidden response a status code equal to that given
func (o *CreateUploadForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create upload forbidden response
func (o *CreateUploadForbidden) Code() int {
	return 403
}

func (o *CreateUploadForbidden) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadForbidden  %+v", 403, o.Payload)
}

func (o *CreateUploadForbidden) String() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadForbidden  %+v", 403, o.Payload)
}

func (o *CreateUploadForbidden) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateUploadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadNotFound creates a CreateUploadNotFound with default headers values
func NewCreateUploadNotFound() *CreateUploadNotFound {
	return &CreateUploadNotFound{}
}

/*
CreateUploadNotFound describes a response with status code 404, with default header values.

The requested resource wasn't found.
*/
type CreateUploadNotFound struct {
	Payload *primemessages.ClientError
}

// IsSuccess returns true when this create upload not found response has a 2xx status code
func (o *CreateUploadNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload not found response has a 3xx status code
func (o *CreateUploadNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload not found response has a 4xx status code
func (o *CreateUploadNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload not found response has a 5xx status code
func (o *CreateUploadNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload not found response a status code equal to that given
func (o *CreateUploadNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create upload not found response
func (o *CreateUploadNotFound) Code() int {
	return 404
}

func (o *CreateUploadNotFound) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadNotFound  %+v", 404, o.Payload)
}

func (o *CreateUploadNotFound) String() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadNotFound  %+v", 404, o.Payload)
}

func (o *CreateUploadNotFound) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateUploadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadUnprocessableEntity creates a CreateUploadUnprocessableEntity with default headers values
func NewCreateUploadUnprocessableEntity() *CreateUploadUnprocessableEntity {
	return &CreateUploadUnprocessableEntity{}
}

/*
CreateUploadUnprocessableEntity describes a response with status code 422, with default header values.

The request was unprocessable, likely due to bad input from the requester.
*/
type CreateUploadUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

// IsSuccess returns true when this create upload unprocessable entity response has a 2xx status code
func (o *CreateUploadUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload unprocessable entity response has a 3xx status code
func (o *CreateUploadUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload unprocessable entity response has a 4xx status code
func (o *CreateUploadUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create upload unprocessable entity response has a 5xx status code
func (o *CreateUploadUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create upload unprocessable entity response a status code equal to that given
func (o *CreateUploadUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create upload unprocessable entity response
func (o *CreateUploadUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateUploadUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateUploadUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateUploadUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *CreateUploadUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUploadInternalServerError creates a CreateUploadInternalServerError with default headers values
func NewCreateUploadInternalServerError() *CreateUploadInternalServerError {
	return &CreateUploadInternalServerError{}
}

/*
CreateUploadInternalServerError describes a response with status code 500, with default header values.

A server error occurred.
*/
type CreateUploadInternalServerError struct {
	Payload *primemessages.Error
}

// IsSuccess returns true when this create upload internal server error response has a 2xx status code
func (o *CreateUploadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create upload internal server error response has a 3xx status code
func (o *CreateUploadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create upload internal server error response has a 4xx status code
func (o *CreateUploadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create upload internal server error response has a 5xx status code
func (o *CreateUploadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create upload internal server error response a status code equal to that given
func (o *CreateUploadInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create upload internal server error response
func (o *CreateUploadInternalServerError) Code() int {
	return 500
}

func (o *CreateUploadInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUploadInternalServerError) String() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/uploads][%d] createUploadInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUploadInternalServerError) GetPayload() *primemessages.Error {
	return o.Payload
}

func (o *CreateUploadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
