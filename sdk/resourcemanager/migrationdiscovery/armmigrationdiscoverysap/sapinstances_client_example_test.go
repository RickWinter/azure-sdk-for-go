//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmigrationdiscoverysap_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscoverysap"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/475747ff6322e9bf180b8911d871561b264379c3/specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_List.json
func ExampleSapInstancesClient_NewListBySapDiscoverySitePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationdiscoverysap.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSapInstancesClient().NewListBySapDiscoverySitePager("test-rg", "SampleSite", nil)
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
		// page.SAPInstanceListResult = armmigrationdiscoverysap.SAPInstanceListResult{
		// 	Value: []*armmigrationdiscoverysap.SAPInstance{
		// 		{
		// 			Name: to.Ptr("MPP_MPP"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapDiscoverySites/SampleSite/sapInstances/MPP_MPP"),
		// 			SystemData: &armmigrationdiscoverysap.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("UserName"),
		// 				CreatedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("UserName"),
		// 				LastModifiedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"property1": to.Ptr("value1"),
		// 				"property2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmigrationdiscoverysap.SAPInstanceProperties{
		// 				Application: to.Ptr("HR"),
		// 				Environment: to.Ptr(armmigrationdiscoverysap.SapInstanceEnvironmentProduction),
		// 				LandscapeSid: to.Ptr("MPP"),
		// 				ProvisioningState: to.Ptr(armmigrationdiscoverysap.ProvisioningStateSucceeded),
		// 				SystemSid: to.Ptr("MPP"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MPP_MPP_DR"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapDiscoverySites/SampleSite/sapInstances/MPP_MPP_DR"),
		// 			SystemData: &armmigrationdiscoverysap.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("UserName"),
		// 				CreatedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("UserName"),
		// 				LastModifiedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"property1": to.Ptr("value1"),
		// 				"property2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmigrationdiscoverysap.SAPInstanceProperties{
		// 				Application: to.Ptr("HR"),
		// 				Environment: to.Ptr(armmigrationdiscoverysap.SapInstanceEnvironmentDisasterRecovery),
		// 				LandscapeSid: to.Ptr("MPP"),
		// 				ProvisioningState: to.Ptr(armmigrationdiscoverysap.ProvisioningStateSucceeded),
		// 				SystemSid: to.Ptr("MPP"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MPP_MPD"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapDiscoverySites/SampleSite/sapInstances/MPP_MPD"),
		// 			SystemData: &armmigrationdiscoverysap.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("UserName"),
		// 				CreatedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("UserName"),
		// 				LastModifiedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"property1": to.Ptr("value1"),
		// 				"property2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmigrationdiscoverysap.SAPInstanceProperties{
		// 				Application: to.Ptr("HR"),
		// 				Environment: to.Ptr(armmigrationdiscoverysap.SapInstanceEnvironmentDevelopment),
		// 				LandscapeSid: to.Ptr("MPP"),
		// 				ProvisioningState: to.Ptr(armmigrationdiscoverysap.ProvisioningStateSucceeded),
		// 				SystemSid: to.Ptr("MPD"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/475747ff6322e9bf180b8911d871561b264379c3/specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_Get.json
func ExampleSapInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationdiscoverysap.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSapInstancesClient().Get(ctx, "test-rg", "SampleSite", "MPP_MPP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPInstance = armmigrationdiscoverysap.SAPInstance{
	// 	Name: to.Ptr("MPP_MPP"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapDiscoverySites/SampleSite/sapInstances/MPP_MPP"),
	// 	SystemData: &armmigrationdiscoverysap.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("UserName"),
	// 		CreatedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("UserName"),
	// 		LastModifiedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"property1": to.Ptr("value1"),
	// 		"property2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armmigrationdiscoverysap.SAPInstanceProperties{
	// 		Application: to.Ptr("HR"),
	// 		Environment: to.Ptr(armmigrationdiscoverysap.SapInstanceEnvironmentProduction),
	// 		LandscapeSid: to.Ptr("MPP"),
	// 		ProvisioningState: to.Ptr(armmigrationdiscoverysap.ProvisioningStateSucceeded),
	// 		SystemSid: to.Ptr("MPP"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/475747ff6322e9bf180b8911d871561b264379c3/specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_Create.json
func ExampleSapInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationdiscoverysap.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSapInstancesClient().BeginCreate(ctx, "test-rg", "SampleSite", "MPP_MPP", armmigrationdiscoverysap.SAPInstance{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"property1": to.Ptr("value1"),
			"property2": to.Ptr("value2"),
		},
		Properties: &armmigrationdiscoverysap.SAPInstanceProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPInstance = armmigrationdiscoverysap.SAPInstance{
	// 	Name: to.Ptr("MPP_MPP"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapDiscoverySites/SampleSite/sapInstances/MPP_MPP"),
	// 	SystemData: &armmigrationdiscoverysap.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("UserName"),
	// 		CreatedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("UserName"),
	// 		LastModifiedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"property1": to.Ptr("value1"),
	// 		"property2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armmigrationdiscoverysap.SAPInstanceProperties{
	// 		Application: to.Ptr("HR"),
	// 		Environment: to.Ptr(armmigrationdiscoverysap.SapInstanceEnvironmentProduction),
	// 		LandscapeSid: to.Ptr("MPP"),
	// 		ProvisioningState: to.Ptr(armmigrationdiscoverysap.ProvisioningStateSucceeded),
	// 		SystemSid: to.Ptr("MPP"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/475747ff6322e9bf180b8911d871561b264379c3/specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_Update.json
func ExampleSapInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationdiscoverysap.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSapInstancesClient().Update(ctx, "test-rg", "SampleSite", "MPP_MPP", armmigrationdiscoverysap.SAPInstanceTagsUpdate{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPInstance = armmigrationdiscoverysap.SAPInstance{
	// 	Name: to.Ptr("MPP_MPP"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapDiscoverySites/SampleSite/sapInstances/MPP_MPP"),
	// 	SystemData: &armmigrationdiscoverysap.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("UserName"),
	// 		CreatedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("UserName"),
	// 		LastModifiedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armmigrationdiscoverysap.SAPInstanceProperties{
	// 		Environment: to.Ptr(armmigrationdiscoverysap.SapInstanceEnvironmentProduction),
	// 		LandscapeSid: to.Ptr("MPP"),
	// 		ProvisioningState: to.Ptr(armmigrationdiscoverysap.ProvisioningStateSucceeded),
	// 		SystemSid: to.Ptr("MPP"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/475747ff6322e9bf180b8911d871561b264379c3/specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_Delete.json
func ExampleSapInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationdiscoverysap.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSapInstancesClient().BeginDelete(ctx, "test-rg", "SampleSite", "MPP_MPP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}