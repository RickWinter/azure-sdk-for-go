// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armplanetarycomputer

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/planetarycomputer/armplanetarycomputer"
	moduleVersion = "v0.1.0"
)

// AutoGeneratedDomainNameLabelScope - The scope at which the auto-generated domain name label is generated and at which the
// resource name can be reused.
type AutoGeneratedDomainNameLabelScope string

const (
	// AutoGeneratedDomainNameLabelScopeNoReuse - The domain name label is randomly generated. The resource name cannot be reused
	// within the same region.
	AutoGeneratedDomainNameLabelScopeNoReuse AutoGeneratedDomainNameLabelScope = "NoReuse"
	// AutoGeneratedDomainNameLabelScopeResourceGroupReuse - The domain name label is deterministically generated using the resource
	// name, tenant id, subscription id, and resource group name. The resource name cannot be reused within the same resource
	// group.
	AutoGeneratedDomainNameLabelScopeResourceGroupReuse AutoGeneratedDomainNameLabelScope = "ResourceGroupReuse"
	// AutoGeneratedDomainNameLabelScopeSubscriptionReuse - The domain name label is deterministically generated using the resource
	// name, tenant id, and subscription id. The resource name cannot be reused within the same region and subscription.
	AutoGeneratedDomainNameLabelScopeSubscriptionReuse AutoGeneratedDomainNameLabelScope = "SubscriptionReuse"
	// AutoGeneratedDomainNameLabelScopeTenantReuse - The domain name label is deterministically generated using the resource
	// name and tenant id. The resource name cannot be reused within the same region and tenant.
	AutoGeneratedDomainNameLabelScopeTenantReuse AutoGeneratedDomainNameLabelScope = "TenantReuse"
)

// PossibleAutoGeneratedDomainNameLabelScopeValues returns the possible values for the AutoGeneratedDomainNameLabelScope const type.
func PossibleAutoGeneratedDomainNameLabelScopeValues() []AutoGeneratedDomainNameLabelScope {
	return []AutoGeneratedDomainNameLabelScope{
		AutoGeneratedDomainNameLabelScopeNoReuse,
		AutoGeneratedDomainNameLabelScopeResourceGroupReuse,
		AutoGeneratedDomainNameLabelScopeSubscriptionReuse,
		AutoGeneratedDomainNameLabelScopeTenantReuse,
	}
}

// CatalogTier - The Microsoft Planetary Computer Pro GeoCatalog tier
type CatalogTier string

const (
	// CatalogTierBasic - The basic tier that utilizes shared resources across catalog instances
	CatalogTierBasic CatalogTier = "Basic"
)

// PossibleCatalogTierValues returns the possible values for the CatalogTier const type.
func PossibleCatalogTierValues() []CatalogTier {
	return []CatalogTier{
		CatalogTierBasic,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	// ManagedServiceIdentityTypeNone - No managed identity.
	ManagedServiceIdentityTypeNone ManagedServiceIdentityType = "None"
	// ManagedServiceIdentityTypeSystemAssigned - System assigned managed identity.
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
	// ManagedServiceIdentityTypeSystemAssignedUserAssigned - System and user assigned managed identity.
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	// ManagedServiceIdentityTypeUserAssigned - User assigned managed identity.
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// ProvisioningState - The status of the current operation.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - The catalog request has been accepted.
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - The catalog is being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - The catalog is being provisioned.
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - The catalog is being updated.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}
