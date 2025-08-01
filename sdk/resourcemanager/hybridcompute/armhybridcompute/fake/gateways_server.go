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

// GatewaysServer is a fake server for instances of the armhybridcompute.GatewaysClient type.
type GatewaysServer struct {
	// BeginCreateOrUpdate is the fake for method GatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, gatewayName string, parameters armhybridcompute.Gateway, options *armhybridcompute.GatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhybridcompute.GatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, gatewayName string, options *armhybridcompute.GatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridcompute.GatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gatewayName string, options *armhybridcompute.GatewaysClientGetOptions) (resp azfake.Responder[armhybridcompute.GatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method GatewaysClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armhybridcompute.GatewaysClientListByResourceGroupOptions) (resp azfake.PagerResponder[armhybridcompute.GatewaysClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method GatewaysClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armhybridcompute.GatewaysClientListBySubscriptionOptions) (resp azfake.PagerResponder[armhybridcompute.GatewaysClientListBySubscriptionResponse])

	// Update is the fake for method GatewaysClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, gatewayName string, parameters armhybridcompute.GatewayUpdate, options *armhybridcompute.GatewaysClientUpdateOptions) (resp azfake.Responder[armhybridcompute.GatewaysClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGatewaysServerTransport creates a new instance of GatewaysServerTransport with the provided implementation.
// The returned GatewaysServerTransport instance is connected to an instance of armhybridcompute.GatewaysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGatewaysServerTransport(srv *GatewaysServer) *GatewaysServerTransport {
	return &GatewaysServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armhybridcompute.GatewaysClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armhybridcompute.GatewaysClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armhybridcompute.GatewaysClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armhybridcompute.GatewaysClientListBySubscriptionResponse]](),
	}
}

// GatewaysServerTransport connects instances of armhybridcompute.GatewaysClient to instances of GatewaysServer.
// Don't use this type directly, use NewGatewaysServerTransport instead.
type GatewaysServerTransport struct {
	srv                         *GatewaysServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armhybridcompute.GatewaysClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armhybridcompute.GatewaysClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armhybridcompute.GatewaysClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armhybridcompute.GatewaysClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for GatewaysServerTransport.
func (g *GatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GatewaysServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if gatewaysServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = gatewaysServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GatewaysClient.BeginCreateOrUpdate":
				res.resp, res.err = g.dispatchBeginCreateOrUpdate(req)
			case "GatewaysClient.BeginDelete":
				res.resp, res.err = g.dispatchBeginDelete(req)
			case "GatewaysClient.Get":
				res.resp, res.err = g.dispatchGet(req)
			case "GatewaysClient.NewListByResourceGroupPager":
				res.resp, res.err = g.dispatchNewListByResourceGroupPager(req)
			case "GatewaysClient.NewListBySubscriptionPager":
				res.resp, res.err = g.dispatchNewListBySubscriptionPager(req)
			case "GatewaysClient.Update":
				res.resp, res.err = g.dispatchUpdate(req)
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

func (g *GatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := g.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/gateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridcompute.Gateway](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, gatewayNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		g.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		g.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		g.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (g *GatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := g.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/gateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), resourceGroupNameParam, gatewayNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		g.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		g.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		g.beginDelete.remove(req)
	}

	return resp, nil
}

func (g *GatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/gateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, gatewayNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Gateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GatewaysServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := g.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/gateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		g.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armhybridcompute.GatewaysClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		g.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (g *GatewaysServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := g.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/gateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := g.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		g.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armhybridcompute.GatewaysClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		g.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (g *GatewaysServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/gateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridcompute.GatewayUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Update(req.Context(), resourceGroupNameParam, gatewayNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Gateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to GatewaysServerTransport
var gatewaysServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
