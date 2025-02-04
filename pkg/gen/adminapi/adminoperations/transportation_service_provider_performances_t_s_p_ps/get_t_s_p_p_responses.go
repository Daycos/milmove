// Code generated by go-swagger; DO NOT EDIT.

package transportation_service_provider_performances_t_s_p_ps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/adminmessages"
)

// GetTSPPOKCode is the HTTP code returned for type GetTSPPOK
const GetTSPPOKCode int = 200

/*
GetTSPPOK success

swagger:response getTSPPOK
*/
type GetTSPPOK struct {

	/*
	  In: Body
	*/
	Payload *adminmessages.TransportationServiceProviderPerformance `json:"body,omitempty"`
}

// NewGetTSPPOK creates GetTSPPOK with default headers values
func NewGetTSPPOK() *GetTSPPOK {

	return &GetTSPPOK{}
}

// WithPayload adds the payload to the get t s p p o k response
func (o *GetTSPPOK) WithPayload(payload *adminmessages.TransportationServiceProviderPerformance) *GetTSPPOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get t s p p o k response
func (o *GetTSPPOK) SetPayload(payload *adminmessages.TransportationServiceProviderPerformance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTSPPOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTSPPBadRequestCode is the HTTP code returned for type GetTSPPBadRequest
const GetTSPPBadRequestCode int = 400

/*
GetTSPPBadRequest invalid request

swagger:response getTSPPBadRequest
*/
type GetTSPPBadRequest struct {
}

// NewGetTSPPBadRequest creates GetTSPPBadRequest with default headers values
func NewGetTSPPBadRequest() *GetTSPPBadRequest {

	return &GetTSPPBadRequest{}
}

// WriteResponse to the client
func (o *GetTSPPBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetTSPPUnauthorizedCode is the HTTP code returned for type GetTSPPUnauthorized
const GetTSPPUnauthorizedCode int = 401

/*
GetTSPPUnauthorized request requires user authentication

swagger:response getTSPPUnauthorized
*/
type GetTSPPUnauthorized struct {
}

// NewGetTSPPUnauthorized creates GetTSPPUnauthorized with default headers values
func NewGetTSPPUnauthorized() *GetTSPPUnauthorized {

	return &GetTSPPUnauthorized{}
}

// WriteResponse to the client
func (o *GetTSPPUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetTSPPNotFoundCode is the HTTP code returned for type GetTSPPNotFound
const GetTSPPNotFoundCode int = 404

/*
GetTSPPNotFound Transportation Service Provider Performances (TSPPs) not found

swagger:response getTSPPNotFound
*/
type GetTSPPNotFound struct {
}

// NewGetTSPPNotFound creates GetTSPPNotFound with default headers values
func NewGetTSPPNotFound() *GetTSPPNotFound {

	return &GetTSPPNotFound{}
}

// WriteResponse to the client
func (o *GetTSPPNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTSPPInternalServerErrorCode is the HTTP code returned for type GetTSPPInternalServerError
const GetTSPPInternalServerErrorCode int = 500

/*
GetTSPPInternalServerError server error

swagger:response getTSPPInternalServerError
*/
type GetTSPPInternalServerError struct {
}

// NewGetTSPPInternalServerError creates GetTSPPInternalServerError with default headers values
func NewGetTSPPInternalServerError() *GetTSPPInternalServerError {

	return &GetTSPPInternalServerError{}
}

// WriteResponse to the client
func (o *GetTSPPInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
