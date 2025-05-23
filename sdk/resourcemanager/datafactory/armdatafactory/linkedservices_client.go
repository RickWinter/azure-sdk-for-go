// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

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

// LinkedServicesClient contains the methods for the LinkedServices group.
// Don't use this type directly, use NewLinkedServicesClient() instead.
type LinkedServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLinkedServicesClient creates a new instance of LinkedServicesClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLinkedServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkedServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LinkedServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - linkedServiceName - The linked service name.
//   - linkedService - Linked service resource definition.
//   - options - LinkedServicesClientCreateOrUpdateOptions contains the optional parameters for the LinkedServicesClient.CreateOrUpdate
//     method.
func (client *LinkedServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServicesClientCreateOrUpdateOptions) (LinkedServicesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "LinkedServicesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, linkedServiceName, linkedService, options)
	if err != nil {
		return LinkedServicesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedServicesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LinkedServicesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LinkedServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, linkedService LinkedServiceResource, options *LinkedServicesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, linkedService); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LinkedServicesClient) createOrUpdateHandleResponse(resp *http.Response) (LinkedServicesClientCreateOrUpdateResponse, error) {
	result := LinkedServicesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServiceResource); err != nil {
		return LinkedServicesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - linkedServiceName - The linked service name.
//   - options - LinkedServicesClientDeleteOptions contains the optional parameters for the LinkedServicesClient.Delete method.
func (client *LinkedServicesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, options *LinkedServicesClientDeleteOptions) (LinkedServicesClientDeleteResponse, error) {
	var err error
	const operationName = "LinkedServicesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, linkedServiceName, options)
	if err != nil {
		return LinkedServicesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedServicesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return LinkedServicesClientDeleteResponse{}, err
	}
	return LinkedServicesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkedServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, _ *LinkedServicesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - linkedServiceName - The linked service name.
//   - options - LinkedServicesClientGetOptions contains the optional parameters for the LinkedServicesClient.Get method.
func (client *LinkedServicesClient) Get(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, options *LinkedServicesClientGetOptions) (LinkedServicesClientGetResponse, error) {
	var err error
	const operationName = "LinkedServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, linkedServiceName, options)
	if err != nil {
		return LinkedServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(httpResp)
		return LinkedServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LinkedServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, options *LinkedServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedServicesClient) getHandleResponse(resp *http.Response) (LinkedServicesClientGetResponse, error) {
	result := LinkedServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServiceResource); err != nil {
		return LinkedServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - Lists linked services.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - options - LinkedServicesClientListByFactoryOptions contains the optional parameters for the LinkedServicesClient.NewListByFactoryPager
//     method.
func (client *LinkedServicesClient) NewListByFactoryPager(resourceGroupName string, factoryName string, options *LinkedServicesClientListByFactoryOptions) *runtime.Pager[LinkedServicesClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkedServicesClientListByFactoryResponse]{
		More: func(page LinkedServicesClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LinkedServicesClientListByFactoryResponse) (LinkedServicesClientListByFactoryResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LinkedServicesClient.NewListByFactoryPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			}, nil)
			if err != nil {
				return LinkedServicesClientListByFactoryResponse{}, err
			}
			return client.listByFactoryHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *LinkedServicesClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, _ *LinkedServicesClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *LinkedServicesClient) listByFactoryHandleResponse(resp *http.Response) (LinkedServicesClientListByFactoryResponse, error) {
	result := LinkedServicesClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServiceListResponse); err != nil {
		return LinkedServicesClientListByFactoryResponse{}, err
	}
	return result, nil
}
