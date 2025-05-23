// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// WorkspaceAPIOperationPolicyClient contains the methods for the WorkspaceAPIOperationPolicy group.
// Don't use this type directly, use NewWorkspaceAPIOperationPolicyClient() instead.
type WorkspaceAPIOperationPolicyClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceAPIOperationPolicyClient creates a new instance of WorkspaceAPIOperationPolicyClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceAPIOperationPolicyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceAPIOperationPolicyClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceAPIOperationPolicyClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates policy configuration for the API Operation level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
//   - policyID - The identifier of the Policy.
//   - parameters - The policy contents to apply.
//   - options - WorkspaceAPIOperationPolicyClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceAPIOperationPolicyClient.CreateOrUpdate
//     method.
func (client *WorkspaceAPIOperationPolicyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID PolicyIDName, parameters PolicyContract, options *WorkspaceAPIOperationPolicyClientCreateOrUpdateOptions) (WorkspaceAPIOperationPolicyClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkspaceAPIOperationPolicyClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, operationID, policyID, parameters, options)
	if err != nil {
		return WorkspaceAPIOperationPolicyClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceAPIOperationPolicyClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceAPIOperationPolicyClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceAPIOperationPolicyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID PolicyIDName, parameters PolicyContract, options *WorkspaceAPIOperationPolicyClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/operations/{operationId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceAPIOperationPolicyClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceAPIOperationPolicyClientCreateOrUpdateResponse, error) {
	result := WorkspaceAPIOperationPolicyClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return WorkspaceAPIOperationPolicyClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the policy configuration at the Api Operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
//   - policyID - The identifier of the Policy.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - WorkspaceAPIOperationPolicyClientDeleteOptions contains the optional parameters for the WorkspaceAPIOperationPolicyClient.Delete
//     method.
func (client *WorkspaceAPIOperationPolicyClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID PolicyIDName, ifMatch string, options *WorkspaceAPIOperationPolicyClientDeleteOptions) (WorkspaceAPIOperationPolicyClientDeleteResponse, error) {
	var err error
	const operationName = "WorkspaceAPIOperationPolicyClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, operationID, policyID, ifMatch, options)
	if err != nil {
		return WorkspaceAPIOperationPolicyClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceAPIOperationPolicyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceAPIOperationPolicyClientDeleteResponse{}, err
	}
	return WorkspaceAPIOperationPolicyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceAPIOperationPolicyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID PolicyIDName, ifMatch string, _ *WorkspaceAPIOperationPolicyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/operations/{operationId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	return req, nil
}

// Get - Get the policy configuration at the API Operation level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
//   - policyID - The identifier of the Policy.
//   - options - WorkspaceAPIOperationPolicyClientGetOptions contains the optional parameters for the WorkspaceAPIOperationPolicyClient.Get
//     method.
func (client *WorkspaceAPIOperationPolicyClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID PolicyIDName, options *WorkspaceAPIOperationPolicyClientGetOptions) (WorkspaceAPIOperationPolicyClientGetResponse, error) {
	var err error
	const operationName = "WorkspaceAPIOperationPolicyClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, operationID, policyID, options)
	if err != nil {
		return WorkspaceAPIOperationPolicyClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceAPIOperationPolicyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceAPIOperationPolicyClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceAPIOperationPolicyClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID PolicyIDName, options *WorkspaceAPIOperationPolicyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/operations/{operationId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	if options != nil && options.Format != nil {
		reqQP.Set("format", string(*options.Format))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceAPIOperationPolicyClient) getHandleResponse(resp *http.Response) (WorkspaceAPIOperationPolicyClientGetResponse, error) {
	result := WorkspaceAPIOperationPolicyClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return WorkspaceAPIOperationPolicyClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the API operation policy specified by its identifier.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
//   - policyID - The identifier of the Policy.
//   - options - WorkspaceAPIOperationPolicyClientGetEntityTagOptions contains the optional parameters for the WorkspaceAPIOperationPolicyClient.GetEntityTag
//     method.
func (client *WorkspaceAPIOperationPolicyClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID PolicyIDName, options *WorkspaceAPIOperationPolicyClientGetEntityTagOptions) (WorkspaceAPIOperationPolicyClientGetEntityTagResponse, error) {
	var err error
	const operationName = "WorkspaceAPIOperationPolicyClient.GetEntityTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, operationID, policyID, options)
	if err != nil {
		return WorkspaceAPIOperationPolicyClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceAPIOperationPolicyClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceAPIOperationPolicyClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *WorkspaceAPIOperationPolicyClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID PolicyIDName, _ *WorkspaceAPIOperationPolicyClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/operations/{operationId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *WorkspaceAPIOperationPolicyClient) getEntityTagHandleResponse(resp *http.Response) (WorkspaceAPIOperationPolicyClientGetEntityTagResponse, error) {
	result := WorkspaceAPIOperationPolicyClientGetEntityTagResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// NewListByOperationPager - Get the list of policy configuration at the API Operation level.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
//   - options - WorkspaceAPIOperationPolicyClientListByOperationOptions contains the optional parameters for the WorkspaceAPIOperationPolicyClient.NewListByOperationPager
//     method.
func (client *WorkspaceAPIOperationPolicyClient) NewListByOperationPager(resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, options *WorkspaceAPIOperationPolicyClientListByOperationOptions) *runtime.Pager[WorkspaceAPIOperationPolicyClientListByOperationResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceAPIOperationPolicyClientListByOperationResponse]{
		More: func(page WorkspaceAPIOperationPolicyClientListByOperationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceAPIOperationPolicyClientListByOperationResponse) (WorkspaceAPIOperationPolicyClientListByOperationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspaceAPIOperationPolicyClient.NewListByOperationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByOperationCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, operationID, options)
			}, nil)
			if err != nil {
				return WorkspaceAPIOperationPolicyClientListByOperationResponse{}, err
			}
			return client.listByOperationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByOperationCreateRequest creates the ListByOperation request.
func (client *WorkspaceAPIOperationPolicyClient) listByOperationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, _ *WorkspaceAPIOperationPolicyClientListByOperationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}/operations/{operationId}/policies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByOperationHandleResponse handles the ListByOperation response.
func (client *WorkspaceAPIOperationPolicyClient) listByOperationHandleResponse(resp *http.Response) (WorkspaceAPIOperationPolicyClientListByOperationResponse, error) {
	result := WorkspaceAPIOperationPolicyClientListByOperationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyCollection); err != nil {
		return WorkspaceAPIOperationPolicyClientListByOperationResponse{}, err
	}
	return result, nil
}
