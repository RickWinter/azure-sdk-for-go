//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d0d3a9b4fe0fce880fded7a617e71f84406bacbd/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armnetworkcloud.OperationListResult{
		// 	Value: []*armnetworkcloud.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.NetworkCloud/clusters/Read"),
		// 			ActionType: to.Ptr(armnetworkcloud.ActionTypeInternal),
		// 			Display: &armnetworkcloud.OperationDisplay{
		// 				Description: to.Ptr("Reads Network Cloud cluster(s)"),
		// 				Operation: to.Ptr("Get/list Network Cloud cluster resources"),
		// 				Provider: to.Ptr("Microsoft Network Cloud"),
		// 				Resource: to.Ptr("Network Cloud cluster"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armnetworkcloud.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.NetworkCloud/clusters/Write"),
		// 			ActionType: to.Ptr(armnetworkcloud.ActionTypeInternal),
		// 			Display: &armnetworkcloud.OperationDisplay{
		// 				Description: to.Ptr("Writes Network Cloud cluster(s)"),
		// 				Operation: to.Ptr("Create/update Network Cloud cluster resources"),
		// 				Provider: to.Ptr("Microsoft Network Cloud"),
		// 				Resource: to.Ptr("Network Cloud cluster"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armnetworkcloud.OriginUserSystem),
		// 	}},
		// }
	}
}
