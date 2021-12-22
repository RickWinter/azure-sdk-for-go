//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ProtectionContainersClient contains the methods for the ProtectionContainers group.
// Don't use this type directly, use NewProtectionContainersClient() instead.
type ProtectionContainersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewProtectionContainersClient creates a new instance of ProtectionContainersClient with the specified values.
func NewProtectionContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProtectionContainersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ProtectionContainersClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets details of the specific container registered to your Recovery Services Vault.
// If the operation fails it returns the *CloudError error type.
func (client *ProtectionContainersClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersGetOptions) (ProtectionContainersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, options)
	if err != nil {
		return ProtectionContainersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProtectionContainersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProtectionContainersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProtectionContainersClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProtectionContainersClient) getHandleResponse(resp *http.Response) (ProtectionContainersGetResponse, error) {
	result := ProtectionContainersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerResource); err != nil {
		return ProtectionContainersGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ProtectionContainersClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Inquire - This is an async operation and the results should be tracked using location header or Azure-async-url.
// If the operation fails it returns the *CloudError error type.
func (client *ProtectionContainersClient) Inquire(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersInquireOptions) (ProtectionContainersInquireResponse, error) {
	req, err := client.inquireCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, options)
	if err != nil {
		return ProtectionContainersInquireResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProtectionContainersInquireResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ProtectionContainersInquireResponse{}, client.inquireHandleError(resp)
	}
	return ProtectionContainersInquireResponse{RawResponse: resp}, nil
}

// inquireCreateRequest creates the Inquire request.
func (client *ProtectionContainersClient) inquireCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersInquireOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// inquireHandleError handles the Inquire error response.
func (client *ProtectionContainersClient) inquireHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Refresh - Discovers all the containers in the subscription that can be backed up to Recovery Services Vault. This is an asynchronous operation. To know
// the status of the operation, call
// GetRefreshOperationResult API.
// If the operation fails it returns the *CloudError error type.
func (client *ProtectionContainersClient) Refresh(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, options *ProtectionContainersRefreshOptions) (ProtectionContainersRefreshResponse, error) {
	req, err := client.refreshCreateRequest(ctx, vaultName, resourceGroupName, fabricName, options)
	if err != nil {
		return ProtectionContainersRefreshResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProtectionContainersRefreshResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ProtectionContainersRefreshResponse{}, client.refreshHandleError(resp)
	}
	return ProtectionContainersRefreshResponse{RawResponse: resp}, nil
}

// refreshCreateRequest creates the Refresh request.
func (client *ProtectionContainersClient) refreshCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, options *ProtectionContainersRefreshOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// refreshHandleError handles the Refresh error response.
func (client *ProtectionContainersClient) refreshHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Register - Registers the container with Recovery Services vault. This is an asynchronous operation. To track the operation status, use location header
// to call get latest status of the operation.
// If the operation fails it returns the *CloudError error type.
func (client *ProtectionContainersClient) Register(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters ProtectionContainerResource, options *ProtectionContainersRegisterOptions) (ProtectionContainersRegisterResponse, error) {
	req, err := client.registerCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, parameters, options)
	if err != nil {
		return ProtectionContainersRegisterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProtectionContainersRegisterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return ProtectionContainersRegisterResponse{}, client.registerHandleError(resp)
	}
	return client.registerHandleResponse(resp)
}

// registerCreateRequest creates the Register request.
func (client *ProtectionContainersClient) registerCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters ProtectionContainerResource, options *ProtectionContainersRegisterOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// registerHandleResponse handles the Register response.
func (client *ProtectionContainersClient) registerHandleResponse(resp *http.Response) (ProtectionContainersRegisterResponse, error) {
	result := ProtectionContainersRegisterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerResource); err != nil {
		return ProtectionContainersRegisterResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// registerHandleError handles the Register error response.
func (client *ProtectionContainersClient) registerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Unregister - Unregisters the given container from your Recovery Services Vault. This is an asynchronous operation. To determine whether the backend service
// has finished processing the request, call Get Container
// Operation Result API.
// If the operation fails it returns the *CloudError error type.
func (client *ProtectionContainersClient) Unregister(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersUnregisterOptions) (ProtectionContainersUnregisterResponse, error) {
	req, err := client.unregisterCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, options)
	if err != nil {
		return ProtectionContainersUnregisterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProtectionContainersUnregisterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return ProtectionContainersUnregisterResponse{}, client.unregisterHandleError(resp)
	}
	return ProtectionContainersUnregisterResponse{RawResponse: resp}, nil
}

// unregisterCreateRequest creates the Unregister request.
func (client *ProtectionContainersClient) unregisterCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *ProtectionContainersUnregisterOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// unregisterHandleError handles the Unregister error response.
func (client *ProtectionContainersClient) unregisterHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
