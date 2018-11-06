// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/e154/smart-home/api/server_v1/stub/models"
)

// DeleteNodeByIDOKCode is the HTTP code returned for type DeleteNodeByIDOK
const DeleteNodeByIDOKCode int = 200

/*DeleteNodeByIDOK success

swagger:response deleteNodeByIdOK
*/
type DeleteNodeByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteNodeByIDOKBody `json:"body,omitempty"`
}

// NewDeleteNodeByIDOK creates DeleteNodeByIDOK with default headers values
func NewDeleteNodeByIDOK() *DeleteNodeByIDOK {
	return &DeleteNodeByIDOK{}
}

// WithPayload adds the payload to the delete node by Id o k response
func (o *DeleteNodeByIDOK) WithPayload(payload *models.DeleteNodeByIDOKBody) *DeleteNodeByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete node by Id o k response
func (o *DeleteNodeByIDOK) SetPayload(payload *models.DeleteNodeByIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNodeByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteNodeByIDDefault error

swagger:response deleteNodeByIdDefault
*/
type DeleteNodeByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewDeleteNodeByIDDefault creates DeleteNodeByIDDefault with default headers values
func NewDeleteNodeByIDDefault(code int) *DeleteNodeByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteNodeByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete node by Id default response
func (o *DeleteNodeByIDDefault) WithStatusCode(code int) *DeleteNodeByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete node by Id default response
func (o *DeleteNodeByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete node by Id default response
func (o *DeleteNodeByIDDefault) WithPayload(payload *models.ErrorModel) *DeleteNodeByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete node by Id default response
func (o *DeleteNodeByIDDefault) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNodeByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
