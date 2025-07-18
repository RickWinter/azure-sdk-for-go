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

// MachineRunCommandsClient contains the methods for the MachineRunCommands group.
// Don't use this type directly, use NewMachineRunCommandsClient() instead.
type MachineRunCommandsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMachineRunCommandsClient creates a new instance of MachineRunCommandsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMachineRunCommandsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MachineRunCommandsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MachineRunCommandsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - runCommandName - The name of the run command.
//   - runCommandProperties - Parameters supplied to the Create Run Command.
//   - options - MachineRunCommandsClientBeginCreateOrUpdateOptions contains the optional parameters for the MachineRunCommandsClient.BeginCreateOrUpdate
//     method.
func (client *MachineRunCommandsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommand, options *MachineRunCommandsClientBeginCreateOrUpdateOptions) (*runtime.Poller[MachineRunCommandsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, machineName, runCommandName, runCommandProperties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachineRunCommandsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MachineRunCommandsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - The operation to create or update a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
func (client *MachineRunCommandsClient) createOrUpdate(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommand, options *MachineRunCommandsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MachineRunCommandsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, machineName, runCommandName, runCommandProperties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MachineRunCommandsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommand, _ *MachineRunCommandsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, runCommandProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - runCommandName - The name of the run command.
//   - options - MachineRunCommandsClientBeginDeleteOptions contains the optional parameters for the MachineRunCommandsClient.BeginDelete
//     method.
func (client *MachineRunCommandsClient) BeginDelete(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, options *MachineRunCommandsClientBeginDeleteOptions) (*runtime.Poller[MachineRunCommandsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, machineName, runCommandName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachineRunCommandsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MachineRunCommandsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The operation to delete a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
func (client *MachineRunCommandsClient) deleteOperation(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, options *MachineRunCommandsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MachineRunCommandsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, machineName, runCommandName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MachineRunCommandsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, _ *MachineRunCommandsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation to get a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - runCommandName - The name of the run command.
//   - options - MachineRunCommandsClientGetOptions contains the optional parameters for the MachineRunCommandsClient.Get method.
func (client *MachineRunCommandsClient) Get(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, options *MachineRunCommandsClientGetOptions) (MachineRunCommandsClientGetResponse, error) {
	var err error
	const operationName = "MachineRunCommandsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, machineName, runCommandName, options)
	if err != nil {
		return MachineRunCommandsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachineRunCommandsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MachineRunCommandsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MachineRunCommandsClient) getCreateRequest(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, _ *MachineRunCommandsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// getHandleResponse handles the Get response.
func (client *MachineRunCommandsClient) getHandleResponse(resp *http.Response) (MachineRunCommandsClientGetResponse, error) {
	result := MachineRunCommandsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineRunCommand); err != nil {
		return MachineRunCommandsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation to get all the run commands of a non-Azure machine.
//
// Generated from API version 2025-02-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - options - MachineRunCommandsClientListOptions contains the optional parameters for the MachineRunCommandsClient.NewListPager
//     method.
func (client *MachineRunCommandsClient) NewListPager(resourceGroupName string, machineName string, options *MachineRunCommandsClientListOptions) *runtime.Pager[MachineRunCommandsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MachineRunCommandsClientListResponse]{
		More: func(page MachineRunCommandsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MachineRunCommandsClientListResponse) (MachineRunCommandsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MachineRunCommandsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, machineName, options)
			}, nil)
			if err != nil {
				return MachineRunCommandsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MachineRunCommandsClient) listCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *MachineRunCommandsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2025-02-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MachineRunCommandsClient) listHandleResponse(resp *http.Response) (MachineRunCommandsClientListResponse, error) {
	result := MachineRunCommandsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineRunCommandsListResult); err != nil {
		return MachineRunCommandsClientListResponse{}, err
	}
	return result, nil
}
