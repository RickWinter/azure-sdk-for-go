//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/NetworkSecurityPerimeterConfigurationList.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().NewListPager("res4410", "cosmosTest", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.NetworkSecurityPerimeterConfigurationListResult = armcosmos.NetworkSecurityPerimeterConfigurationListResult{
		// 	Value: []*armcosmos.NetworkSecurityPerimeterConfiguration{
		// 		{
		// 			Name: to.Ptr("dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/networkSecurityPerimeterConfigurations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/res4410/providers/Microsoft.DocumentDB/databaseAccounts/cosmosTest/networkSecurityPerimeterConfigurations/dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1"),
		// 			Properties: &armcosmos.NetworkSecurityPerimeterConfigurationProperties{
		// 				NetworkSecurityPerimeter: &armcosmos.NetworkSecurityPerimeter{
		// 					ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/res4794/providers/Microsoft.Network/networkSecurityPerimeters/nsp1"),
		// 					Location: to.Ptr("East US"),
		// 					PerimeterGUID: to.Ptr("ce2d5953-5c15-40ca-9d51-cc3f4a63b0f5"),
		// 				},
		// 				Profile: &armcosmos.NetworkSecurityProfile{
		// 					Name: to.Ptr("profile1"),
		// 					AccessRules: []*armcosmos.AccessRule{
		// 						{
		// 							Name: to.Ptr("inVpnRule"),
		// 							Properties: &armcosmos.AccessRuleProperties{
		// 								AddressPrefixes: []*string{
		// 									to.Ptr("148.0.0.0/8"),
		// 									to.Ptr("152.4.6.0/24")},
		// 									Direction: to.Ptr(armcosmos.AccessRuleDirectionInbound),
		// 								},
		// 						}},
		// 						AccessRulesVersion: to.Ptr[int32](10),
		// 						DiagnosticSettingsVersion: to.Ptr[int32](5),
		// 						EnabledLogCategories: []*string{
		// 							to.Ptr("NspPublicInboundPerimeterRulesAllowed"),
		// 							to.Ptr("NspPublicInboundPerimeterRulesDenied")},
		// 						},
		// 						ProvisioningState: to.Ptr(armcosmos.NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
		// 						ResourceAssociation: &armcosmos.ResourceAssociation{
		// 							Name: to.Ptr("association1"),
		// 							AccessMode: to.Ptr(armcosmos.ResourceAssociationAccessModeEnforced),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/NetworkSecurityPerimeterConfigurationGet.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().Get(ctx, "res4410", "cosmosTest", "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSecurityPerimeterConfiguration = armcosmos.NetworkSecurityPerimeterConfiguration{
	// 	Name: to.Ptr("dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/networkSecurityPerimeterConfigurations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/res4410/providers/Microsoft.DocumentDB/databaseAccounts/cosmosTest/networkSecurityPerimeterConfigurations/dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1"),
	// 	Properties: &armcosmos.NetworkSecurityPerimeterConfigurationProperties{
	// 		NetworkSecurityPerimeter: &armcosmos.NetworkSecurityPerimeter{
	// 			ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/res4794/providers/Microsoft.Network/networkSecurityPerimeters/nsp1"),
	// 			Location: to.Ptr("East US"),
	// 			PerimeterGUID: to.Ptr("ce2d5953-5c15-40ca-9d51-cc3f4a63b0f5"),
	// 		},
	// 		Profile: &armcosmos.NetworkSecurityProfile{
	// 			Name: to.Ptr("profile1"),
	// 			AccessRules: []*armcosmos.AccessRule{
	// 				{
	// 					Name: to.Ptr("allowedSubscriptions"),
	// 					Properties: &armcosmos.AccessRuleProperties{
	// 						Direction: to.Ptr(armcosmos.AccessRuleDirectionInbound),
	// 						Subscriptions: []*armcosmos.AccessRulePropertiesSubscriptionsItem{
	// 							{
	// 								ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774"),
	// 						}},
	// 					},
	// 			}},
	// 			AccessRulesVersion: to.Ptr[int32](10),
	// 			DiagnosticSettingsVersion: to.Ptr[int32](5),
	// 			EnabledLogCategories: []*string{
	// 				to.Ptr("NspPublicInboundPerimeterRulesAllowed"),
	// 				to.Ptr("NspPublicInboundPerimeterRulesDenied")},
	// 			},
	// 			ProvisioningState: to.Ptr(armcosmos.NetworkSecurityPerimeterConfigurationProvisioningStateSucceeded),
	// 			ResourceAssociation: &armcosmos.ResourceAssociation{
	// 				Name: to.Ptr("resourceAssociation1"),
	// 				AccessMode: to.Ptr(armcosmos.ResourceAssociationAccessModeEnforced),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/NetworkSecurityPerimeterConfigurationReconcile.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_BeginReconcile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().BeginReconcile(ctx, "res4410", "sto8607", "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
