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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// BookmarksServer is a fake server for instances of the armsecurityinsights.BookmarksClient type.
type BookmarksServer struct {
	// CreateOrUpdate is the fake for method BookmarksClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, bookmark armsecurityinsights.Bookmark, options *armsecurityinsights.BookmarksClientCreateOrUpdateOptions) (resp azfake.Responder[armsecurityinsights.BookmarksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method BookmarksClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, options *armsecurityinsights.BookmarksClientDeleteOptions) (resp azfake.Responder[armsecurityinsights.BookmarksClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BookmarksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, options *armsecurityinsights.BookmarksClientGetOptions) (resp azfake.Responder[armsecurityinsights.BookmarksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BookmarksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsecurityinsights.BookmarksClientListOptions) (resp azfake.PagerResponder[armsecurityinsights.BookmarksClientListResponse])
}

// NewBookmarksServerTransport creates a new instance of BookmarksServerTransport with the provided implementation.
// The returned BookmarksServerTransport instance is connected to an instance of armsecurityinsights.BookmarksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBookmarksServerTransport(srv *BookmarksServer) *BookmarksServerTransport {
	return &BookmarksServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurityinsights.BookmarksClientListResponse]](),
	}
}

// BookmarksServerTransport connects instances of armsecurityinsights.BookmarksClient to instances of BookmarksServer.
// Don't use this type directly, use NewBookmarksServerTransport instead.
type BookmarksServerTransport struct {
	srv          *BookmarksServer
	newListPager *tracker[azfake.PagerResponder[armsecurityinsights.BookmarksClientListResponse]]
}

// Do implements the policy.Transporter interface for BookmarksServerTransport.
func (b *BookmarksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BookmarksClient.CreateOrUpdate":
		resp, err = b.dispatchCreateOrUpdate(req)
	case "BookmarksClient.Delete":
		resp, err = b.dispatchDelete(req)
	case "BookmarksClient.Get":
		resp, err = b.dispatchGet(req)
	case "BookmarksClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BookmarksServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/bookmarks/(?P<bookmarkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.Bookmark](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	bookmarkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("bookmarkId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, bookmarkIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Bookmark, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BookmarksServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if b.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/bookmarks/(?P<bookmarkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	bookmarkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("bookmarkId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, bookmarkIDParam, nil)
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

func (b *BookmarksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/bookmarks/(?P<bookmarkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	bookmarkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("bookmarkId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, bookmarkIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Bookmark, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BookmarksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/bookmarks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, nil)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurityinsights.BookmarksClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}
