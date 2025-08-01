// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

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

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the operation supported by the given provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - operationsPutContent - The operations content properties supplied to the CreateOrUpdate operation.
//   - options - OperationsClientCreateOrUpdateOptions contains the optional parameters for the OperationsClient.CreateOrUpdate
//     method.
func (client *OperationsClient) CreateOrUpdate(ctx context.Context, providerNamespace string, operationsPutContent OperationsPutContent, options *OperationsClientCreateOrUpdateOptions) (OperationsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "OperationsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, providerNamespace, operationsPutContent, options)
	if err != nil {
		return OperationsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OperationsClient) createOrUpdateCreateRequest(ctx context.Context, providerNamespace string, operationsPutContent OperationsPutContent, _ *OperationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/operations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, operationsPutContent); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *OperationsClient) createOrUpdateHandleResponse(resp *http.Response) (OperationsClientCreateOrUpdateResponse, error) {
	result := OperationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsPutContent); err != nil {
		return OperationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - options - OperationsClientDeleteOptions contains the optional parameters for the OperationsClient.Delete method.
func (client *OperationsClient) Delete(ctx context.Context, providerNamespace string, options *OperationsClientDeleteOptions) (OperationsClientDeleteResponse, error) {
	var err error
	const operationName = "OperationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return OperationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientDeleteResponse{}, err
	}
	return OperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OperationsClient) deleteCreateRequest(ctx context.Context, providerNamespace string, _ *OperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/operations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListPager - Lists all the operations supported by Microsoft.ProviderHub.
//
// Generated from API version 2024-09-01
//   - options - OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
func (client *OperationsClient) NewListPager(options *OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationsClientListResponse]{
		More: func(page OperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationsClientListResponse) (OperationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OperationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return OperationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, _ *OperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ProviderHub/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsClientListResponse, error) {
	result := OperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsDefinitionArrayResponseWithContinuation); err != nil {
		return OperationsClientListResponse{}, err
	}
	return result, nil
}

// ListByProviderRegistration - Gets the operations supported by the given provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - options - OperationsClientListByProviderRegistrationOptions contains the optional parameters for the OperationsClient.ListByProviderRegistration
//     method.
func (client *OperationsClient) ListByProviderRegistration(ctx context.Context, providerNamespace string, options *OperationsClientListByProviderRegistrationOptions) (OperationsClientListByProviderRegistrationResponse, error) {
	var err error
	const operationName = "OperationsClient.ListByProviderRegistration"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByProviderRegistrationCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return OperationsClientListByProviderRegistrationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientListByProviderRegistrationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientListByProviderRegistrationResponse{}, err
	}
	resp, err := client.listByProviderRegistrationHandleResponse(httpResp)
	return resp, err
}

// listByProviderRegistrationCreateRequest creates the ListByProviderRegistration request.
func (client *OperationsClient) listByProviderRegistrationCreateRequest(ctx context.Context, providerNamespace string, _ *OperationsClientListByProviderRegistrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/operations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProviderRegistrationHandleResponse handles the ListByProviderRegistration response.
func (client *OperationsClient) listByProviderRegistrationHandleResponse(resp *http.Response) (OperationsClientListByProviderRegistrationResponse, error) {
	result := OperationsClientListByProviderRegistrationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsDefinitionArray); err != nil {
		return OperationsClientListByProviderRegistrationResponse{}, err
	}
	return result, nil
}
