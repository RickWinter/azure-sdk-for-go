// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armonlineexperimentation_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/onlineexperimentation/armonlineexperimentation"
	"log"
)

// Generated from example definition: 2025-05-31-preview/OnlineExperimentationWorkspaces_OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armonlineexperimentation.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armonlineexperimentation.OperationsClientListResponse{
		// 	OperationListResult: armonlineexperimentation.OperationListResult{
		// 		Value: []*armonlineexperimentation.Operation{
		// 			{
		// 				Name: to.Ptr("aaaaaaaaa"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armonlineexperimentation.OperationDisplay{
		// 					Provider: to.Ptr("aaaaaaaa"),
		// 					Resource: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 					Operation: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 					Description: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				Origin: to.Ptr(armonlineexperimentation.OriginUser),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/providers/Microsoft.OnlineExperimentation/operations?api-version=2025-05-31-preview&$skiptoken=aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 	},
		// }
	}
}
