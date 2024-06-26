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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
	"net/http"
	"net/url"
	"regexp"
)

// DbSystemShapesServer is a fake server for instances of the armoracledatabase.DbSystemShapesClient type.
type DbSystemShapesServer struct {
	// Get is the fake for method DbSystemShapesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, dbsystemshapename string, options *armoracledatabase.DbSystemShapesClientGetOptions) (resp azfake.Responder[armoracledatabase.DbSystemShapesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByLocationPager is the fake for method DbSystemShapesClient.NewListByLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByLocationPager func(location string, options *armoracledatabase.DbSystemShapesClientListByLocationOptions) (resp azfake.PagerResponder[armoracledatabase.DbSystemShapesClientListByLocationResponse])
}

// NewDbSystemShapesServerTransport creates a new instance of DbSystemShapesServerTransport with the provided implementation.
// The returned DbSystemShapesServerTransport instance is connected to an instance of armoracledatabase.DbSystemShapesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDbSystemShapesServerTransport(srv *DbSystemShapesServer) *DbSystemShapesServerTransport {
	return &DbSystemShapesServerTransport{
		srv:                    srv,
		newListByLocationPager: newTracker[azfake.PagerResponder[armoracledatabase.DbSystemShapesClientListByLocationResponse]](),
	}
}

// DbSystemShapesServerTransport connects instances of armoracledatabase.DbSystemShapesClient to instances of DbSystemShapesServer.
// Don't use this type directly, use NewDbSystemShapesServerTransport instead.
type DbSystemShapesServerTransport struct {
	srv                    *DbSystemShapesServer
	newListByLocationPager *tracker[azfake.PagerResponder[armoracledatabase.DbSystemShapesClientListByLocationResponse]]
}

// Do implements the policy.Transporter interface for DbSystemShapesServerTransport.
func (d *DbSystemShapesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DbSystemShapesClient.Get":
		resp, err = d.dispatchGet(req)
	case "DbSystemShapesClient.NewListByLocationPager":
		resp, err = d.dispatchNewListByLocationPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DbSystemShapesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dbSystemShapes/(?P<dbsystemshapename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	dbsystemshapenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dbsystemshapename")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), locationParam, dbsystemshapenameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DbSystemShape, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DbSystemShapesServerTransport) dispatchNewListByLocationPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByLocationPager not implemented")}
	}
	newListByLocationPager := d.newListByLocationPager.get(req)
	if newListByLocationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dbSystemShapes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByLocationPager(locationParam, nil)
		newListByLocationPager = &resp
		d.newListByLocationPager.add(req, newListByLocationPager)
		server.PagerResponderInjectNextLinks(newListByLocationPager, req, func(page *armoracledatabase.DbSystemShapesClientListByLocationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByLocationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByLocationPager) {
		d.newListByLocationPager.remove(req)
	}
	return resp, nil
}
