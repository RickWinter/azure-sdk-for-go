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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
	"net/http"
	"net/url"
	"regexp"
)

// SimplifiedSolutionsServer is a fake server for instances of the armselfhelp.SimplifiedSolutionsClient type.
type SimplifiedSolutionsServer struct {
	// BeginCreate is the fake for method SimplifiedSolutionsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, scope string, simplifiedSolutionsResourceName string, simplifiedSolutionsRequestBody armselfhelp.SimplifiedSolutionsResource, options *armselfhelp.SimplifiedSolutionsClientBeginCreateOptions) (resp azfake.PollerResponder[armselfhelp.SimplifiedSolutionsClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SimplifiedSolutionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, simplifiedSolutionsResourceName string, options *armselfhelp.SimplifiedSolutionsClientGetOptions) (resp azfake.Responder[armselfhelp.SimplifiedSolutionsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewSimplifiedSolutionsServerTransport creates a new instance of SimplifiedSolutionsServerTransport with the provided implementation.
// The returned SimplifiedSolutionsServerTransport instance is connected to an instance of armselfhelp.SimplifiedSolutionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSimplifiedSolutionsServerTransport(srv *SimplifiedSolutionsServer) *SimplifiedSolutionsServerTransport {
	return &SimplifiedSolutionsServerTransport{
		srv:         srv,
		beginCreate: newTracker[azfake.PollerResponder[armselfhelp.SimplifiedSolutionsClientCreateResponse]](),
	}
}

// SimplifiedSolutionsServerTransport connects instances of armselfhelp.SimplifiedSolutionsClient to instances of SimplifiedSolutionsServer.
// Don't use this type directly, use NewSimplifiedSolutionsServerTransport instead.
type SimplifiedSolutionsServerTransport struct {
	srv         *SimplifiedSolutionsServer
	beginCreate *tracker[azfake.PollerResponder[armselfhelp.SimplifiedSolutionsClientCreateResponse]]
}

// Do implements the policy.Transporter interface for SimplifiedSolutionsServerTransport.
func (s *SimplifiedSolutionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SimplifiedSolutionsClient.BeginCreate":
		resp, err = s.dispatchBeginCreate(req)
	case "SimplifiedSolutionsClient.Get":
		resp, err = s.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SimplifiedSolutionsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := s.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Help/simplifiedSolutions/(?P<simplifiedSolutionsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armselfhelp.SimplifiedSolutionsResource](req)
		if err != nil {
			return nil, err
		}
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		simplifiedSolutionsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("simplifiedSolutionsResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreate(req.Context(), scopeParam, simplifiedSolutionsResourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		s.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		s.beginCreate.remove(req)
	}

	return resp, nil
}

func (s *SimplifiedSolutionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Help/simplifiedSolutions/(?P<simplifiedSolutionsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	simplifiedSolutionsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("simplifiedSolutionsResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), scopeParam, simplifiedSolutionsResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SimplifiedSolutionsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
