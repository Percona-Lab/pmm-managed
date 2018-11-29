// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// RemoveAgentReader is a Reader for the RemoveAgent structure.
type RemoveAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveAgentOK creates a RemoveAgentOK with default headers values
func NewRemoveAgentOK() *RemoveAgentOK {
	return &RemoveAgentOK{}
}

/*RemoveAgentOK handles this case with default header values.

(empty)
*/
type RemoveAgentOK struct {
	Payload models.InventoryRemoveAgentResponse
}

func (o *RemoveAgentOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/Remove][%d] removeAgentOK  %+v", 200, o.Payload)
}

func (o *RemoveAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
