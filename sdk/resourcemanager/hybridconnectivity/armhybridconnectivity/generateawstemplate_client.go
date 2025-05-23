// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhybridconnectivity

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

// GenerateAwsTemplateClient contains the methods for the GenerateAwsTemplate group.
// Don't use this type directly, use NewGenerateAwsTemplateClient() instead.
type GenerateAwsTemplateClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGenerateAwsTemplateClient creates a new instance of GenerateAwsTemplateClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGenerateAwsTemplateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GenerateAwsTemplateClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GenerateAwsTemplateClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Post - Retrieve AWS Cloud Formation template
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01
//   - generateAwsTemplateRequest - ConnectorId and SolutionTypes and their properties to Generate AWS CFT Template.
//   - options - GenerateAwsTemplateClientPostOptions contains the optional parameters for the GenerateAwsTemplateClient.Post
//     method.
func (client *GenerateAwsTemplateClient) Post(ctx context.Context, generateAwsTemplateRequest GenerateAwsTemplateRequest, options *GenerateAwsTemplateClientPostOptions) (GenerateAwsTemplateClientPostResponse, error) {
	var err error
	const operationName = "GenerateAwsTemplateClient.Post"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.postCreateRequest(ctx, generateAwsTemplateRequest, options)
	if err != nil {
		return GenerateAwsTemplateClientPostResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GenerateAwsTemplateClientPostResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GenerateAwsTemplateClientPostResponse{}, err
	}
	resp, err := client.postHandleResponse(httpResp)
	return resp, err
}

// postCreateRequest creates the Post request.
func (client *GenerateAwsTemplateClient) postCreateRequest(ctx context.Context, generateAwsTemplateRequest GenerateAwsTemplateRequest, _ *GenerateAwsTemplateClientPostOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridConnectivity/generateAwsTemplate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, generateAwsTemplateRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// postHandleResponse handles the Post response.
func (client *GenerateAwsTemplateClient) postHandleResponse(resp *http.Response) (GenerateAwsTemplateClientPostResponse, error) {
	result := GenerateAwsTemplateClientPostResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PostResponse); err != nil {
		return GenerateAwsTemplateClientPostResponse{}, err
	}
	return result, nil
}
