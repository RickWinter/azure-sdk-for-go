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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"net/http"
	"net/url"
	"regexp"
)

// SubscriptionsServer is a fake server for instances of the armbilling.SubscriptionsClient type.
type SubscriptionsServer struct {
	// Get is the fake for method SubscriptionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, billingAccountName string, options *armbilling.SubscriptionsClientGetOptions) (resp azfake.Responder[armbilling.SubscriptionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByBillingAccountPager is the fake for method SubscriptionsClient.NewListByBillingAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingAccountPager func(billingAccountName string, options *armbilling.SubscriptionsClientListByBillingAccountOptions) (resp azfake.PagerResponder[armbilling.SubscriptionsClientListByBillingAccountResponse])

	// NewListByBillingProfilePager is the fake for method SubscriptionsClient.NewListByBillingProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingProfilePager func(billingAccountName string, billingProfileName string, options *armbilling.SubscriptionsClientListByBillingProfileOptions) (resp azfake.PagerResponder[armbilling.SubscriptionsClientListByBillingProfileResponse])

	// NewListByCustomerPager is the fake for method SubscriptionsClient.NewListByCustomerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCustomerPager func(billingAccountName string, customerName string, options *armbilling.SubscriptionsClientListByCustomerOptions) (resp azfake.PagerResponder[armbilling.SubscriptionsClientListByCustomerResponse])

	// NewListByInvoiceSectionPager is the fake for method SubscriptionsClient.NewListByInvoiceSectionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInvoiceSectionPager func(billingAccountName string, billingProfileName string, invoiceSectionName string, options *armbilling.SubscriptionsClientListByInvoiceSectionOptions) (resp azfake.PagerResponder[armbilling.SubscriptionsClientListByInvoiceSectionResponse])

	// BeginMove is the fake for method SubscriptionsClient.BeginMove
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginMove func(ctx context.Context, billingAccountName string, parameters armbilling.TransferBillingSubscriptionRequestProperties, options *armbilling.SubscriptionsClientBeginMoveOptions) (resp azfake.PollerResponder[armbilling.SubscriptionsClientMoveResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method SubscriptionsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, billingAccountName string, parameters armbilling.Subscription, options *armbilling.SubscriptionsClientUpdateOptions) (resp azfake.Responder[armbilling.SubscriptionsClientUpdateResponse], errResp azfake.ErrorResponder)

	// ValidateMove is the fake for method SubscriptionsClient.ValidateMove
	// HTTP status codes to indicate success: http.StatusOK
	ValidateMove func(ctx context.Context, billingAccountName string, parameters armbilling.TransferBillingSubscriptionRequestProperties, options *armbilling.SubscriptionsClientValidateMoveOptions) (resp azfake.Responder[armbilling.SubscriptionsClientValidateMoveResponse], errResp azfake.ErrorResponder)
}

// NewSubscriptionsServerTransport creates a new instance of SubscriptionsServerTransport with the provided implementation.
// The returned SubscriptionsServerTransport instance is connected to an instance of armbilling.SubscriptionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSubscriptionsServerTransport(srv *SubscriptionsServer) *SubscriptionsServerTransport {
	return &SubscriptionsServerTransport{
		srv:                          srv,
		newListByBillingAccountPager: newTracker[azfake.PagerResponder[armbilling.SubscriptionsClientListByBillingAccountResponse]](),
		newListByBillingProfilePager: newTracker[azfake.PagerResponder[armbilling.SubscriptionsClientListByBillingProfileResponse]](),
		newListByCustomerPager:       newTracker[azfake.PagerResponder[armbilling.SubscriptionsClientListByCustomerResponse]](),
		newListByInvoiceSectionPager: newTracker[azfake.PagerResponder[armbilling.SubscriptionsClientListByInvoiceSectionResponse]](),
		beginMove:                    newTracker[azfake.PollerResponder[armbilling.SubscriptionsClientMoveResponse]](),
	}
}

// SubscriptionsServerTransport connects instances of armbilling.SubscriptionsClient to instances of SubscriptionsServer.
// Don't use this type directly, use NewSubscriptionsServerTransport instead.
type SubscriptionsServerTransport struct {
	srv                          *SubscriptionsServer
	newListByBillingAccountPager *tracker[azfake.PagerResponder[armbilling.SubscriptionsClientListByBillingAccountResponse]]
	newListByBillingProfilePager *tracker[azfake.PagerResponder[armbilling.SubscriptionsClientListByBillingProfileResponse]]
	newListByCustomerPager       *tracker[azfake.PagerResponder[armbilling.SubscriptionsClientListByCustomerResponse]]
	newListByInvoiceSectionPager *tracker[azfake.PagerResponder[armbilling.SubscriptionsClientListByInvoiceSectionResponse]]
	beginMove                    *tracker[azfake.PollerResponder[armbilling.SubscriptionsClientMoveResponse]]
}

