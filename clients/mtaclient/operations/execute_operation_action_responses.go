// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExecuteOperationActionReader is a Reader for the ExecuteOperationAction structure.
type ExecuteOperationActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteOperationActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewExecuteOperationActionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExecuteOperationActionAccepted creates a ExecuteOperationActionAccepted with default headers values
func NewExecuteOperationActionAccepted() *ExecuteOperationActionAccepted {
	return &ExecuteOperationActionAccepted{}
}

/*ExecuteOperationActionAccepted handles this case with default header values.

Accepted
*/
type ExecuteOperationActionAccepted struct {
	Location strfmt.URI
}

func (o *ExecuteOperationActionAccepted) Error() string {
	return fmt.Sprintf("[POST /operations/{operationId}][%d] executeOperationActionAccepted ", 202)
}

func (o *ExecuteOperationActionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location

	location, err := formats.Parse("uri", response.GetHeader("Location"))
	if err != nil {
		return errors.InvalidType("Location", "header", "strfmt.URI", response.GetHeader("Location"))
	}
	o.Location = *(location.(*strfmt.URI))

	return nil
}