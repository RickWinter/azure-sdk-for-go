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

// VirtualHubRouteTableV2SClient contains the methods for the VirtualHubRouteTableV2S group.
// Don't use this type directly, use NewVirtualHubRouteTableV2SClient() instead.
type VirtualHubRouteTableV2SClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualHubRouteTableV2SClient creates a new instance of VirtualHubRouteTableV2SClient with the specified values.
func NewVirtualHubRouteTableV2SClient(con *armcore.Connection, subscriptionID string) VirtualHubRouteTableV2SClient {
	return VirtualHubRouteTableV2SClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VirtualHubRouteTableV2SClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing VirtualHubRouteTableV2.
func (client VirtualHubRouteTableV2SClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SBeginCreateOrUpdateOptions) (VirtualHubRouteTableV2PollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, routeTableName, virtualHubRouteTableV2Parameters, options)
	if err != nil {
		return VirtualHubRouteTableV2PollerResponse{}, err
	}
	result := VirtualHubRouteTableV2PollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualHubRouteTableV2SClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualHubRouteTableV2PollerResponse{}, err
	}
	poller := &virtualHubRouteTableV2Poller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualHubRouteTableV2Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualHubRouteTableV2Poller from the specified resume token.
// token - The value must come from a previous call to VirtualHubRouteTableV2Poller.ResumeToken().
func (client VirtualHubRouteTableV2SClient) ResumeCreateOrUpdate(token string) (VirtualHubRouteTableV2Poller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualHubRouteTableV2SClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualHubRouteTableV2Poller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing VirtualHubRouteTableV2.
func (client VirtualHubRouteTableV2SClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, virtualHubRouteTableV2Parameters, options)
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
func (client VirtualHubRouteTableV2SClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, virtualHubRouteTableV2Parameters VirtualHubRouteTableV2, options *VirtualHubRouteTableV2SBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(virtualHubRouteTableV2Parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client VirtualHubRouteTableV2SClient) createOrUpdateHandleResponse(resp *azcore.Response) (VirtualHubRouteTableV2Response, error) {
	result := VirtualHubRouteTableV2Response{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualHubRouteTableV2)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client VirtualHubRouteTableV2SClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes a VirtualHubRouteTableV2.
func (client VirtualHubRouteTableV2SClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualHubRouteTableV2SClient.Delete", "location", resp, client.deleteHandleError)
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
func (client VirtualHubRouteTableV2SClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualHubRouteTableV2SClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes a VirtualHubRouteTableV2.
func (client VirtualHubRouteTableV2SClient) delete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
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
func (client VirtualHubRouteTableV2SClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
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
func (client VirtualHubRouteTableV2SClient) deleteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves the details of a VirtualHubRouteTableV2.
func (client VirtualHubRouteTableV2SClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SGetOptions) (VirtualHubRouteTableV2Response, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return VirtualHubRouteTableV2Response{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualHubRouteTableV2Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualHubRouteTableV2Response{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return VirtualHubRouteTableV2Response{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client VirtualHubRouteTableV2SClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *VirtualHubRouteTableV2SGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
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
func (client VirtualHubRouteTableV2SClient) getHandleResponse(resp *azcore.Response) (VirtualHubRouteTableV2Response, error) {
	result := VirtualHubRouteTableV2Response{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualHubRouteTableV2)
	return result, err
}

// getHandleError handles the Get error response.
func (client VirtualHubRouteTableV2SClient) getHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Retrieves the details of all VirtualHubRouteTableV2s.
func (client VirtualHubRouteTableV2SClient) List(resourceGroupName string, virtualHubName string, options *VirtualHubRouteTableV2SListOptions) ListVirtualHubRouteTableV2SResultPager {
	return &listVirtualHubRouteTableV2SResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ListVirtualHubRouteTableV2SResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVirtualHubRouteTableV2SResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client VirtualHubRouteTableV2SClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubRouteTableV2SListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
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
func (client VirtualHubRouteTableV2SClient) listHandleResponse(resp *azcore.Response) (ListVirtualHubRouteTableV2SResultResponse, error) {
	result := ListVirtualHubRouteTableV2SResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ListVirtualHubRouteTableV2SResult)
	return result, err
}

// listHandleError handles the List error response.
func (client VirtualHubRouteTableV2SClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
