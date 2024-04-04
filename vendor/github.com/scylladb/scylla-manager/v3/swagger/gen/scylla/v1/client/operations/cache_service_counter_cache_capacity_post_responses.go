// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v1/models"
)

// CacheServiceCounterCacheCapacityPostReader is a Reader for the CacheServiceCounterCacheCapacityPost structure.
type CacheServiceCounterCacheCapacityPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceCounterCacheCapacityPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceCounterCacheCapacityPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceCounterCacheCapacityPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceCounterCacheCapacityPostOK creates a CacheServiceCounterCacheCapacityPostOK with default headers values
func NewCacheServiceCounterCacheCapacityPostOK() *CacheServiceCounterCacheCapacityPostOK {
	return &CacheServiceCounterCacheCapacityPostOK{}
}

/*
CacheServiceCounterCacheCapacityPostOK handles this case with default header values.

Success
*/
type CacheServiceCounterCacheCapacityPostOK struct {
}

func (o *CacheServiceCounterCacheCapacityPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCacheServiceCounterCacheCapacityPostDefault creates a CacheServiceCounterCacheCapacityPostDefault with default headers values
func NewCacheServiceCounterCacheCapacityPostDefault(code int) *CacheServiceCounterCacheCapacityPostDefault {
	return &CacheServiceCounterCacheCapacityPostDefault{
		_statusCode: code,
	}
}

/*
CacheServiceCounterCacheCapacityPostDefault handles this case with default header values.

internal server error
*/
type CacheServiceCounterCacheCapacityPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service counter cache capacity post default response
func (o *CacheServiceCounterCacheCapacityPostDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceCounterCacheCapacityPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceCounterCacheCapacityPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceCounterCacheCapacityPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}