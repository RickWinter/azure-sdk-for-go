//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/GenerateManifest.json
func ExampleClient_GenerateManifest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().GenerateManifest(ctx, "Microsoft.Contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceProviderManifest = armproviderhub.ResourceProviderManifest{
	// 	Capabilities: []*armproviderhub.ResourceProviderCapabilities{
	// 		{
	// 			Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 			QuotaID: to.Ptr("CSP_2015-05-01"),
	// 		},
	// 		{
	// 			Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 			QuotaID: to.Ptr("CSP_MG_2017-12-01"),
	// 	}},
	// 	CrossTenantTokenValidation: to.Ptr(armproviderhub.CrossTenantTokenValidationEnsureSecureValidation),
	// 	GlobalNotificationEndpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 		{
	// 			Enabled: to.Ptr(true),
	// 			EndpointURI: to.Ptr("https://notificationendpoint.com"),
	// 	}},
	// 	Management: &armproviderhub.ResourceProviderManifestManagement{
	// 		AuthorizationOwners: []*string{
	// 			to.Ptr("authorizationOwners-group")},
	// 			IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 			IncidentRoutingService: to.Ptr(""),
	// 			IncidentRoutingTeam: to.Ptr(""),
	// 			ManifestOwners: []*string{
	// 				to.Ptr("manifestOwners-group")},
	// 				ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
	// 			},
	// 			Metadata: map[string]any{
	// 				"onboardedVia": "ProviderHub",
	// 			},
	// 			Namespace: to.Ptr("microsoft.contoso"),
	// 			ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
	// 				{
	// 					ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 					RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 			}},
	// 			ProviderType: to.Ptr(armproviderhub.ResourceProviderType("Internal, Hidden")),
	// 			ProviderVersion: to.Ptr("2.0"),
	// 			ReRegisterSubscriptionMetadata: &armproviderhub.ResourceProviderManifestReRegisterSubscriptionMetadata{
	// 				ConcurrencyLimit: to.Ptr[int32](100),
	// 				Enabled: to.Ptr(true),
	// 			},
	// 			ResourceProviderAuthorizationRules: &armproviderhub.ResourceProviderAuthorizationRules{
	// 				AsyncOperationPollingRules: &armproviderhub.AsyncOperationPollingRules{
	// 					AuthorizationActions: []*string{
	// 						to.Ptr("Microsoft.Contoso/classicAdministrators/operationStatuses/read")},
	// 					},
	// 				},
	// 				ResourceTypes: []*armproviderhub.ResourceType{
	// 					{
	// 						Name: to.Ptr("Operations"),
	// 						AllowedUnauthorizedActions: []*string{
	// 							to.Ptr("microsoft.contoso/operations/read")},
	// 							AllowedUnauthorizedActionsExtensions: []*armproviderhub.AllowedUnauthorizedActionsExtension{
	// 								{
	// 									Action: to.Ptr("Microsoft.BizTalkServices/bizTalk/read"),
	// 									Intent: to.Ptr(armproviderhub.IntentDEFERREDACCESSCHECK),
	// 							}},
	// 							Endpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 								{
	// 									APIVersions: []*string{
	// 										to.Ptr("2020-01-01-preview")},
	// 										EndpointURI: to.Ptr("https://resource-endpoint.com/"),
	// 										Locations: []*string{
	// 											to.Ptr("")},
	// 											Timeout: to.Ptr("PT20S"),
	// 									}},
	// 									LinkedOperationRules: []*armproviderhub.LinkedOperationRule{
	// 									},
	// 									Notifications: []*armproviderhub.Notification{
	// 										{
	// 											NotificationType: to.Ptr(armproviderhub.NotificationTypeSubscriptionNotification),
	// 											SkipNotifications: to.Ptr(armproviderhub.SkipNotificationsDisabled),
	// 									}},
	// 									ResourceValidation: to.Ptr(armproviderhub.ResourceValidation("ReservedWords, ProfaneWords")),
	// 									RoutingType: to.Ptr(armproviderhub.RoutingType("ProxyOnly, Tenant")),
	// 								},
	// 								{
	// 									Name: to.Ptr("Locations"),
	// 									Endpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 										{
	// 											APIVersions: []*string{
	// 												to.Ptr("2020-01-01-preview")},
	// 												EndpointURI: to.Ptr("https://resource-endpoint.com/"),
	// 												Locations: []*string{
	// 													to.Ptr("")},
	// 													Timeout: to.Ptr("PT20S"),
	// 											}},
	// 											LinkedOperationRules: []*armproviderhub.LinkedOperationRule{
	// 											},
	// 											ResourceValidation: to.Ptr(armproviderhub.ResourceValidation("ReservedWords, ProfaneWords")),
	// 											RoutingType: to.Ptr(armproviderhub.RoutingTypeProxyOnly),
	// 										},
	// 										{
	// 											Name: to.Ptr("Locations/OperationStatuses"),
	// 											AdditionalOptions: to.Ptr(armproviderhub.AdditionalOptionsProtectedAsyncOperationPolling),
	// 											Endpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 												{
	// 													APIVersions: []*string{
	// 														to.Ptr("2020-01-01-preview")},
	// 														EndpointURI: to.Ptr("https://resource-endpoint.com/"),
	// 														Locations: []*string{
	// 															to.Ptr("")},
	// 															Timeout: to.Ptr("PT20S"),
	// 													}},
	// 													LinkedOperationRules: []*armproviderhub.LinkedOperationRule{
	// 													},
	// 													ResourceValidation: to.Ptr(armproviderhub.ResourceValidation("ReservedWords, ProfaneWords")),
	// 													RoutingType: to.Ptr(armproviderhub.RoutingType("ProxyOnly, LocationBased")),
	// 											}},
	// 											ServiceName: to.Ptr("root"),
	// 											Services: []*armproviderhub.ResourceProviderService{
	// 												{
	// 													ServiceName: to.Ptr("tags"),
	// 													Status: to.Ptr(armproviderhub.ServiceStatusInactive),
	// 											}},
	// 										}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/CheckinManifest.json
func ExampleClient_CheckinManifest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().CheckinManifest(ctx, "Microsoft.Contoso", armproviderhub.CheckinManifestParams{
		BaselineArmManifestLocation: to.Ptr("EastUS2EUAP"),
		Environment:                 to.Ptr("Prod"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckinManifestInfo = armproviderhub.CheckinManifestInfo{
	// 	IsCheckedIn: to.Ptr(false),
	// 	StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
	// }
}
