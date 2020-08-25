// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/care-pet/go/cmd/server/models"
)

// FindSensorDataBySensorIDAndTimeRangeReader is a Reader for the FindSensorDataBySensorIDAndTimeRange structure.
type FindSensorDataBySensorIDAndTimeRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindSensorDataBySensorIDAndTimeRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindSensorDataBySensorIDAndTimeRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindSensorDataBySensorIDAndTimeRangeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindSensorDataBySensorIDAndTimeRangeOK creates a FindSensorDataBySensorIDAndTimeRangeOK with default headers values
func NewFindSensorDataBySensorIDAndTimeRangeOK() *FindSensorDataBySensorIDAndTimeRangeOK {
	return &FindSensorDataBySensorIDAndTimeRangeOK{}
}

/*FindSensorDataBySensorIDAndTimeRangeOK handles this case with default header values.

sensors response
*/
type FindSensorDataBySensorIDAndTimeRangeOK struct {
	Payload []*models.Measure
}

func (o *FindSensorDataBySensorIDAndTimeRangeOK) Error() string {
	return fmt.Sprintf("[GET /sensor/{id}/{p}/values][%d] findSensorDataBySensorIdAndTimeRangeOK  %+v", 200, o.Payload)
}

func (o *FindSensorDataBySensorIDAndTimeRangeOK) GetPayload() []*models.Measure {
	return o.Payload
}

func (o *FindSensorDataBySensorIDAndTimeRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindSensorDataBySensorIDAndTimeRangeDefault creates a FindSensorDataBySensorIDAndTimeRangeDefault with default headers values
func NewFindSensorDataBySensorIDAndTimeRangeDefault(code int) *FindSensorDataBySensorIDAndTimeRangeDefault {
	return &FindSensorDataBySensorIDAndTimeRangeDefault{
		_statusCode: code,
	}
}

/*FindSensorDataBySensorIDAndTimeRangeDefault handles this case with default header values.

error
*/
type FindSensorDataBySensorIDAndTimeRangeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find sensor data by sensor id and time range default response
func (o *FindSensorDataBySensorIDAndTimeRangeDefault) Code() int {
	return o._statusCode
}

func (o *FindSensorDataBySensorIDAndTimeRangeDefault) Error() string {
	return fmt.Sprintf("[GET /sensor/{id}/{p}/values][%d] find sensor data by sensor id and time range default  %+v", o._statusCode, o.Payload)
}

func (o *FindSensorDataBySensorIDAndTimeRangeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindSensorDataBySensorIDAndTimeRangeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}