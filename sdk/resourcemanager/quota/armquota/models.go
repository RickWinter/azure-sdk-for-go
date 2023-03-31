//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armquota

import "time"

// ClientBeginCreateOrUpdateOptions contains the optional parameters for the Client.BeginCreateOrUpdate method.
type ClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientBeginUpdateOptions contains the optional parameters for the Client.BeginUpdate method.
type ClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClientListOptions contains the optional parameters for the Client.NewListPager method.
type ClientListOptions struct {
	// placeholder for future optional parameters
}

// CommonResourceProperties - Resource properties.
type CommonResourceProperties struct {
	// READ-ONLY; Resource ID
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type. Example: "Microsoft.Quota/quotas"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CreateGenericQuotaRequestParameters - Quota change requests information.
type CreateGenericQuotaRequestParameters struct {
	// Quota change requests.
	Value []*CurrentQuotaLimitBase `json:"value,omitempty"`
}

// CurrentQuotaLimitBase - Quota limit.
type CurrentQuotaLimitBase struct {
	// Quota properties for the specified resource, based on the API called, Quotas or Usages.
	Properties *Properties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CurrentUsagesBase - Resource usage.
type CurrentUsagesBase struct {
	// Usage properties for the specified resource.
	Properties *UsagesProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ExceptionResponse - Error.
type ExceptionResponse struct {
	// API error details.
	Error *ServiceError `json:"error,omitempty"`
}

// LimitJSONObjectClassification provides polymorphic access to related types.
// Call the interface's GetLimitJSONObject() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *LimitJSONObject, *LimitObject
type LimitJSONObjectClassification interface {
	// GetLimitJSONObject returns the LimitJSONObject content of the underlying type.
	GetLimitJSONObject() *LimitJSONObject
}

// LimitJSONObject - LimitJson abstract class.
type LimitJSONObject struct {
	// REQUIRED; The limit object type.
	LimitObjectType *LimitType `json:"limitObjectType,omitempty"`
}

// GetLimitJSONObject implements the LimitJSONObjectClassification interface for type LimitJSONObject.
func (l *LimitJSONObject) GetLimitJSONObject() *LimitJSONObject { return l }

// LimitObject - The resource quota limit value.
type LimitObject struct {
	// REQUIRED; The limit object type.
	LimitObjectType *LimitType `json:"limitObjectType,omitempty"`

	// REQUIRED; The quota/limit value
	Value *int32 `json:"value,omitempty"`

	// The quota or usages limit types.
	LimitType *QuotaLimitTypes `json:"limitType,omitempty"`
}

// GetLimitJSONObject implements the LimitJSONObjectClassification interface for type LimitObject.
func (l *LimitObject) GetLimitJSONObject() *LimitJSONObject {
	return &LimitJSONObject{
		LimitObjectType: l.LimitObjectType,
	}
}

// Limits - Quota limits.
type Limits struct {
	// The URI used to fetch the next page of quota limits. When there are no more pages, this string is null.
	NextLink *string `json:"nextLink,omitempty"`

	// List of quota limits.
	Value []*CurrentQuotaLimitBase `json:"value,omitempty"`
}

// LimitsResponse - Quota limits request response.
type LimitsResponse struct {
	// The URI used to fetch the next page of quota limits. When there are no more pages, this is null.
	NextLink *string `json:"nextLink,omitempty"`

	// List of quota limits with the quota request status.
	Value []*CurrentQuotaLimitBase `json:"value,omitempty"`
}

// OperationClientListOptions contains the optional parameters for the OperationClient.NewListPager method.
type OperationClientListOptions struct {
	// placeholder for future optional parameters
}

type OperationDisplay struct {
	// Operation description.
	Description *string `json:"description,omitempty"`

	// Operation name.
	Operation *string `json:"operation,omitempty"`

	// Provider name.
	Provider *string `json:"provider,omitempty"`

	// Resource name.
	Resource *string `json:"resource,omitempty"`
}

type OperationList struct {
	// URL to get the next page of items.
	NextLink *string              `json:"nextLink,omitempty"`
	Value    []*OperationResponse `json:"value,omitempty"`
}

type OperationResponse struct {
	Display *OperationDisplay `json:"display,omitempty"`
	Name    *string           `json:"name,omitempty"`
	Origin  *string           `json:"origin,omitempty"`
}

// Properties - Quota properties for the specified resource.
type Properties struct {
	// Resource quota limit properties.
	Limit LimitJSONObjectClassification `json:"limit,omitempty"`

	// Resource name provided by the resource provider. Use this property name when requesting quota.
	Name *ResourceName `json:"name,omitempty"`

	// Additional properties for the specific resource provider.
	Properties any `json:"properties,omitempty"`

	// Resource type name.
	ResourceType *string `json:"resourceType,omitempty"`

	// READ-ONLY; States if quota can be requested for this resource.
	IsQuotaApplicable *bool `json:"isQuotaApplicable,omitempty" azure:"ro"`

	// READ-ONLY; The time period over which the quota usage values are summarized. For example: *P1D (per one day) *PT1M (per
	// one minute) *PT1S (per one second). This parameter is optional because, for some resources
	// like compute, the period is irrelevant.
	QuotaPeriod *string `json:"quotaPeriod,omitempty" azure:"ro"`

	// READ-ONLY; The quota units, such as Count and Bytes. When requesting quota, use the unit value returned in the GET response
	// in the request body of your PUT operation.
	Unit *string `json:"unit,omitempty" azure:"ro"`
}

// RequestDetails - List of quota requests with details.
type RequestDetails struct {
	// Quota request details.
	Properties *RequestProperties `json:"properties,omitempty"`

	// READ-ONLY; Quota request ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Quota request name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type. "Microsoft.Quota/quotas".
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RequestDetailsList - Quota request information.
type RequestDetailsList struct {
	// The URI for fetching the next page of quota limits. When there are no more pages, this string is null.
	NextLink *string `json:"nextLink,omitempty"`

	// Quota request details.
	Value []*RequestDetails `json:"value,omitempty"`
}

// RequestOneResourceProperties - Quota request.
type RequestOneResourceProperties struct {
	// Error details of the quota request.
	Error *ServiceErrorDetail `json:"error,omitempty"`

	// Resource quota limit properties.
	Limit *LimitObject `json:"limit,omitempty"`

	// Resource name provided by the resource provider. Use this property name when requesting quota.
	Name *ResourceName `json:"name,omitempty"`

	// Additional properties for the specific resource provider.
	Properties any `json:"properties,omitempty"`

	// Resource type name.
	ResourceType *string `json:"resourceType,omitempty"`

	// The quota limit units, such as Count and Bytes. When requesting quota, use the unit value returned in the GET response
	// in the request body of your PUT operation.
	Unit *string `json:"unit,omitempty"`

	// READ-ONLY; Usage information for the current resource.
	CurrentValue *int32 `json:"currentValue,omitempty" azure:"ro"`

	// READ-ONLY; States if quota can be requested for this resource.
	IsQuotaApplicable *bool `json:"isQuotaApplicable,omitempty" azure:"ro"`

	// READ-ONLY; User-friendly status message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Quota request status.
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The time period over which the quota usage values are summarized. For example: *P1D (per one day) *PT1M (per
	// one minute) *PT1S (per one second). This parameter is optional because, for some resources
	// like compute, the period is irrelevant.
	QuotaPeriod *string `json:"quotaPeriod,omitempty" azure:"ro"`

	// READ-ONLY; Quota request submission time. The date conforms to the following ISO 8601 standard format: yyyy-MM-ddTHH:mm:ssZ.
	RequestSubmitTime *time.Time `json:"requestSubmitTime,omitempty" azure:"ro"`
}

// RequestOneResourceSubmitResponse - Quota request response.
type RequestOneResourceSubmitResponse struct {
	// Quota request details.
	Properties *RequestOneResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; Quota request ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the quota request.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type. "Microsoft.Quota/ServiceLimitRequests"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RequestProperties - Quota request properties.
type RequestProperties struct {
	// Error details of the quota request.
	Error *ServiceErrorDetail `json:"error,omitempty"`

	// Quota request details.
	Value []*SubRequest `json:"value,omitempty"`

	// READ-ONLY; User-friendly status message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The quota request status.
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The quota request submission time. The date conforms to the following format specified by the ISO 8601 standard:
	// yyyy-MM-ddTHH:mm:ssZ
	RequestSubmitTime *time.Time `json:"requestSubmitTime,omitempty" azure:"ro"`
}

// RequestStatusClientGetOptions contains the optional parameters for the RequestStatusClient.Get method.
type RequestStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// RequestStatusClientListOptions contains the optional parameters for the RequestStatusClient.NewListPager method.
type RequestStatusClientListOptions struct {
	// FIELD SUPPORTED OPERATORS
	// requestSubmitTime ge, le, eq, gt, lt
	// provisioningState eq {QuotaRequestState}
	// resourceName eq {resourceName}
	Filter *string
	// The Skiptoken parameter is used only if a previous operation returned a partial result. If a previous response contains
	// a nextLink element, its value includes a skiptoken parameter that specifies a
	// starting point to use for subsequent calls.
	Skiptoken *string
	// Number of records to return.
	Top *int32
}

// RequestStatusDetails - Quota request status details.
type RequestStatusDetails struct {
	// Resource quota limit properties.
	Limit *LimitObject `json:"limit,omitempty"`

	// Resource name provided by the resource provider. Use this property name when requesting quota.
	Name *ResourceName `json:"name,omitempty"`

	// Additional properties for the specific resource provider.
	Properties any `json:"properties,omitempty"`

	// Resource type name.
	ResourceType *string `json:"resourceType,omitempty"`

	// The quota limit units, such as Count and Bytes. When requesting quota, use the unit value returned in the GET response
	// in the request body of your PUT operation.
	Unit *string `json:"unit,omitempty"`

	// READ-ONLY; User-friendly message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Quota request status.
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The time period over which the quota usage values are summarized. For example: *P1D (per one day) *PT1M (per
	// one minute) *PT1S (per one second). This parameter is optional because, for some resources
	// like compute, the period is irrelevant.
	QuotaPeriod *string `json:"quotaPeriod,omitempty" azure:"ro"`
}

// RequestSubmitResponse - Quota request response.
type RequestSubmitResponse struct {
	// Quota request details.
	Properties *RequestProperties `json:"properties,omitempty"`

	// READ-ONLY; Quota request ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Quota request name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type. "Microsoft.Quota/quotas".
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RequestSubmitResponse202 - The quota request response with the quota request ID.
type RequestSubmitResponse202 struct {
	// Quota request status.
	Properties *RequestStatusDetails `json:"properties,omitempty"`

	// READ-ONLY; The quota request ID. To check the request status, use the id value in a Quota Request Status [https://docs.microsoft.com/en-us/rest/api/reserved-vm-instances/quotarequeststatus/get]
	// GET operation.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Operation ID.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceName - Name of the resource provided by the resource Provider. When requesting quota, use this property name.
type ResourceName struct {
	// Resource name.
	Value *string `json:"value,omitempty"`

	// READ-ONLY; Resource display name.
	LocalizedValue *string `json:"localizedValue,omitempty" azure:"ro"`
}

// ServiceError - API error details.
type ServiceError struct {
	// Error code.
	Code *string `json:"code,omitempty"`

	// Error message.
	Message *string `json:"message,omitempty"`

	// READ-ONLY; List of error details.
	Details []*ServiceErrorDetail `json:"details,omitempty" azure:"ro"`
}

// ServiceErrorDetail - Error details.
type ServiceErrorDetail struct {
	// READ-ONLY; Error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// SubRequest - Request property.
type SubRequest struct {
	// Resource quota limit properties.
	Limit LimitJSONObjectClassification `json:"limit,omitempty"`

	// Resource name.
	Name *ResourceName `json:"name,omitempty"`

	// Quota limit units, such as Count and Bytes. When requesting quota, use the unit value returned in the GET response in the
	// request body of your PUT operation.
	Unit *string `json:"unit,omitempty"`

	// READ-ONLY; User-friendly status message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The quota request status.
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Resource type for which the quota properties were requested.
	ResourceType *string `json:"resourceType,omitempty" azure:"ro"`

	// READ-ONLY; Quota request ID.
	SubRequestID *string `json:"subRequestId,omitempty" azure:"ro"`
}

// UsagesClientGetOptions contains the optional parameters for the UsagesClient.Get method.
type UsagesClientGetOptions struct {
	// placeholder for future optional parameters
}

// UsagesClientListOptions contains the optional parameters for the UsagesClient.NewListPager method.
type UsagesClientListOptions struct {
	// placeholder for future optional parameters
}

// UsagesLimits - Quota limits.
type UsagesLimits struct {
	// The URI used to fetch the next page of quota limits. When there are no more pages, this is null.
	NextLink *string `json:"nextLink,omitempty"`

	// List of quota limits.
	Value []*CurrentUsagesBase `json:"value,omitempty"`
}

// UsagesObject - The resource usages value.
type UsagesObject struct {
	// REQUIRED; The usages value.
	Value *int32 `json:"value,omitempty"`

	// The quota or usages limit types.
	UsagesType *UsagesTypes `json:"usagesType,omitempty"`
}

// UsagesProperties - Usage properties for the specified resource.
type UsagesProperties struct {
	// Resource name provided by the resource provider. Use this property name when requesting quota.
	Name *ResourceName `json:"name,omitempty"`

	// Additional properties for the specific resource provider.
	Properties any `json:"properties,omitempty"`

	// The name of the resource type.
	ResourceType *string `json:"resourceType,omitempty"`

	// The quota limit properties for this resource.
	Usages *UsagesObject `json:"usages,omitempty"`

	// READ-ONLY; States if quota can be requested for this resource.
	IsQuotaApplicable *bool `json:"isQuotaApplicable,omitempty" azure:"ro"`

	// READ-ONLY; The time period for the summary of the quota usage values. For example: *P1D (per one day) *PT1M (per one minute)
	// *PT1S (per one second). This parameter is optional because it is not relevant for all
	// resources such as compute.
	QuotaPeriod *string `json:"quotaPeriod,omitempty" azure:"ro"`

	// READ-ONLY; The units for the quota usage, such as Count and Bytes. When requesting quota, use the unit value returned in
	// the GET response in the request body of your PUT operation.
	Unit *string `json:"unit,omitempty" azure:"ro"`
}