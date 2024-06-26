//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// SKUsServer is a fake server for instances of the armdevcenter.SKUsClient type.
type SKUsServer struct {
	// NewListBySubscriptionPager is the fake for method SKUsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armdevcenter.SKUsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armdevcenter.SKUsClientListBySubscriptionResponse])
}

// NewSKUsServerTransport creates a new instance of SKUsServerTransport with the provided implementation.
// The returned SKUsServerTransport instance is connected to an instance of armdevcenter.SKUsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSKUsServerTransport(srv *SKUsServer) *SKUsServerTransport {
	return &SKUsServerTransport{
		srv:                        srv,
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armdevcenter.SKUsClientListBySubscriptionResponse]](),
	}
}

// SKUsServerTransport connects instances of armdevcenter.SKUsClient to instances of SKUsServer.
// Don't use this type directly, use NewSKUsServerTransport instead.
type SKUsServerTransport struct {
	srv                        *SKUsServer
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armdevcenter.SKUsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for SKUsServerTransport.
func (s *SKUsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SKUsClient.NewListBySubscriptionPager":
		resp, err = s.dispatchNewListBySubscriptionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SKUsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := s.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/skus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armdevcenter.SKUsClientListBySubscriptionOptions
		if topParam != nil {
			options = &armdevcenter.SKUsClientListBySubscriptionOptions{
				Top: topParam,
			}
		}
		resp := s.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		s.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armdevcenter.SKUsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		s.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}
