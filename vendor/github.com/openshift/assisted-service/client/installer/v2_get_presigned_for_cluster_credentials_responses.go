// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2GetPresignedForClusterCredentialsReader is a Reader for the V2GetPresignedForClusterCredentials structure.
type V2GetPresignedForClusterCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2GetPresignedForClusterCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2GetPresignedForClusterCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2GetPresignedForClusterCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2GetPresignedForClusterCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2GetPresignedForClusterCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2GetPresignedForClusterCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2GetPresignedForClusterCredentialsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2GetPresignedForClusterCredentialsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2GetPresignedForClusterCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2GetPresignedForClusterCredentialsOK creates a V2GetPresignedForClusterCredentialsOK with default headers values
func NewV2GetPresignedForClusterCredentialsOK() *V2GetPresignedForClusterCredentialsOK {
	return &V2GetPresignedForClusterCredentialsOK{}
}

/*
V2GetPresignedForClusterCredentialsOK describes a response with status code 200, with default header values.

Success.
*/
type V2GetPresignedForClusterCredentialsOK struct {
	Payload *models.PresignedURL
}

// IsSuccess returns true when this v2 get presigned for cluster credentials o k response has a 2xx status code
func (o *V2GetPresignedForClusterCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 get presigned for cluster credentials o k response has a 3xx status code
func (o *V2GetPresignedForClusterCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get presigned for cluster credentials o k response has a 4xx status code
func (o *V2GetPresignedForClusterCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 get presigned for cluster credentials o k response has a 5xx status code
func (o *V2GetPresignedForClusterCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get presigned for cluster credentials o k response a status code equal to that given
func (o *V2GetPresignedForClusterCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *V2GetPresignedForClusterCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsOK  %+v", 200, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsOK) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsOK  %+v", 200, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsOK) GetPayload() *models.PresignedURL {
	return o.Payload
}

func (o *V2GetPresignedForClusterCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PresignedURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPresignedForClusterCredentialsBadRequest creates a V2GetPresignedForClusterCredentialsBadRequest with default headers values
func NewV2GetPresignedForClusterCredentialsBadRequest() *V2GetPresignedForClusterCredentialsBadRequest {
	return &V2GetPresignedForClusterCredentialsBadRequest{}
}

/*
V2GetPresignedForClusterCredentialsBadRequest describes a response with status code 400, with default header values.

Error.
*/
type V2GetPresignedForClusterCredentialsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get presigned for cluster credentials bad request response has a 2xx status code
func (o *V2GetPresignedForClusterCredentialsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get presigned for cluster credentials bad request response has a 3xx status code
func (o *V2GetPresignedForClusterCredentialsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get presigned for cluster credentials bad request response has a 4xx status code
func (o *V2GetPresignedForClusterCredentialsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get presigned for cluster credentials bad request response has a 5xx status code
func (o *V2GetPresignedForClusterCredentialsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get presigned for cluster credentials bad request response a status code equal to that given
func (o *V2GetPresignedForClusterCredentialsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *V2GetPresignedForClusterCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetPresignedForClusterCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPresignedForClusterCredentialsUnauthorized creates a V2GetPresignedForClusterCredentialsUnauthorized with default headers values
func NewV2GetPresignedForClusterCredentialsUnauthorized() *V2GetPresignedForClusterCredentialsUnauthorized {
	return &V2GetPresignedForClusterCredentialsUnauthorized{}
}

/*
V2GetPresignedForClusterCredentialsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2GetPresignedForClusterCredentialsUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 get presigned for cluster credentials unauthorized response has a 2xx status code
func (o *V2GetPresignedForClusterCredentialsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get presigned for cluster credentials unauthorized response has a 3xx status code
func (o *V2GetPresignedForClusterCredentialsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get presigned for cluster credentials unauthorized response has a 4xx status code
func (o *V2GetPresignedForClusterCredentialsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get presigned for cluster credentials unauthorized response has a 5xx status code
func (o *V2GetPresignedForClusterCredentialsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get presigned for cluster credentials unauthorized response a status code equal to that given
func (o *V2GetPresignedForClusterCredentialsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2GetPresignedForClusterCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2GetPresignedForClusterCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPresignedForClusterCredentialsForbidden creates a V2GetPresignedForClusterCredentialsForbidden with default headers values
func NewV2GetPresignedForClusterCredentialsForbidden() *V2GetPresignedForClusterCredentialsForbidden {
	return &V2GetPresignedForClusterCredentialsForbidden{}
}

/*
V2GetPresignedForClusterCredentialsForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2GetPresignedForClusterCredentialsForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 get presigned for cluster credentials forbidden response has a 2xx status code
func (o *V2GetPresignedForClusterCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get presigned for cluster credentials forbidden response has a 3xx status code
func (o *V2GetPresignedForClusterCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get presigned for cluster credentials forbidden response has a 4xx status code
func (o *V2GetPresignedForClusterCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get presigned for cluster credentials forbidden response has a 5xx status code
func (o *V2GetPresignedForClusterCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get presigned for cluster credentials forbidden response a status code equal to that given
func (o *V2GetPresignedForClusterCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2GetPresignedForClusterCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsForbidden) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2GetPresignedForClusterCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPresignedForClusterCredentialsNotFound creates a V2GetPresignedForClusterCredentialsNotFound with default headers values
func NewV2GetPresignedForClusterCredentialsNotFound() *V2GetPresignedForClusterCredentialsNotFound {
	return &V2GetPresignedForClusterCredentialsNotFound{}
}

/*
V2GetPresignedForClusterCredentialsNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2GetPresignedForClusterCredentialsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get presigned for cluster credentials not found response has a 2xx status code
func (o *V2GetPresignedForClusterCredentialsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get presigned for cluster credentials not found response has a 3xx status code
func (o *V2GetPresignedForClusterCredentialsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get presigned for cluster credentials not found response has a 4xx status code
func (o *V2GetPresignedForClusterCredentialsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get presigned for cluster credentials not found response has a 5xx status code
func (o *V2GetPresignedForClusterCredentialsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get presigned for cluster credentials not found response a status code equal to that given
func (o *V2GetPresignedForClusterCredentialsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2GetPresignedForClusterCredentialsNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsNotFound) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetPresignedForClusterCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPresignedForClusterCredentialsMethodNotAllowed creates a V2GetPresignedForClusterCredentialsMethodNotAllowed with default headers values
func NewV2GetPresignedForClusterCredentialsMethodNotAllowed() *V2GetPresignedForClusterCredentialsMethodNotAllowed {
	return &V2GetPresignedForClusterCredentialsMethodNotAllowed{}
}

/*
V2GetPresignedForClusterCredentialsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2GetPresignedForClusterCredentialsMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get presigned for cluster credentials method not allowed response has a 2xx status code
func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get presigned for cluster credentials method not allowed response has a 3xx status code
func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get presigned for cluster credentials method not allowed response has a 4xx status code
func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get presigned for cluster credentials method not allowed response has a 5xx status code
func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get presigned for cluster credentials method not allowed response a status code equal to that given
func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetPresignedForClusterCredentialsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPresignedForClusterCredentialsConflict creates a V2GetPresignedForClusterCredentialsConflict with default headers values
func NewV2GetPresignedForClusterCredentialsConflict() *V2GetPresignedForClusterCredentialsConflict {
	return &V2GetPresignedForClusterCredentialsConflict{}
}

/*
V2GetPresignedForClusterCredentialsConflict describes a response with status code 409, with default header values.

Error.
*/
type V2GetPresignedForClusterCredentialsConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get presigned for cluster credentials conflict response has a 2xx status code
func (o *V2GetPresignedForClusterCredentialsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get presigned for cluster credentials conflict response has a 3xx status code
func (o *V2GetPresignedForClusterCredentialsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get presigned for cluster credentials conflict response has a 4xx status code
func (o *V2GetPresignedForClusterCredentialsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get presigned for cluster credentials conflict response has a 5xx status code
func (o *V2GetPresignedForClusterCredentialsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get presigned for cluster credentials conflict response a status code equal to that given
func (o *V2GetPresignedForClusterCredentialsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *V2GetPresignedForClusterCredentialsConflict) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsConflict  %+v", 409, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsConflict) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsConflict  %+v", 409, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetPresignedForClusterCredentialsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetPresignedForClusterCredentialsInternalServerError creates a V2GetPresignedForClusterCredentialsInternalServerError with default headers values
func NewV2GetPresignedForClusterCredentialsInternalServerError() *V2GetPresignedForClusterCredentialsInternalServerError {
	return &V2GetPresignedForClusterCredentialsInternalServerError{}
}

/*
V2GetPresignedForClusterCredentialsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2GetPresignedForClusterCredentialsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get presigned for cluster credentials internal server error response has a 2xx status code
func (o *V2GetPresignedForClusterCredentialsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get presigned for cluster credentials internal server error response has a 3xx status code
func (o *V2GetPresignedForClusterCredentialsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get presigned for cluster credentials internal server error response has a 4xx status code
func (o *V2GetPresignedForClusterCredentialsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 get presigned for cluster credentials internal server error response has a 5xx status code
func (o *V2GetPresignedForClusterCredentialsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 get presigned for cluster credentials internal server error response a status code equal to that given
func (o *V2GetPresignedForClusterCredentialsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2GetPresignedForClusterCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/downloads/credentials-presigned][%d] v2GetPresignedForClusterCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2GetPresignedForClusterCredentialsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetPresignedForClusterCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
