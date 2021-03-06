// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package daemonset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostIpamIPReader is a Reader for the PostIpamIP structure.
type PostIpamIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIpamIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIpamIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostIpamIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIpamIPOK creates a PostIpamIPOK with default headers values
func NewPostIpamIPOK() *PostIpamIPOK {
	return &PostIpamIPOK{}
}

/* PostIpamIPOK describes a response with status code 200, with default header values.

Success
*/
type PostIpamIPOK struct {
}

func (o *PostIpamIPOK) Error() string {
	return fmt.Sprintf("[POST /ipam/ip][%d] postIpamIpOK ", 200)
}

func (o *PostIpamIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIpamIPInternalServerError creates a PostIpamIPInternalServerError with default headers values
func NewPostIpamIPInternalServerError() *PostIpamIPInternalServerError {
	return &PostIpamIPInternalServerError{}
}

/* PostIpamIPInternalServerError describes a response with status code 500, with default header values.

Allocation failure
*/
type PostIpamIPInternalServerError struct {
}

func (o *PostIpamIPInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ipam/ip][%d] postIpamIpInternalServerError ", 500)
}

func (o *PostIpamIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
