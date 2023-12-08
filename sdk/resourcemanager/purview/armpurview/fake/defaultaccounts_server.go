//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
	"net/http"
	"net/url"
)

// DefaultAccountsServer is a fake server for instances of the armpurview.DefaultAccountsClient type.
type DefaultAccountsServer struct {
	// Get is the fake for method DefaultAccountsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scopeTenantID string, scopeType armpurview.ScopeType, options *armpurview.DefaultAccountsClientGetOptions) (resp azfake.Responder[armpurview.DefaultAccountsClientGetResponse], errResp azfake.ErrorResponder)

	// Remove is the fake for method DefaultAccountsClient.Remove
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Remove func(ctx context.Context, scopeTenantID string, scopeType armpurview.ScopeType, options *armpurview.DefaultAccountsClientRemoveOptions) (resp azfake.Responder[armpurview.DefaultAccountsClientRemoveResponse], errResp azfake.ErrorResponder)

	// Set is the fake for method DefaultAccountsClient.Set
	// HTTP status codes to indicate success: http.StatusOK
	Set func(ctx context.Context, defaultAccountPayload armpurview.DefaultAccountPayload, options *armpurview.DefaultAccountsClientSetOptions) (resp azfake.Responder[armpurview.DefaultAccountsClientSetResponse], errResp azfake.ErrorResponder)
}

// NewDefaultAccountsServerTransport creates a new instance of DefaultAccountsServerTransport with the provided implementation.
// The returned DefaultAccountsServerTransport instance is connected to an instance of armpurview.DefaultAccountsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDefaultAccountsServerTransport(srv *DefaultAccountsServer) *DefaultAccountsServerTransport {
	return &DefaultAccountsServerTransport{srv: srv}
}

// DefaultAccountsServerTransport connects instances of armpurview.DefaultAccountsClient to instances of DefaultAccountsServer.
// Don't use this type directly, use NewDefaultAccountsServerTransport instead.
type DefaultAccountsServerTransport struct {
	srv *DefaultAccountsServer
}

// Do implements the policy.Transporter interface for DefaultAccountsServerTransport.
func (d *DefaultAccountsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DefaultAccountsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DefaultAccountsClient.Remove":
		resp, err = d.dispatchRemove(req)
	case "DefaultAccountsClient.Set":
		resp, err = d.dispatchSet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DefaultAccountsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	qp := req.URL.Query()
	scopeTenantIDParam, err := url.QueryUnescape(qp.Get("scopeTenantId"))
	if err != nil {
		return nil, err
	}
	scopeTypeParam, err := parseWithCast(qp.Get("scopeType"), func(v string) (armpurview.ScopeType, error) {
		p, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpurview.ScopeType(p), nil
	})
	if err != nil {
		return nil, err
	}
	scopeUnescaped, err := url.QueryUnescape(qp.Get("scope"))
	if err != nil {
		return nil, err
	}
	scopeParam := getOptional(scopeUnescaped)
	var options *armpurview.DefaultAccountsClientGetOptions
	if scopeParam != nil {
		options = &armpurview.DefaultAccountsClientGetOptions{
			Scope: scopeParam,
		}
	}
	respr, errRespr := d.srv.Get(req.Context(), scopeTenantIDParam, scopeTypeParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DefaultAccountPayload, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefaultAccountsServerTransport) dispatchRemove(req *http.Request) (*http.Response, error) {
	if d.srv.Remove == nil {
		return nil, &nonRetriableError{errors.New("fake for method Remove not implemented")}
	}
	qp := req.URL.Query()
	scopeTenantIDParam, err := url.QueryUnescape(qp.Get("scopeTenantId"))
	if err != nil {
		return nil, err
	}
	scopeTypeParam, err := parseWithCast(qp.Get("scopeType"), func(v string) (armpurview.ScopeType, error) {
		p, unescapeErr := url.QueryUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armpurview.ScopeType(p), nil
	})
	if err != nil {
		return nil, err
	}
	scopeUnescaped, err := url.QueryUnescape(qp.Get("scope"))
	if err != nil {
		return nil, err
	}
	scopeParam := getOptional(scopeUnescaped)
	var options *armpurview.DefaultAccountsClientRemoveOptions
	if scopeParam != nil {
		options = &armpurview.DefaultAccountsClientRemoveOptions{
			Scope: scopeParam,
		}
	}
	respr, errRespr := d.srv.Remove(req.Context(), scopeTenantIDParam, scopeTypeParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefaultAccountsServerTransport) dispatchSet(req *http.Request) (*http.Response, error) {
	if d.srv.Set == nil {
		return nil, &nonRetriableError{errors.New("fake for method Set not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[armpurview.DefaultAccountPayload](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Set(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DefaultAccountPayload, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}