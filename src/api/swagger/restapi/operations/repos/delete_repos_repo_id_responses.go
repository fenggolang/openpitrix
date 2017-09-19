// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/src/api/swagger/models"
)

// DeleteReposRepoIDNoContentCode is the HTTP code returned for type DeleteReposRepoIDNoContent
const DeleteReposRepoIDNoContentCode int = 204

/*DeleteReposRepoIDNoContent Repo successfully deleted.

swagger:response deleteReposRepoIdNoContent
*/
type DeleteReposRepoIDNoContent struct {
}

// NewDeleteReposRepoIDNoContent creates DeleteReposRepoIDNoContent with default headers values
func NewDeleteReposRepoIDNoContent() *DeleteReposRepoIDNoContent {
	return &DeleteReposRepoIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteReposRepoIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// DeleteReposRepoIDNotFoundCode is the HTTP code returned for type DeleteReposRepoIDNotFound
const DeleteReposRepoIDNotFoundCode int = 404

/*DeleteReposRepoIDNotFound The repo does not exist.

swagger:response deleteReposRepoIdNotFound
*/
type DeleteReposRepoIDNotFound struct {
}

// NewDeleteReposRepoIDNotFound creates DeleteReposRepoIDNotFound with default headers values
func NewDeleteReposRepoIDNotFound() *DeleteReposRepoIDNotFound {
	return &DeleteReposRepoIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteReposRepoIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// DeleteReposRepoIDInternalServerErrorCode is the HTTP code returned for type DeleteReposRepoIDInternalServerError
const DeleteReposRepoIDInternalServerErrorCode int = 500

/*DeleteReposRepoIDInternalServerError An unexpected error occured.

swagger:response deleteReposRepoIdInternalServerError
*/
type DeleteReposRepoIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteReposRepoIDInternalServerError creates DeleteReposRepoIDInternalServerError with default headers values
func NewDeleteReposRepoIDInternalServerError() *DeleteReposRepoIDInternalServerError {
	return &DeleteReposRepoIDInternalServerError{}
}

// WithPayload adds the payload to the delete repos repo Id internal server error response
func (o *DeleteReposRepoIDInternalServerError) WithPayload(payload *models.Error) *DeleteReposRepoIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete repos repo Id internal server error response
func (o *DeleteReposRepoIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteReposRepoIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
