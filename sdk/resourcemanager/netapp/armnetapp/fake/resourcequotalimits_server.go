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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
	"net/http"
	"net/url"
	"regexp"
)

// ResourceQuotaLimitsServer is a fake server for instances of the armnetapp.ResourceQuotaLimitsClient type.
type ResourceQuotaLimitsServer struct {
	// Get is the fake for method ResourceQuotaLimitsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, quotaLimitName string, options *armnetapp.ResourceQuotaLimitsClientGetOptions) (resp azfake.Responder[armnetapp.ResourceQuotaLimitsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ResourceQuotaLimitsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armnetapp.ResourceQuotaLimitsClientListOptions) (resp azfake.PagerResponder[armnetapp.ResourceQuotaLimitsClientListResponse])
}

// NewResourceQuotaLimitsServerTransport creates a new instance of ResourceQuotaLimitsServerTransport with the provided implementation.
// The returned ResourceQuotaLimitsServerTransport instance is connected to an instance of armnetapp.ResourceQuotaLimitsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResourceQuotaLimitsServerTransport(srv *ResourceQuotaLimitsServer) *ResourceQuotaLimitsServerTransport {
	return &ResourceQuotaLimitsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetapp.ResourceQuotaLimitsClientListResponse]](),
	}
}

// ResourceQuotaLimitsServerTransport connects instances of armnetapp.ResourceQuotaLimitsClient to instances of ResourceQuotaLimitsServer.
// Don't use this type directly, use NewResourceQuotaLimitsServerTransport instead.
type ResourceQuotaLimitsServerTransport struct {
	srv          *ResourceQuotaLimitsServer
	newListPager *tracker[azfake.PagerResponder[armnetapp.ResourceQuotaLimitsClientListResponse]]
}

// Do implements the policy.Transporter interface for ResourceQuotaLimitsServerTransport.
func (r *ResourceQuotaLimitsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ResourceQuotaLimitsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if resourceQuotaLimitsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = resourceQuotaLimitsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ResourceQuotaLimitsClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ResourceQuotaLimitsClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
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

func (r *ResourceQuotaLimitsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotaLimits/(?P<quotaLimitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	quotaLimitNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("quotaLimitName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), locationParam, quotaLimitNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionQuotaItem, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceQuotaLimitsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotaLimits`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(locationParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetapp.ResourceQuotaLimitsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ResourceQuotaLimitsServerTransport
var resourceQuotaLimitsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
