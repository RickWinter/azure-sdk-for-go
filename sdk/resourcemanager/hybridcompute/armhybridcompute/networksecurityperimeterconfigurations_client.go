// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// NetworkSecurityPerimeterConfigurationsClient contains the methods for the NetworkSecurityPerimeterConfigurations group.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationsClient() instead.
type NetworkSecurityPerimeterConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkSecurityPerimeterConfigurationsClient creates a new instance of NetworkSecurityPerimeterConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkSecurityPerimeterConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkSecurityPerimeterConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkSecurityPerimeterConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetByPrivateLinkScope - Gets the network security perimeter configuration for a private link scope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scopeName - The name of the Azure Arc PrivateLinkScope resource.
//   - perimeterName - The name, in the format {perimeterGuid}.{associationName}, of the Network Security Perimeter resource.
//   - options - NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeOptions contains the optional parameters for
//     the NetworkSecurityPerimeterConfigurationsClient.GetByPrivateLinkScope method.
func (client *NetworkSecurityPerimeterConfigurationsClient) GetByPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string, perimeterName string, options *NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeOptions) (NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterConfigurationsClient.GetByPrivateLinkScope"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByPrivateLinkScopeCreateRequest(ctx, resourceGroupName, scopeName, perimeterName, options)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse{}, err
	}
	resp, err := client.getByPrivateLinkScopeHandleResponse(httpResp)
	return resp, err
}

// getByPrivateLinkScopeCreateRequest creates the GetByPrivateLinkScope request.
func (client *NetworkSecurityPerimeterConfigurationsClient) getByPrivateLinkScopeCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, perimeterName string, _ *NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}/networkSecurityPerimeterConfigurations/{perimeterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	if perimeterName == "" {
		return nil, errors.New("parameter perimeterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{perimeterName}", url.PathEscape(perimeterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByPrivateLinkScopeHandleResponse handles the GetByPrivateLinkScope response.
func (client *NetworkSecurityPerimeterConfigurationsClient) getByPrivateLinkScopeHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfiguration); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse{}, err
	}
	return result, nil
}

// NewListByPrivateLinkScopePager - Lists the network security perimeter configurations for a private link scope.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scopeName - The name of the Azure Arc PrivateLinkScope resource.
//   - options - NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeOptions contains the optional parameters for
//     the NetworkSecurityPerimeterConfigurationsClient.NewListByPrivateLinkScopePager method.
func (client *NetworkSecurityPerimeterConfigurationsClient) NewListByPrivateLinkScopePager(resourceGroupName string, scopeName string, options *NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeOptions) *runtime.Pager[NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse]{
		More: func(page NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse) (NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkSecurityPerimeterConfigurationsClient.NewListByPrivateLinkScopePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByPrivateLinkScopeCreateRequest(ctx, resourceGroupName, scopeName, options)
			}, nil)
			if err != nil {
				return NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse{}, err
			}
			return client.listByPrivateLinkScopeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByPrivateLinkScopeCreateRequest creates the ListByPrivateLinkScope request.
func (client *NetworkSecurityPerimeterConfigurationsClient) listByPrivateLinkScopeCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, _ *NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}/networkSecurityPerimeterConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPrivateLinkScopeHandleResponse handles the ListByPrivateLinkScope response.
func (client *NetworkSecurityPerimeterConfigurationsClient) listByPrivateLinkScopeHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfigurationListResult); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse{}, err
	}
	return result, nil
}

// BeginReconcileForPrivateLinkScope - Forces the network security perimeter configuration to refresh for a private link scope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - scopeName - The name of the Azure Arc PrivateLinkScope resource.
//   - perimeterName - The name, in the format {perimeterGuid}.{associationName}, of the Network Security Perimeter resource.
//   - options - NetworkSecurityPerimeterConfigurationsClientBeginReconcileForPrivateLinkScopeOptions contains the optional parameters
//     for the NetworkSecurityPerimeterConfigurationsClient.BeginReconcileForPrivateLinkScope method.
func (client *NetworkSecurityPerimeterConfigurationsClient) BeginReconcileForPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string, perimeterName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileForPrivateLinkScopeOptions) (*runtime.Poller[NetworkSecurityPerimeterConfigurationsClientReconcileForPrivateLinkScopeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reconcileForPrivateLinkScope(ctx, resourceGroupName, scopeName, perimeterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkSecurityPerimeterConfigurationsClientReconcileForPrivateLinkScopeResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkSecurityPerimeterConfigurationsClientReconcileForPrivateLinkScopeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ReconcileForPrivateLinkScope - Forces the network security perimeter configuration to refresh for a private link scope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcileForPrivateLinkScope(ctx context.Context, resourceGroupName string, scopeName string, perimeterName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileForPrivateLinkScopeOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterConfigurationsClient.BeginReconcileForPrivateLinkScope"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.reconcileForPrivateLinkScopeCreateRequest(ctx, resourceGroupName, scopeName, perimeterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// reconcileForPrivateLinkScopeCreateRequest creates the ReconcileForPrivateLinkScope request.
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcileForPrivateLinkScopeCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, perimeterName string, _ *NetworkSecurityPerimeterConfigurationsClientBeginReconcileForPrivateLinkScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/privateLinkScopes/{scopeName}/networkSecurityPerimeterConfigurations/{perimeterName}/reconcile"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	if perimeterName == "" {
		return nil, errors.New("parameter perimeterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{perimeterName}", url.PathEscape(perimeterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
