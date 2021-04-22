// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// LibraryClient contains the methods for the Library group.
// Don't use this type directly, use NewLibraryClient() instead.
type LibraryClient struct {
	con *Connection
}

// NewLibraryClient creates a new instance of LibraryClient with the specified values.
func NewLibraryClient(con *Connection) *LibraryClient {
	return &LibraryClient{con: con}
}

// Append - Append the content to the library resource created using the create operation. The maximum content size is 4MiB. Content larger than 4MiB must
// be appended in 4MiB chunks
func (client *LibraryClient) Append(ctx context.Context, libraryName string, content azcore.ReadSeekCloser, options *LibraryAppendOptions) (*http.Response, error) {
	req, err := client.appendCreateRequest(ctx, libraryName, content, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.appendHandleError(resp)
	}
	return resp.Response, nil
}

// appendCreateRequest creates the Append request.
func (client *LibraryClient) appendCreateRequest(ctx context.Context, libraryName string, content azcore.ReadSeekCloser, options *LibraryAppendOptions) (*azcore.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("comp", "appendblock")
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.BlobConditionAppendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*options.BlobConditionAppendPosition, 10))
	}
	req.Header.Set("Accept", "application/json")
	return req, req.SetBody(content, "application/octet-stream")
}

// appendHandleError handles the Append error response.
func (client *LibraryClient) appendHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginCreate - Creates a library with the library name.
func (client *LibraryClient) BeginCreate(ctx context.Context, libraryName string, options *LibraryBeginCreateOptions) (LibraryResourceInfoPollerResponse, error) {
	resp, err := client.create(ctx, libraryName, options)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	result := LibraryResourceInfoPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("LibraryClient.Create", resp, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	poller := &libraryResourceInfoPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LibraryResourceInfoResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreate creates a new LibraryResourceInfoPoller from the specified resume token.
// token - The value must come from a previous call to LibraryResourceInfoPoller.ResumeToken().
func (client *LibraryClient) ResumeCreate(ctx context.Context, token string) (LibraryResourceInfoPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("LibraryClient.Create", token, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	poller := &libraryResourceInfoPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	result := LibraryResourceInfoPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LibraryResourceInfoResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Create - Creates a library with the library name.
func (client *LibraryClient) create(ctx context.Context, libraryName string, options *LibraryBeginCreateOptions) (*azcore.Response, error) {
	req, err := client.createCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *LibraryClient) createCreateRequest(ctx context.Context, libraryName string, options *LibraryBeginCreateOptions) (*azcore.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *LibraryClient) createHandleResponse(resp *azcore.Response) (LibraryResourceInfoResponse, error) {
	var val *LibraryResourceInfo
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LibraryResourceInfoResponse{}, err
	}
	return LibraryResourceInfoResponse{RawResponse: resp.Response, LibraryResourceInfo: val}, nil
}

// createHandleError handles the Create error response.
func (client *LibraryClient) createHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Delete Library
func (client *LibraryClient) BeginDelete(ctx context.Context, libraryName string, options *LibraryBeginDeleteOptions) (LibraryResourceInfoPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, libraryName, options)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	result := LibraryResourceInfoPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("LibraryClient.Delete", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	poller := &libraryResourceInfoPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LibraryResourceInfoResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new LibraryResourceInfoPoller from the specified resume token.
// token - The value must come from a previous call to LibraryResourceInfoPoller.ResumeToken().
func (client *LibraryClient) ResumeDelete(ctx context.Context, token string) (LibraryResourceInfoPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("LibraryClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	poller := &libraryResourceInfoPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	result := LibraryResourceInfoPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LibraryResourceInfoResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Delete Library
func (client *LibraryClient) deleteOperation(ctx context.Context, libraryName string, options *LibraryBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusConflict) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LibraryClient) deleteCreateRequest(ctx context.Context, libraryName string, options *LibraryBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *LibraryClient) deleteHandleResponse(resp *azcore.Response) (LibraryResourceInfoResponse, error) {
	var val *LibraryResourceInfo
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LibraryResourceInfoResponse{}, err
	}
	return LibraryResourceInfoResponse{RawResponse: resp.Response, LibraryResourceInfo: val}, nil
}

// deleteHandleError handles the Delete error response.
func (client *LibraryClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginFlush - Flush Library
func (client *LibraryClient) BeginFlush(ctx context.Context, libraryName string, options *LibraryBeginFlushOptions) (LibraryResourceInfoPollerResponse, error) {
	resp, err := client.flush(ctx, libraryName, options)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	result := LibraryResourceInfoPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("LibraryClient.Flush", resp, client.con.Pipeline(), client.flushHandleError)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	poller := &libraryResourceInfoPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LibraryResourceInfoResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeFlush creates a new LibraryResourceInfoPoller from the specified resume token.
// token - The value must come from a previous call to LibraryResourceInfoPoller.ResumeToken().
func (client *LibraryClient) ResumeFlush(ctx context.Context, token string) (LibraryResourceInfoPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("LibraryClient.Flush", token, client.con.Pipeline(), client.flushHandleError)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	poller := &libraryResourceInfoPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LibraryResourceInfoPollerResponse{}, err
	}
	result := LibraryResourceInfoPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LibraryResourceInfoResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Flush - Flush Library
func (client *LibraryClient) flush(ctx context.Context, libraryName string, options *LibraryBeginFlushOptions) (*azcore.Response, error) {
	req, err := client.flushCreateRequest(ctx, libraryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.flushHandleError(resp)
	}
	return resp, nil
}

// flushCreateRequest creates the Flush request.
func (client *LibraryClient) flushCreateRequest(ctx context.Context, libraryName string, options *LibraryBeginFlushOptions) (*azcore.Request, error) {
	urlPath := "/libraries/{libraryName}/flush"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// flushHandleResponse handles the Flush response.
func (client *LibraryClient) flushHandleResponse(resp *azcore.Response) (LibraryResourceInfoResponse, error) {
	var val *LibraryResourceInfo
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LibraryResourceInfoResponse{}, err
	}
	return LibraryResourceInfoResponse{RawResponse: resp.Response, LibraryResourceInfo: val}, nil
}

// flushHandleError handles the Flush error response.
func (client *LibraryClient) flushHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Get Library
func (client *LibraryClient) Get(ctx context.Context, libraryName string, options *LibraryGetOptions) (LibraryResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, libraryName, options)
	if err != nil {
		return LibraryResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LibraryResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return LibraryResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LibraryClient) getCreateRequest(ctx context.Context, libraryName string, options *LibraryGetOptions) (*azcore.Request, error) {
	urlPath := "/libraries/{libraryName}"
	if libraryName == "" {
		return nil, errors.New("parameter libraryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{libraryName}", url.PathEscape(libraryName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LibraryClient) getHandleResponse(resp *azcore.Response) (LibraryResourceResponse, error) {
	var val *LibraryResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LibraryResourceResponse{}, err
	}
	return LibraryResourceResponse{RawResponse: resp.Response, LibraryResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *LibraryClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetOperationResult - Get Operation result for Library
// Possible return types are *LibraryResourceResponse, *OperationResultResponse
func (client *LibraryClient) GetOperationResult(ctx context.Context, operationID string, options *LibraryGetOperationResultOptions) (interface{}, error) {
	req, err := client.getOperationResultCreateRequest(ctx, operationID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.getOperationResultHandleError(resp)
	}
	return client.getOperationResultHandleResponse(resp)
}

// getOperationResultCreateRequest creates the GetOperationResult request.
func (client *LibraryClient) getOperationResultCreateRequest(ctx context.Context, operationID string, options *LibraryGetOperationResultOptions) (*azcore.Request, error) {
	urlPath := "/libraryOperationResults/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getOperationResultHandleResponse handles the GetOperationResult response.
func (client *LibraryClient) getOperationResultHandleResponse(resp *azcore.Response) (interface{}, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		var val *LibraryResource
		if err := resp.UnmarshalAsJSON(&val); err != nil {
			return nil, err
		}
		return LibraryResourceResponse{RawResponse: resp.Response, LibraryResource: val}, nil
	case http.StatusAccepted:
		var val *OperationResult
		if err := resp.UnmarshalAsJSON(&val); err != nil {
			return nil, err
		}
		return OperationResultResponse{RawResponse: resp.Response, OperationResult: val}, nil
	default:
		return nil, fmt.Errorf("unhandled HTTP status code %d", resp.StatusCode)
	}
}

// getOperationResultHandleError handles the GetOperationResult error response.
func (client *LibraryClient) getOperationResultHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists Library.
func (client *LibraryClient) List(options *LibraryListOptions) LibraryListResponsePager {
	return &libraryListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp LibraryListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LibraryListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *LibraryClient) listCreateRequest(ctx context.Context, options *LibraryListOptions) (*azcore.Request, error) {
	urlPath := "/libraries"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LibraryClient) listHandleResponse(resp *azcore.Response) (LibraryListResponseResponse, error) {
	var val *LibraryListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LibraryListResponseResponse{}, err
	}
	return LibraryListResponseResponse{RawResponse: resp.Response, LibraryListResponse: val}, nil
}

// listHandleError handles the List error response.
func (client *LibraryClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
