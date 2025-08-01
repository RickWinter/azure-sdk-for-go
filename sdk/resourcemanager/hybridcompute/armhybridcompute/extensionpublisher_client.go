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

// ExtensionPublisherClient contains the methods for the ExtensionPublisher group.
// Don't use this type directly, use NewExtensionPublisherClient() instead.
type ExtensionPublisherClient struct {
	internal *arm.Client
}

// NewExtensionPublisherClient creates a new instance of ExtensionPublisherClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExtensionPublisherClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ExtensionPublisherClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExtensionPublisherClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Gets all Extension publishers based on the location
//
// Generated from API version 2025-02-19-preview
//   - location - The name of Azure region.
//   - options - ExtensionPublisherClientListOptions contains the optional parameters for the ExtensionPublisherClient.NewListPager
//     method.
func (client *ExtensionPublisherClient) NewListPager(location string, options *ExtensionPublisherClientListOptions) *runtime.Pager[ExtensionPublisherClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExtensionPublisherClientListResponse]{
		More: func(page ExtensionPublisherClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExtensionPublisherClientListResponse) (ExtensionPublisherClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExtensionPublisherClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return ExtensionPublisherClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ExtensionPublisherClient) listCreateRequest(ctx context.Context, location string, _ *ExtensionPublisherClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.HybridCompute/locations/{location}/publishers"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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

// listHandleResponse handles the List response.
func (client *ExtensionPublisherClient) listHandleResponse(resp *http.Response) (ExtensionPublisherClientListResponse, error) {
	result := ExtensionPublisherClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionPublisherListResult); err != nil {
		return ExtensionPublisherClientListResponse{}, err
	}
	return result, nil
}
