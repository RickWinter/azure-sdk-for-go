// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
	"net/http"
	"net/url"
	"regexp"
)

// StandbyVirtualMachinesServer is a fake server for instances of the armstandbypool.StandbyVirtualMachinesClient type.
type StandbyVirtualMachinesServer struct {
	// Get is the fake for method StandbyVirtualMachinesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, standbyVirtualMachinePoolName string, standbyVirtualMachineName string, options *armstandbypool.StandbyVirtualMachinesClientGetOptions) (resp azfake.Responder[armstandbypool.StandbyVirtualMachinesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByStandbyVirtualMachinePoolResourcePager is the fake for method StandbyVirtualMachinesClient.NewListByStandbyVirtualMachinePoolResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByStandbyVirtualMachinePoolResourcePager func(resourceGroupName string, standbyVirtualMachinePoolName string, options *armstandbypool.StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceOptions) (resp azfake.PagerResponder[armstandbypool.StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse])
}

// NewStandbyVirtualMachinesServerTransport creates a new instance of StandbyVirtualMachinesServerTransport with the provided implementation.
// The returned StandbyVirtualMachinesServerTransport instance is connected to an instance of armstandbypool.StandbyVirtualMachinesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewStandbyVirtualMachinesServerTransport(srv *StandbyVirtualMachinesServer) *StandbyVirtualMachinesServerTransport {
	return &StandbyVirtualMachinesServerTransport{
		srv: srv,
		newListByStandbyVirtualMachinePoolResourcePager: newTracker[azfake.PagerResponder[armstandbypool.StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse]](),
	}
}

// StandbyVirtualMachinesServerTransport connects instances of armstandbypool.StandbyVirtualMachinesClient to instances of StandbyVirtualMachinesServer.
// Don't use this type directly, use NewStandbyVirtualMachinesServerTransport instead.
type StandbyVirtualMachinesServerTransport struct {
	srv                                             *StandbyVirtualMachinesServer
	newListByStandbyVirtualMachinePoolResourcePager *tracker[azfake.PagerResponder[armstandbypool.StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse]]
}

// Do implements the policy.Transporter interface for StandbyVirtualMachinesServerTransport.
func (s *StandbyVirtualMachinesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *StandbyVirtualMachinesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "StandbyVirtualMachinesClient.Get":
		resp, err = s.dispatchGet(req)
	case "StandbyVirtualMachinesClient.NewListByStandbyVirtualMachinePoolResourcePager":
		resp, err = s.dispatchNewListByStandbyVirtualMachinePoolResourcePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (s *StandbyVirtualMachinesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StandbyPool/standbyVirtualMachinePools/(?P<standbyVirtualMachinePoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/standbyVirtualMachines/(?P<standbyVirtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	standbyVirtualMachinePoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("standbyVirtualMachinePoolName")])
	if err != nil {
		return nil, err
	}
	standbyVirtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("standbyVirtualMachineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, standbyVirtualMachinePoolNameParam, standbyVirtualMachineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StandbyVirtualMachineResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StandbyVirtualMachinesServerTransport) dispatchNewListByStandbyVirtualMachinePoolResourcePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByStandbyVirtualMachinePoolResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByStandbyVirtualMachinePoolResourcePager not implemented")}
	}
	newListByStandbyVirtualMachinePoolResourcePager := s.newListByStandbyVirtualMachinePoolResourcePager.get(req)
	if newListByStandbyVirtualMachinePoolResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StandbyPool/standbyVirtualMachinePools/(?P<standbyVirtualMachinePoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/standbyVirtualMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		standbyVirtualMachinePoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("standbyVirtualMachinePoolName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByStandbyVirtualMachinePoolResourcePager(resourceGroupNameParam, standbyVirtualMachinePoolNameParam, nil)
		newListByStandbyVirtualMachinePoolResourcePager = &resp
		s.newListByStandbyVirtualMachinePoolResourcePager.add(req, newListByStandbyVirtualMachinePoolResourcePager)
		server.PagerResponderInjectNextLinks(newListByStandbyVirtualMachinePoolResourcePager, req, func(page *armstandbypool.StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByStandbyVirtualMachinePoolResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByStandbyVirtualMachinePoolResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByStandbyVirtualMachinePoolResourcePager) {
		s.newListByStandbyVirtualMachinePoolResourcePager.remove(req)
	}
	return resp, nil
}
