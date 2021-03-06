// Code generated by go-swagger; DO NOT EDIT.

package token_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new token manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for token manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateClient create client API
*/
func (a *Client) CreateClient(params *CreateClientParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateClient",
		Method:             "POST",
		PathPattern:        "/v1/oauth2/{user_id}/client",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClientOK), nil

}

/*
Token token API
*/
func (a *Client) Token(params *TokenParams) (*TokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Token",
		Method:             "POST",
		PathPattern:        "/v1/oauth2/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TokenOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
