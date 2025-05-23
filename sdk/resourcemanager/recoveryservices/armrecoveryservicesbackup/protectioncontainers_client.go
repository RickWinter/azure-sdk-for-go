// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// ProtectionContainersClient contains the methods for the ProtectionContainers group.
// Don't use this type directly, use NewProtectionContainersClient() instead.
type ProtectionContainersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProtectionContainersClient creates a new instance of ProtectionContainersClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProtectionContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProtectionContainersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProtectionContainersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets details of the specific container registered to your Recovery Services Vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Name of the fabric where the container belongs.
//   - containerName - Name of the container whose details need to be fetched.
//   - options - ProtectionContainersClientGetOptions contains the optional parameters for the ProtectionContainersClient.Get
//     method.
func (client *ProtectionContainersClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersClientGetOptions) (ProtectionContainersClientGetResponse, error) {
	var err error
	const operationName = "ProtectionContainersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, options)
	if err != nil {
		return ProtectionContainersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionContainersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProtectionContainersClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, _ *ProtectionContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProtectionContainersClient) getHandleResponse(resp *http.Response) (ProtectionContainersClientGetResponse, error) {
	result := ProtectionContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerResource); err != nil {
		return ProtectionContainersClientGetResponse{}, err
	}
	return result, nil
}

// Inquire - This is an async operation and the results should be tracked using location header or Azure-async-url.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric Name associated with the container.
//   - containerName - Name of the container in which inquiry needs to be triggered.
//   - options - ProtectionContainersClientInquireOptions contains the optional parameters for the ProtectionContainersClient.Inquire
//     method.
func (client *ProtectionContainersClient) Inquire(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersClientInquireOptions) (ProtectionContainersClientInquireResponse, error) {
	var err error
	const operationName = "ProtectionContainersClient.Inquire"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.inquireCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, options)
	if err != nil {
		return ProtectionContainersClientInquireResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionContainersClientInquireResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionContainersClientInquireResponse{}, err
	}
	return ProtectionContainersClientInquireResponse{}, nil
}

// inquireCreateRequest creates the Inquire request.
func (client *ProtectionContainersClient) inquireCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersClientInquireOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/inquire"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Refresh - Discovers all the containers in the subscription that can be backed up to Recovery Services Vault. This is an
// asynchronous operation. To know the status of the operation, call
// GetRefreshOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated the container.
//   - options - ProtectionContainersClientRefreshOptions contains the optional parameters for the ProtectionContainersClient.Refresh
//     method.
func (client *ProtectionContainersClient) Refresh(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, options *ProtectionContainersClientRefreshOptions) (ProtectionContainersClientRefreshResponse, error) {
	var err error
	const operationName = "ProtectionContainersClient.Refresh"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.refreshCreateRequest(ctx, vaultName, resourceGroupName, fabricName, options)
	if err != nil {
		return ProtectionContainersClientRefreshResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionContainersClientRefreshResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionContainersClientRefreshResponse{}, err
	}
	return ProtectionContainersClientRefreshResponse{}, nil
}

// refreshCreateRequest creates the Refresh request.
func (client *ProtectionContainersClient) refreshCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, options *ProtectionContainersClientRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/refreshContainers"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginRegister - Registers the container with Recovery Services vault. This is an asynchronous operation. To track the operation
// status, use location header to call get latest status of the operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the container.
//   - containerName - Name of the container to be registered.
//   - parameters - Request body for operation
//   - options - ProtectionContainersClientBeginRegisterOptions contains the optional parameters for the ProtectionContainersClient.BeginRegister
//     method.
func (client *ProtectionContainersClient) BeginRegister(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters ProtectionContainerResource, options *ProtectionContainersClientBeginRegisterOptions) (*runtime.Poller[ProtectionContainersClientRegisterResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.register(ctx, vaultName, resourceGroupName, fabricName, containerName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProtectionContainersClientRegisterResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProtectionContainersClientRegisterResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Register - Registers the container with Recovery Services vault. This is an asynchronous operation. To track the operation
// status, use location header to call get latest status of the operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
func (client *ProtectionContainersClient) register(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters ProtectionContainerResource, options *ProtectionContainersClientBeginRegisterOptions) (*http.Response, error) {
	var err error
	const operationName = "ProtectionContainersClient.BeginRegister"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.registerCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// registerCreateRequest creates the Register request.
func (client *ProtectionContainersClient) registerCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters ProtectionContainerResource, _ *ProtectionContainersClientBeginRegisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Unregister - Unregisters the given container from your Recovery Services Vault. This is an asynchronous operation. To determine
// whether the backend service has finished processing the request, call Get Container
// Operation Result API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Name of the fabric where the container belongs.
//   - containerName - Name of the container which needs to be unregistered from the Recovery Services Vault.
//   - options - ProtectionContainersClientUnregisterOptions contains the optional parameters for the ProtectionContainersClient.Unregister
//     method.
func (client *ProtectionContainersClient) Unregister(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersClientUnregisterOptions) (ProtectionContainersClientUnregisterResponse, error) {
	var err error
	const operationName = "ProtectionContainersClient.Unregister"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unregisterCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, options)
	if err != nil {
		return ProtectionContainersClientUnregisterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionContainersClientUnregisterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionContainersClientUnregisterResponse{}, err
	}
	return ProtectionContainersClientUnregisterResponse{}, nil
}

// unregisterCreateRequest creates the Unregister request.
func (client *ProtectionContainersClient) unregisterCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, _ *ProtectionContainersClientUnregisterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