// Do implements the policy.Transporter interface for SubscriptionsServerTransport.
func (s *SubscriptionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SubscriptionsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SubscriptionsClient.NewListByBillingAccountPager":
		resp, err = s.dispatchNewListByBillingAccountPager(req)
	case "SubscriptionsClient.NewListByBillingProfilePager":
		resp, err = s.dispatchNewListByBillingProfilePager(req)
	case "SubscriptionsClient.NewListByCustomerPager":
		resp, err = s.dispatchNewListByCustomerPager(req)
	case "SubscriptionsClient.NewListByInvoiceSectionPager":
		resp, err = s.dispatchNewListByInvoiceSectionPager(req)
	case "SubscriptionsClient.BeginMove":
		resp, err = s.dispatchBeginMove(req)
	case "SubscriptionsClient.Update":
		resp, err = s.dispatchUpdate(req)
	case "SubscriptionsClient.ValidateMove":
		resp, err = s.dispatchValidateMove(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingSubscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), billingAccountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Subscription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchNewListByBillingAccountPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByBillingAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingAccountPager not implemented")}
	}
	newListByBillingAccountPager := s.newListByBillingAccountPager.get(req)
	if newListByBillingAccountPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingSubscriptions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByBillingAccountPager(billingAccountNameParam, nil)
		newListByBillingAccountPager = &resp
		s.newListByBillingAccountPager.add(req, newListByBillingAccountPager)
		server.PagerResponderInjectNextLinks(newListByBillingAccountPager, req, func(page *armbilling.SubscriptionsClientListByBillingAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByBillingAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingAccountPager) {
		s.newListByBillingAccountPager.remove(req)
	}
	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchNewListByBillingProfilePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByBillingProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingProfilePager not implemented")}
	}
	newListByBillingProfilePager := s.newListByBillingProfilePager.get(req)
	if newListByBillingProfilePager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingSubscriptions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByBillingProfilePager(billingAccountNameParam, billingProfileNameParam, nil)
		newListByBillingProfilePager = &resp
		s.newListByBillingProfilePager.add(req, newListByBillingProfilePager)
		server.PagerResponderInjectNextLinks(newListByBillingProfilePager, req, func(page *armbilling.SubscriptionsClientListByBillingProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByBillingProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingProfilePager) {
		s.newListByBillingProfilePager.remove(req)
	}
	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchNewListByCustomerPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByCustomerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCustomerPager not implemented")}
	}
	newListByCustomerPager := s.newListByCustomerPager.get(req)
	if newListByCustomerPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/customers/(?P<customerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingSubscriptions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		customerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByCustomerPager(billingAccountNameParam, customerNameParam, nil)
		newListByCustomerPager = &resp
		s.newListByCustomerPager.add(req, newListByCustomerPager)
		server.PagerResponderInjectNextLinks(newListByCustomerPager, req, func(page *armbilling.SubscriptionsClientListByCustomerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCustomerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByCustomerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCustomerPager) {
		s.newListByCustomerPager.remove(req)
	}
	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchNewListByInvoiceSectionPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByInvoiceSectionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInvoiceSectionPager not implemented")}
	}
	newListByInvoiceSectionPager := s.newListByInvoiceSectionPager.get(req)
	if newListByInvoiceSectionPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoiceSections/(?P<invoiceSectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingSubscriptions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		invoiceSectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("invoiceSectionName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByInvoiceSectionPager(billingAccountNameParam, billingProfileNameParam, invoiceSectionNameParam, nil)
		newListByInvoiceSectionPager = &resp
		s.newListByInvoiceSectionPager.add(req, newListByInvoiceSectionPager)
		server.PagerResponderInjectNextLinks(newListByInvoiceSectionPager, req, func(page *armbilling.SubscriptionsClientListByInvoiceSectionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInvoiceSectionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByInvoiceSectionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInvoiceSectionPager) {
		s.newListByInvoiceSectionPager.remove(req)
	}
	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchBeginMove(req *http.Request) (*http.Response, error) {
	if s.srv.BeginMove == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginMove not implemented")}
	}
	beginMove := s.beginMove.get(req)
	if beginMove == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingSubscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/move`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armbilling.TransferBillingSubscriptionRequestProperties](req)
		if err != nil {
			return nil, err
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginMove(req.Context(), billingAccountNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginMove = &respr
		s.beginMove.add(req, beginMove)
	}

	resp, err := server.PollerResponderNext(beginMove, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginMove.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginMove) {
		s.beginMove.remove(req)
	}

	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingSubscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armbilling.Subscription](req)
	if err != nil {
		return nil, err
	}
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Update(req.Context(), billingAccountNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Subscription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchValidateMove(req *http.Request) (*http.Response, error) {
	if s.srv.ValidateMove == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidateMove not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingSubscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validateMoveEligibility`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armbilling.TransferBillingSubscriptionRequestProperties](req)
	if err != nil {
		return nil, err
	}
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ValidateMove(req.Context(), billingAccountNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ValidateSubscriptionTransferEligibilityResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}