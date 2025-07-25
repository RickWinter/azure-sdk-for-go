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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkSecurityPerimeterConfigurationsServer is a fake server for instances of the armhybridcompute.NetworkSecurityPerimeterConfigurationsClient type.
type NetworkSecurityPerimeterConfigurationsServer struct {
	// GetByPrivateLinkScope is the fake for method NetworkSecurityPerimeterConfigurationsClient.GetByPrivateLinkScope
	// HTTP status codes to indicate success: http.StatusOK
	GetByPrivateLinkScope func(ctx context.Context, resourceGroupName string, scopeName string, perimeterName string, options *armhybridcompute.NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeOptions) (resp azfake.Responder[armhybridcompute.NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse], errResp azfake.ErrorResponder)

	// NewListByPrivateLinkScopePager is the fake for method NetworkSecurityPerimeterConfigurationsClient.NewListByPrivateLinkScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByPrivateLinkScopePager func(resourceGroupName string, scopeName string, options *armhybridcompute.NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeOptions) (resp azfake.PagerResponder[armhybridcompute.NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse])

	// BeginReconcileForPrivateLinkScope is the fake for method NetworkSecurityPerimeterConfigurationsClient.BeginReconcileForPrivateLinkScope
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReconcileForPrivateLinkScope func(ctx context.Context, resourceGroupName string, scopeName string, perimeterName string, options *armhybridcompute.NetworkSecurityPerimeterConfigurationsClientBeginReconcileForPrivateLinkScopeOptions) (resp azfake.PollerResponder[armhybridcompute.NetworkSecurityPerimeterConfigurationsClientReconcileForPrivateLinkScopeResponse], errResp azfake.ErrorResponder)
}

// NewNetworkSecurityPerimeterConfigurationsServerTransport creates a new instance of NetworkSecurityPerimeterConfigurationsServerTransport with the provided implementation.
// The returned NetworkSecurityPerimeterConfigurationsServerTransport instance is connected to an instance of armhybridcompute.NetworkSecurityPerimeterConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkSecurityPerimeterConfigurationsServerTransport(srv *NetworkSecurityPerimeterConfigurationsServer) *NetworkSecurityPerimeterConfigurationsServerTransport {
	return &NetworkSecurityPerimeterConfigurationsServerTransport{
		srv:                               srv,
		newListByPrivateLinkScopePager:    newTracker[azfake.PagerResponder[armhybridcompute.NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse]](),
		beginReconcileForPrivateLinkScope: newTracker[azfake.PollerResponder[armhybridcompute.NetworkSecurityPerimeterConfigurationsClientReconcileForPrivateLinkScopeResponse]](),
	}
}

// NetworkSecurityPerimeterConfigurationsServerTransport connects instances of armhybridcompute.NetworkSecurityPerimeterConfigurationsClient to instances of NetworkSecurityPerimeterConfigurationsServer.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationsServerTransport instead.
type NetworkSecurityPerimeterConfigurationsServerTransport struct {
	srv                               *NetworkSecurityPerimeterConfigurationsServer
	newListByPrivateLinkScopePager    *tracker[azfake.PagerResponder[armhybridcompute.NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse]]
	beginReconcileForPrivateLinkScope *tracker[azfake.PollerResponder[armhybridcompute.NetworkSecurityPerimeterConfigurationsClientReconcileForPrivateLinkScopeResponse]]
}

// Do implements the policy.Transporter interface for NetworkSecurityPerimeterConfigurationsServerTransport.
func (n *NetworkSecurityPerimeterConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if networkSecurityPerimeterConfigurationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = networkSecurityPerimeterConfigurationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "NetworkSecurityPerimeterConfigurationsClient.GetByPrivateLinkScope":
				res.resp, res.err = n.dispatchGetByPrivateLinkScope(req)
			case "NetworkSecurityPerimeterConfigurationsClient.NewListByPrivateLinkScopePager":
				res.resp, res.err = n.dispatchNewListByPrivateLinkScopePager(req)
			case "NetworkSecurityPerimeterConfigurationsClient.BeginReconcileForPrivateLinkScope":
				res.resp, res.err = n.dispatchBeginReconcileForPrivateLinkScope(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchGetByPrivateLinkScope(req *http.Request) (*http.Response, error) {
	if n.srv.GetByPrivateLinkScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByPrivateLinkScope not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/privateLinkScopes/(?P<scopeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<perimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	scopeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("scopeName")])
	if err != nil {
		return nil, err
	}
	perimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("perimeterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.GetByPrivateLinkScope(req.Context(), resourceGroupNameParam, scopeNameParam, perimeterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkSecurityPerimeterConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchNewListByPrivateLinkScopePager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByPrivateLinkScopePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByPrivateLinkScopePager not implemented")}
	}
	newListByPrivateLinkScopePager := n.newListByPrivateLinkScopePager.get(req)
	if newListByPrivateLinkScopePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/privateLinkScopes/(?P<scopeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		scopeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("scopeName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByPrivateLinkScopePager(resourceGroupNameParam, scopeNameParam, nil)
		newListByPrivateLinkScopePager = &resp
		n.newListByPrivateLinkScopePager.add(req, newListByPrivateLinkScopePager)
		server.PagerResponderInjectNextLinks(newListByPrivateLinkScopePager, req, func(page *armhybridcompute.NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByPrivateLinkScopePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByPrivateLinkScopePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByPrivateLinkScopePager) {
		n.newListByPrivateLinkScopePager.remove(req)
	}
	return resp, nil
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchBeginReconcileForPrivateLinkScope(req *http.Request) (*http.Response, error) {
	if n.srv.BeginReconcileForPrivateLinkScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReconcileForPrivateLinkScope not implemented")}
	}
	beginReconcileForPrivateLinkScope := n.beginReconcileForPrivateLinkScope.get(req)
	if beginReconcileForPrivateLinkScope == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/privateLinkScopes/(?P<scopeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<perimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reconcile`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		scopeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("scopeName")])
		if err != nil {
			return nil, err
		}
		perimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("perimeterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginReconcileForPrivateLinkScope(req.Context(), resourceGroupNameParam, scopeNameParam, perimeterNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReconcileForPrivateLinkScope = &respr
		n.beginReconcileForPrivateLinkScope.add(req, beginReconcileForPrivateLinkScope)
	}

	resp, err := server.PollerResponderNext(beginReconcileForPrivateLinkScope, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginReconcileForPrivateLinkScope.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReconcileForPrivateLinkScope) {
		n.beginReconcileForPrivateLinkScope.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to NetworkSecurityPerimeterConfigurationsServerTransport
var networkSecurityPerimeterConfigurationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
