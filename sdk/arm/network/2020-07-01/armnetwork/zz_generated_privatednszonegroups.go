// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PrivateDNSZoneGroupsClient contains the methods for the PrivateDNSZoneGroups group.
// Don't use this type directly, use NewPrivateDNSZoneGroupsClient() instead.
type PrivateDNSZoneGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPrivateDNSZoneGroupsClient creates a new instance of PrivateDNSZoneGroupsClient with the specified values.
func NewPrivateDNSZoneGroupsClient(con *armcore.Connection, subscriptionID string) PrivateDNSZoneGroupsClient {
	return PrivateDNSZoneGroupsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client PrivateDNSZoneGroupsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates a private dns zone group in the specified private endpoint.
func (client PrivateDNSZoneGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsBeginCreateOrUpdateOptions) (PrivateDNSZoneGroupPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName, parameters, options)
	if err != nil {
		return PrivateDNSZoneGroupPollerResponse{}, err
	}
	result := PrivateDNSZoneGroupPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateDNSZoneGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return PrivateDNSZoneGroupPollerResponse{}, err
	}
	poller := &privateDnsZoneGroupPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PrivateDNSZoneGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new PrivateDNSZoneGroupPoller from the specified resume token.
// token - The value must come from a previous call to PrivateDNSZoneGroupPoller.ResumeToken().
func (client PrivateDNSZoneGroupsClient) ResumeCreateOrUpdate(token string) (PrivateDNSZoneGroupPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateDNSZoneGroupsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &privateDnsZoneGroupPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a private dns zone group in the specified private endpoint.
func (client PrivateDNSZoneGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client PrivateDNSZoneGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDnsZoneGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client PrivateDNSZoneGroupsClient) createOrUpdateHandleResponse(resp *azcore.Response) (PrivateDNSZoneGroupResponse, error) {
	result := PrivateDNSZoneGroupResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.PrivateDNSZoneGroup)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client PrivateDNSZoneGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified private dns zone group.
func (client PrivateDNSZoneGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, options *PrivateDNSZoneGroupsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PrivateDNSZoneGroupsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client PrivateDNSZoneGroupsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PrivateDNSZoneGroupsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified private dns zone group.
func (client PrivateDNSZoneGroupsClient) delete(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, options *PrivateDNSZoneGroupsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client PrivateDNSZoneGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, options *PrivateDNSZoneGroupsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDnsZoneGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client PrivateDNSZoneGroupsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the private dns zone group resource by specified private dns zone group name.
func (client PrivateDNSZoneGroupsClient) Get(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, options *PrivateDNSZoneGroupsGetOptions) (PrivateDNSZoneGroupResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDnsZoneGroupName, options)
	if err != nil {
		return PrivateDNSZoneGroupResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return PrivateDNSZoneGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateDNSZoneGroupResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return PrivateDNSZoneGroupResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client PrivateDNSZoneGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDnsZoneGroupName string, options *PrivateDNSZoneGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDnsZoneGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client PrivateDNSZoneGroupsClient) getHandleResponse(resp *azcore.Response) (PrivateDNSZoneGroupResponse, error) {
	result := PrivateDNSZoneGroupResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.PrivateDNSZoneGroup)
	return result, err
}

// getHandleError handles the Get error response.
func (client PrivateDNSZoneGroupsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all private dns zone groups in a private endpoint.
func (client PrivateDNSZoneGroupsClient) List(privateEndpointName string, resourceGroupName string, options *PrivateDNSZoneGroupsListOptions) PrivateDNSZoneGroupListResultPager {
	return &privateDnsZoneGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, privateEndpointName, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp PrivateDNSZoneGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PrivateDNSZoneGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client PrivateDNSZoneGroupsClient) listCreateRequest(ctx context.Context, privateEndpointName string, resourceGroupName string, options *PrivateDNSZoneGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups"
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client PrivateDNSZoneGroupsClient) listHandleResponse(resp *azcore.Response) (PrivateDNSZoneGroupListResultResponse, error) {
	result := PrivateDNSZoneGroupListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.PrivateDNSZoneGroupListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client PrivateDNSZoneGroupsClient) listHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
