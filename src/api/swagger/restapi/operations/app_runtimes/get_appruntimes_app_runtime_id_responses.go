// Code generated by go-swagger; DO NOT EDIT.

package app_runtimes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/src/api/swagger/models"
)

// GetAppruntimesAppRuntimeIDOKCode is the HTTP code returned for type GetAppruntimesAppRuntimeIDOK
const GetAppruntimesAppRuntimeIDOKCode int = 200

/*GetAppruntimesAppRuntimeIDOK An app runtime

swagger:response getAppruntimesAppRuntimeIdOK
*/
type GetAppruntimesAppRuntimeIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.AppRuntime `json:"body,omitempty"`
}

// NewGetAppruntimesAppRuntimeIDOK creates GetAppruntimesAppRuntimeIDOK with default headers values
func NewGetAppruntimesAppRuntimeIDOK() *GetAppruntimesAppRuntimeIDOK {
	return &GetAppruntimesAppRuntimeIDOK{}
}

// WithPayload adds the payload to the get appruntimes app runtime Id o k response
func (o *GetAppruntimesAppRuntimeIDOK) WithPayload(payload *models.AppRuntime) *GetAppruntimesAppRuntimeIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get appruntimes app runtime Id o k response
func (o *GetAppruntimesAppRuntimeIDOK) SetPayload(payload *models.AppRuntime) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppruntimesAppRuntimeIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAppruntimesAppRuntimeIDNotFoundCode is the HTTP code returned for type GetAppruntimesAppRuntimeIDNotFound
const GetAppruntimesAppRuntimeIDNotFoundCode int = 404

/*GetAppruntimesAppRuntimeIDNotFound The app runtime does not exist.

swagger:response getAppruntimesAppRuntimeIdNotFound
*/
type GetAppruntimesAppRuntimeIDNotFound struct {
}

// NewGetAppruntimesAppRuntimeIDNotFound creates GetAppruntimesAppRuntimeIDNotFound with default headers values
func NewGetAppruntimesAppRuntimeIDNotFound() *GetAppruntimesAppRuntimeIDNotFound {
	return &GetAppruntimesAppRuntimeIDNotFound{}
}

// WriteResponse to the client
func (o *GetAppruntimesAppRuntimeIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// GetAppruntimesAppRuntimeIDInternalServerErrorCode is the HTTP code returned for type GetAppruntimesAppRuntimeIDInternalServerError
const GetAppruntimesAppRuntimeIDInternalServerErrorCode int = 500

/*GetAppruntimesAppRuntimeIDInternalServerError An unexpected error occured.

swagger:response getAppruntimesAppRuntimeIdInternalServerError
*/
type GetAppruntimesAppRuntimeIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAppruntimesAppRuntimeIDInternalServerError creates GetAppruntimesAppRuntimeIDInternalServerError with default headers values
func NewGetAppruntimesAppRuntimeIDInternalServerError() *GetAppruntimesAppRuntimeIDInternalServerError {
	return &GetAppruntimesAppRuntimeIDInternalServerError{}
}

// WithPayload adds the payload to the get appruntimes app runtime Id internal server error response
func (o *GetAppruntimesAppRuntimeIDInternalServerError) WithPayload(payload *models.Error) *GetAppruntimesAppRuntimeIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get appruntimes app runtime Id internal server error response
func (o *GetAppruntimesAppRuntimeIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppruntimesAppRuntimeIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
