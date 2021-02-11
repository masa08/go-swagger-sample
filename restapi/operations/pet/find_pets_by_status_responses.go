// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-swagger-example/models"
)

// FindPetsByStatusOKCode is the HTTP code returned for type FindPetsByStatusOK
const FindPetsByStatusOKCode int = 200

/*FindPetsByStatusOK successful operation

swagger:response findPetsByStatusOK
*/
type FindPetsByStatusOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Pet `json:"body,omitempty"`
}

// NewFindPetsByStatusOK creates FindPetsByStatusOK with default headers values
func NewFindPetsByStatusOK() *FindPetsByStatusOK {

	return &FindPetsByStatusOK{}
}

// WithPayload adds the payload to the find pets by status o k response
func (o *FindPetsByStatusOK) WithPayload(payload []*models.Pet) *FindPetsByStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pets by status o k response
func (o *FindPetsByStatusOK) SetPayload(payload []*models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPetsByStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Pet, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// FindPetsByStatusBadRequestCode is the HTTP code returned for type FindPetsByStatusBadRequest
const FindPetsByStatusBadRequestCode int = 400

/*FindPetsByStatusBadRequest Invalid status value

swagger:response findPetsByStatusBadRequest
*/
type FindPetsByStatusBadRequest struct {
}

// NewFindPetsByStatusBadRequest creates FindPetsByStatusBadRequest with default headers values
func NewFindPetsByStatusBadRequest() *FindPetsByStatusBadRequest {

	return &FindPetsByStatusBadRequest{}
}

// WriteResponse to the client
func (o *FindPetsByStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
