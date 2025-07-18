// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// TopicTypesClient contains the methods for the TopicTypes group.
// Don't use this type directly, use NewTopicTypesClient() instead.
type TopicTypesClient struct {
	internal *arm.Client
}

// NewTopicTypesClient creates a new instance of TopicTypesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTopicTypesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TopicTypesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TopicTypesClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get information about a topic type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - topicTypeName - Name of the topic type.
//   - options - TopicTypesClientGetOptions contains the optional parameters for the TopicTypesClient.Get method.
func (client *TopicTypesClient) Get(ctx context.Context, topicTypeName string, options *TopicTypesClientGetOptions) (TopicTypesClientGetResponse, error) {
	var err error
	const operationName = "TopicTypesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, topicTypeName, options)
	if err != nil {
		return TopicTypesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicTypesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TopicTypesClient) getCreateRequest(ctx context.Context, topicTypeName string, _ *TopicTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}"
	if topicTypeName == "" {
		return nil, errors.New("parameter topicTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicTypeName}", url.PathEscape(topicTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopicTypesClient) getHandleResponse(resp *http.Response) (TopicTypesClientGetResponse, error) {
	result := TopicTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicTypeInfo); err != nil {
		return TopicTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all registered topic types.
//
// Generated from API version 2025-04-01-preview
//   - options - TopicTypesClientListOptions contains the optional parameters for the TopicTypesClient.NewListPager method.
func (client *TopicTypesClient) NewListPager(options *TopicTypesClientListOptions) *runtime.Pager[TopicTypesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopicTypesClientListResponse]{
		More: func(page TopicTypesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *TopicTypesClientListResponse) (TopicTypesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopicTypesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return TopicTypesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TopicTypesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TopicTypesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TopicTypesClient) listCreateRequest(ctx context.Context, _ *TopicTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.EventGrid/topicTypes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TopicTypesClient) listHandleResponse(resp *http.Response) (TopicTypesClientListResponse, error) {
	result := TopicTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicTypesListResult); err != nil {
		return TopicTypesClientListResponse{}, err
	}
	return result, nil
}

// NewListEventTypesPager - List event types for a topic type.
//
// Generated from API version 2025-04-01-preview
//   - topicTypeName - Name of the topic type.
//   - options - TopicTypesClientListEventTypesOptions contains the optional parameters for the TopicTypesClient.NewListEventTypesPager
//     method.
func (client *TopicTypesClient) NewListEventTypesPager(topicTypeName string, options *TopicTypesClientListEventTypesOptions) *runtime.Pager[TopicTypesClientListEventTypesResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopicTypesClientListEventTypesResponse]{
		More: func(page TopicTypesClientListEventTypesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *TopicTypesClientListEventTypesResponse) (TopicTypesClientListEventTypesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopicTypesClient.NewListEventTypesPager")
			req, err := client.listEventTypesCreateRequest(ctx, topicTypeName, options)
			if err != nil {
				return TopicTypesClientListEventTypesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TopicTypesClientListEventTypesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TopicTypesClientListEventTypesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listEventTypesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listEventTypesCreateRequest creates the ListEventTypes request.
func (client *TopicTypesClient) listEventTypesCreateRequest(ctx context.Context, topicTypeName string, _ *TopicTypesClientListEventTypesOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventTypes"
	if topicTypeName == "" {
		return nil, errors.New("parameter topicTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicTypeName}", url.PathEscape(topicTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listEventTypesHandleResponse handles the ListEventTypes response.
func (client *TopicTypesClient) listEventTypesHandleResponse(resp *http.Response) (TopicTypesClientListEventTypesResponse, error) {
	result := TopicTypesClientListEventTypesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventTypesListResult); err != nil {
		return TopicTypesClientListEventTypesResponse{}, err
	}
	return result, nil
}
