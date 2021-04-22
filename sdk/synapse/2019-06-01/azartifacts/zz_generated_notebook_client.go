// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// NotebookClient contains the methods for the Notebook group.
// Don't use this type directly, use NewNotebookClient() instead.
type NotebookClient struct {
	con *Connection
}

// NewNotebookClient creates a new instance of NotebookClient with the specified values.
func NewNotebookClient(con *Connection) *NotebookClient {
	return &NotebookClient{con: con}
}

// BeginCreateOrUpdateNotebook - Creates or updates a Note Book.
func (client *NotebookClient) BeginCreateOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (NotebookResourcePollerResponse, error) {
	resp, err := client.createOrUpdateNotebook(ctx, notebookName, notebook, options)
	if err != nil {
		return NotebookResourcePollerResponse{}, err
	}
	result := NotebookResourcePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("NotebookClient.CreateOrUpdateNotebook", resp, client.con.Pipeline(), client.createOrUpdateNotebookHandleError)
	if err != nil {
		return NotebookResourcePollerResponse{}, err
	}
	poller := &notebookResourcePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NotebookResourceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdateNotebook creates a new NotebookResourcePoller from the specified resume token.
// token - The value must come from a previous call to NotebookResourcePoller.ResumeToken().
func (client *NotebookClient) ResumeCreateOrUpdateNotebook(ctx context.Context, token string) (NotebookResourcePollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("NotebookClient.CreateOrUpdateNotebook", token, client.con.Pipeline(), client.createOrUpdateNotebookHandleError)
	if err != nil {
		return NotebookResourcePollerResponse{}, err
	}
	poller := &notebookResourcePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return NotebookResourcePollerResponse{}, err
	}
	result := NotebookResourcePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NotebookResourceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdateNotebook - Creates or updates a Note Book.
func (client *NotebookClient) createOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateNotebookCreateRequest(ctx, notebookName, notebook, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateNotebookHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateNotebookCreateRequest creates the CreateOrUpdateNotebook request.
func (client *NotebookClient) createOrUpdateNotebookCreateRequest(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(notebook)
}

// createOrUpdateNotebookHandleResponse handles the CreateOrUpdateNotebook response.
func (client *NotebookClient) createOrUpdateNotebookHandleResponse(resp *azcore.Response) (NotebookResourceResponse, error) {
	var val *NotebookResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NotebookResourceResponse{}, err
	}
	return NotebookResourceResponse{RawResponse: resp.Response, NotebookResource: val}, nil
}

// createOrUpdateNotebookHandleError handles the CreateOrUpdateNotebook error response.
func (client *NotebookClient) createOrUpdateNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDeleteNotebook - Deletes a Note book.
func (client *NotebookClient) BeginDeleteNotebook(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteNotebook(ctx, notebookName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("NotebookClient.DeleteNotebook", resp, client.con.Pipeline(), client.deleteNotebookHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDeleteNotebook creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *NotebookClient) ResumeDeleteNotebook(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("NotebookClient.DeleteNotebook", token, client.con.Pipeline(), client.deleteNotebookHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// DeleteNotebook - Deletes a Note book.
func (client *NotebookClient) deleteNotebook(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (*azcore.Response, error) {
	req, err := client.deleteNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteNotebookHandleError(resp)
	}
	return resp, nil
}

// deleteNotebookCreateRequest creates the DeleteNotebook request.
func (client *NotebookClient) deleteNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
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

// deleteNotebookHandleError handles the DeleteNotebook error response.
func (client *NotebookClient) deleteNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebook - Gets a Note Book.
func (client *NotebookClient) GetNotebook(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (NotebookResourceResponse, error) {
	req, err := client.getNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return NotebookResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NotebookResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return NotebookResourceResponse{}, client.getNotebookHandleError(resp)
	}
	return client.getNotebookHandleResponse(resp)
}

// getNotebookCreateRequest creates the GetNotebook request.
func (client *NotebookClient) getNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNotebookHandleResponse handles the GetNotebook response.
func (client *NotebookClient) getNotebookHandleResponse(resp *azcore.Response) (NotebookResourceResponse, error) {
	var val *NotebookResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NotebookResourceResponse{}, err
	}
	return NotebookResourceResponse{RawResponse: resp.Response, NotebookResource: val}, nil
}

// getNotebookHandleError handles the GetNotebook error response.
func (client *NotebookClient) getNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebookSummaryByWorkSpace - Lists a summary of Notebooks.
func (client *NotebookClient) GetNotebookSummaryByWorkSpace(options *NotebookGetNotebookSummaryByWorkSpaceOptions) NotebookListResponsePager {
	return &notebookListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getNotebookSummaryByWorkSpaceCreateRequest(ctx, options)
		},
		responder: client.getNotebookSummaryByWorkSpaceHandleResponse,
		errorer:   client.getNotebookSummaryByWorkSpaceHandleError,
		advancer: func(ctx context.Context, resp NotebookListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getNotebookSummaryByWorkSpaceCreateRequest creates the GetNotebookSummaryByWorkSpace request.
func (client *NotebookClient) getNotebookSummaryByWorkSpaceCreateRequest(ctx context.Context, options *NotebookGetNotebookSummaryByWorkSpaceOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/summary"
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

// getNotebookSummaryByWorkSpaceHandleResponse handles the GetNotebookSummaryByWorkSpace response.
func (client *NotebookClient) getNotebookSummaryByWorkSpaceHandleResponse(resp *azcore.Response) (NotebookListResponseResponse, error) {
	var val *NotebookListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NotebookListResponseResponse{}, err
	}
	return NotebookListResponseResponse{RawResponse: resp.Response, NotebookListResponse: val}, nil
}

// getNotebookSummaryByWorkSpaceHandleError handles the GetNotebookSummaryByWorkSpace error response.
func (client *NotebookClient) getNotebookSummaryByWorkSpaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotebooksByWorkspace - Lists Notebooks.
func (client *NotebookClient) GetNotebooksByWorkspace(options *NotebookGetNotebooksByWorkspaceOptions) NotebookListResponsePager {
	return &notebookListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getNotebooksByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getNotebooksByWorkspaceHandleResponse,
		errorer:   client.getNotebooksByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp NotebookListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getNotebooksByWorkspaceCreateRequest creates the GetNotebooksByWorkspace request.
func (client *NotebookClient) getNotebooksByWorkspaceCreateRequest(ctx context.Context, options *NotebookGetNotebooksByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/notebooks"
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

// getNotebooksByWorkspaceHandleResponse handles the GetNotebooksByWorkspace response.
func (client *NotebookClient) getNotebooksByWorkspaceHandleResponse(resp *azcore.Response) (NotebookListResponseResponse, error) {
	var val *NotebookListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NotebookListResponseResponse{}, err
	}
	return NotebookListResponseResponse{RawResponse: resp.Response, NotebookListResponse: val}, nil
}

// getNotebooksByWorkspaceHandleError handles the GetNotebooksByWorkspace error response.
func (client *NotebookClient) getNotebooksByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginRenameNotebook - Renames a notebook.
func (client *NotebookClient) BeginRenameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (HTTPPollerResponse, error) {
	resp, err := client.renameNotebook(ctx, notebookName, request, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("NotebookClient.RenameNotebook", resp, client.con.Pipeline(), client.renameNotebookHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRenameNotebook creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *NotebookClient) ResumeRenameNotebook(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("NotebookClient.RenameNotebook", token, client.con.Pipeline(), client.renameNotebookHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// RenameNotebook - Renames a notebook.
func (client *NotebookClient) renameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (*azcore.Response, error) {
	req, err := client.renameNotebookCreateRequest(ctx, notebookName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameNotebookHandleError(resp)
	}
	return resp, nil
}

// renameNotebookCreateRequest creates the RenameNotebook request.
func (client *NotebookClient) renameNotebookCreateRequest(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}/rename"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameNotebookHandleError handles the RenameNotebook error response.
func (client *NotebookClient) renameNotebookHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
