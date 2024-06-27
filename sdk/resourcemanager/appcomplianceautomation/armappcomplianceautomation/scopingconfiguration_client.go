//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ScopingConfigurationClient contains the methods for the ScopingConfiguration group.
// Don't use this type directly, use NewScopingConfigurationClient() instead.
type ScopingConfigurationClient struct {
	internal *arm.Client
}

// NewScopingConfigurationClient creates a new instance of ScopingConfigurationClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScopingConfigurationClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ScopingConfigurationClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScopingConfigurationClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Get the AppComplianceAutomation scoping configuration of the specific report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - scopingConfigurationName - The scoping configuration of the specific report.
//   - properties - Parameters for the create or update operation, this is a singleton resource, so please make sure you're using
//     'default' as the name.
//   - options - ScopingConfigurationClientCreateOrUpdateOptions contains the optional parameters for the ScopingConfigurationClient.CreateOrUpdate
//     method.
func (client *ScopingConfigurationClient) CreateOrUpdate(ctx context.Context, reportName string, scopingConfigurationName string, properties ScopingConfigurationResource, options *ScopingConfigurationClientCreateOrUpdateOptions) (ScopingConfigurationClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ScopingConfigurationClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, reportName, scopingConfigurationName, properties, options)
	if err != nil {
		return ScopingConfigurationClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopingConfigurationClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ScopingConfigurationClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ScopingConfigurationClient) createOrUpdateCreateRequest(ctx context.Context, reportName string, scopingConfigurationName string, properties ScopingConfigurationResource, options *ScopingConfigurationClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/scopingConfigurations/{scopingConfigurationName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if scopingConfigurationName == "" {
		return nil, errors.New("parameter scopingConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopingConfigurationName}", url.PathEscape(scopingConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ScopingConfigurationClient) createOrUpdateHandleResponse(resp *http.Response) (ScopingConfigurationClientCreateOrUpdateResponse, error) {
	result := ScopingConfigurationClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopingConfigurationResource); err != nil {
		return ScopingConfigurationClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Clean the AppComplianceAutomation scoping configuration of the specific report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - scopingConfigurationName - The scoping configuration of the specific report.
//   - options - ScopingConfigurationClientDeleteOptions contains the optional parameters for the ScopingConfigurationClient.Delete
//     method.
func (client *ScopingConfigurationClient) Delete(ctx context.Context, reportName string, scopingConfigurationName string, options *ScopingConfigurationClientDeleteOptions) (ScopingConfigurationClientDeleteResponse, error) {
	var err error
	const operationName = "ScopingConfigurationClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, reportName, scopingConfigurationName, options)
	if err != nil {
		return ScopingConfigurationClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopingConfigurationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScopingConfigurationClientDeleteResponse{}, err
	}
	return ScopingConfigurationClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ScopingConfigurationClient) deleteCreateRequest(ctx context.Context, reportName string, scopingConfigurationName string, options *ScopingConfigurationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/scopingConfigurations/{scopingConfigurationName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if scopingConfigurationName == "" {
		return nil, errors.New("parameter scopingConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopingConfigurationName}", url.PathEscape(scopingConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the AppComplianceAutomation scoping configuration of the specific report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - scopingConfigurationName - The scoping configuration of the specific report.
//   - options - ScopingConfigurationClientGetOptions contains the optional parameters for the ScopingConfigurationClient.Get
//     method.
func (client *ScopingConfigurationClient) Get(ctx context.Context, reportName string, scopingConfigurationName string, options *ScopingConfigurationClientGetOptions) (ScopingConfigurationClientGetResponse, error) {
	var err error
	const operationName = "ScopingConfigurationClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, reportName, scopingConfigurationName, options)
	if err != nil {
		return ScopingConfigurationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopingConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScopingConfigurationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ScopingConfigurationClient) getCreateRequest(ctx context.Context, reportName string, scopingConfigurationName string, options *ScopingConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/scopingConfigurations/{scopingConfigurationName}"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	if scopingConfigurationName == "" {
		return nil, errors.New("parameter scopingConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopingConfigurationName}", url.PathEscape(scopingConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScopingConfigurationClient) getHandleResponse(resp *http.Response) (ScopingConfigurationClientGetResponse, error) {
	result := ScopingConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopingConfigurationResource); err != nil {
		return ScopingConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list format of the singleton scopingConfiguration for a specified report.
//
// Generated from API version 2024-06-27
//   - reportName - Report Name.
//   - options - ScopingConfigurationClientListOptions contains the optional parameters for the ScopingConfigurationClient.NewListPager
//     method.
func (client *ScopingConfigurationClient) NewListPager(reportName string, options *ScopingConfigurationClientListOptions) *runtime.Pager[ScopingConfigurationClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScopingConfigurationClientListResponse]{
		More: func(page ScopingConfigurationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScopingConfigurationClientListResponse) (ScopingConfigurationClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScopingConfigurationClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, reportName, options)
			}, nil)
			if err != nil {
				return ScopingConfigurationClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ScopingConfigurationClient) listCreateRequest(ctx context.Context, reportName string, options *ScopingConfigurationClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports/{reportName}/scopingConfigurations"
	if reportName == "" {
		return nil, errors.New("parameter reportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportName}", url.PathEscape(reportName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScopingConfigurationClient) listHandleResponse(resp *http.Response) (ScopingConfigurationClientListResponse, error) {
	result := ScopingConfigurationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopingConfigurationResourceListResult); err != nil {
		return ScopingConfigurationClientListResponse{}, err
	}
	return result, nil
}