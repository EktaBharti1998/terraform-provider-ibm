// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PostFloatingIpsReader is a Reader for the PostFloatingIps structure.
type PostFloatingIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFloatingIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostFloatingIpsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostFloatingIpsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostFloatingIpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPostFloatingIpsCreated creates a PostFloatingIpsCreated with default headers values
func NewPostFloatingIpsCreated() *PostFloatingIpsCreated {
	return &PostFloatingIpsCreated{}
}

/*PostFloatingIpsCreated handles this case with default header values.

dummy
*/
type PostFloatingIpsCreated struct {
	Payload *models.FloatingIP
}

func (o *PostFloatingIpsCreated) Error() string {
	return fmt.Sprintf("[POST /floating_ips][%d] postFloatingIpsCreated  %+v", 201, o.Payload)
}

func (o *PostFloatingIpsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FloatingIP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFloatingIpsBadRequest creates a PostFloatingIpsBadRequest with default headers values
func NewPostFloatingIpsBadRequest() *PostFloatingIpsBadRequest {
	return &PostFloatingIpsBadRequest{}
}

/*PostFloatingIpsBadRequest handles this case with default header values.

error
*/
type PostFloatingIpsBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostFloatingIpsBadRequest) Error() string {
	return fmt.Sprintf("[POST /floating_ips][%d] postFloatingIpsBadRequest  %+v", 400, o.Payload)
}

func (o *PostFloatingIpsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFloatingIpsDefault creates a PostFloatingIpsDefault with default headers values
func NewPostFloatingIpsDefault(code int) *PostFloatingIpsDefault {
	return &PostFloatingIpsDefault{
		_statusCode: code,
	}
}

/*PostFloatingIpsDefault handles this case with default header values.

unexpectederror
*/
type PostFloatingIpsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the post floating ips default response
func (o *PostFloatingIpsDefault) Code() int {
	return o._statusCode
}

func (o *PostFloatingIpsDefault) Error() string {
	return fmt.Sprintf("[POST /floating_ips][%d] PostFloatingIps default  %+v", o._statusCode, o.Payload)
}

func (o *PostFloatingIpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}