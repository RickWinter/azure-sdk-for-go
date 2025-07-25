// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DomainRegistrationProviderClient contains the methods for the DomainRegistrationProvider group.
// Don't use this type directly, use NewDomainRegistrationProviderClient() instead.
type DomainRegistrationProviderClient struct {
	internal *arm.Client
}

// NewDomainRegistrationProviderClient creates a new instance of DomainRegistrationProviderClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDomainRegistrationProviderClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*DomainRegistrationProviderClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DomainRegistrationProviderClient{
		internal: cl,
	}
	return client, nil
}

// NewListOperationsPager - Description for Implements Csm operations Api to exposes the list of available Csm Apis under
// the resource provider
//
// Generated from API version 2024-11-01
//   - options - DomainRegistrationProviderClientListOperationsOptions contains the optional parameters for the DomainRegistrationProviderClient.NewListOperationsPager
//     method.
func (client *DomainRegistrationProviderClient) NewListOperationsPager(options *DomainRegistrationProviderClientListOperationsOptions) *runtime.Pager[DomainRegistrationProviderClientListOperationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[DomainRegistrationProviderClientListOperationsResponse]{
		More: func(page DomainRegistrationProviderClientListOperationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DomainRegistrationProviderClientListOperationsResponse) (DomainRegistrationProviderClientListOperationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DomainRegistrationProviderClient.NewListOperationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listOperationsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DomainRegistrationProviderClientListOperationsResponse{}, err
			}
			return client.listOperationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *DomainRegistrationProviderClient) listOperationsCreateRequest(ctx context.Context, _ *DomainRegistrationProviderClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DomainRegistration/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *DomainRegistrationProviderClient) listOperationsHandleResponse(resp *http.Response) (DomainRegistrationProviderClientListOperationsResponse, error) {
	result := DomainRegistrationProviderClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CsmOperationCollection); err != nil {
		return DomainRegistrationProviderClientListOperationsResponse{}, err
	}
	return result, nil
}
