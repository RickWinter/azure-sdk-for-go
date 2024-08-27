//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListPredefinedClusterLibraries.json
func ExampleClusterLibrariesClient_NewListPager_listPredefinedClusterLibraries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterLibrariesClient().NewListPager("hiloResourceGroup", "clusterPool", "cluster", armhdinsightcontainers.CategoryPredefined, nil)
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
		// page.ClusterLibraryList = armhdinsightcontainers.ClusterLibraryList{
		// 	Value: []*armhdinsightcontainers.ClusterLibrary{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hiloResourceGroup/providers/Microsoft.HDInsight/clusterPools/clusterPool/clusters/cluster/libraries/pyspark"),
		// 			Properties: &armhdinsightcontainers.PyPiLibraryProperties{
		// 				Type: to.Ptr(armhdinsightcontainers.TypePypi),
		// 				Name: to.Ptr("pyspark"),
		// 				Version: to.Ptr("3.5.0"),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hiloResourceGroup/providers/Microsoft.HDInsight/clusterPools/clusterPool/clusters/cluster/libraries/spark-sql"),
		// 			Properties: &armhdinsightcontainers.MavenLibraryProperties{
		// 				Type: to.Ptr(armhdinsightcontainers.TypeMaven),
		// 				Name: to.Ptr("spark-sql"),
		// 				GroupID: to.Ptr("org.apache.spark"),
		// 				Version: to.Ptr("3.5.0_2.12"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListUserCustomClusterLibraries.json
func ExampleClusterLibrariesClient_NewListPager_listUserCustomClusterLibraries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterLibrariesClient().NewListPager("hiloResourceGroup", "clusterPool", "cluster", armhdinsightcontainers.CategoryCustom, nil)
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
		// page.ClusterLibraryList = armhdinsightcontainers.ClusterLibraryList{
		// 	Value: []*armhdinsightcontainers.ClusterLibrary{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hiloResourceGroup/providers/Microsoft.HDInsight/clusterPools/clusterPool/clusters/cluster/libraries/requests"),
		// 			Properties: &armhdinsightcontainers.PyPiLibraryProperties{
		// 				Type: to.Ptr(armhdinsightcontainers.TypePypi),
		// 				Message: to.Ptr(""),
		// 				Remarks: to.Ptr("PyPi packages."),
		// 				Status: to.Ptr(armhdinsightcontainers.StatusINSTALLED),
		// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00.000Z"); return t}()),
		// 				Name: to.Ptr("requests"),
		// 				Version: to.Ptr("2.31.0"),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hiloResourceGroup/providers/Microsoft.HDInsight/clusterPools/clusterPool/clusters/cluster/libraries/flink-connector-kafka"),
		// 			Properties: &armhdinsightcontainers.MavenLibraryProperties{
		// 				Type: to.Ptr(armhdinsightcontainers.TypeMaven),
		// 				Message: to.Ptr(""),
		// 				Remarks: to.Ptr("Maven packages."),
		// 				Status: to.Ptr(armhdinsightcontainers.StatusINSTALLING),
		// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00.000Z"); return t}()),
		// 				Name: to.Ptr("flink-connector-kafka"),
		// 				GroupID: to.Ptr("org.apache.flink"),
		// 				Version: to.Ptr("3.0.2-1.18"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/InstallNewClusterLibraries.json
func ExampleClusterLibrariesClient_BeginManageLibraries_installNewClusterLibraries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterLibrariesClient().BeginManageLibraries(ctx, "hiloResourceGroup", "clusterPool", "cluster", armhdinsightcontainers.ClusterLibraryManagementOperation{
		Properties: &armhdinsightcontainers.ClusterLibraryManagementOperationProperties{
			Action: to.Ptr(armhdinsightcontainers.LibraryManagementActionInstall),
			Libraries: []*armhdinsightcontainers.ClusterLibrary{
				{
					Properties: &armhdinsightcontainers.PyPiLibraryProperties{
						Type:    to.Ptr(armhdinsightcontainers.TypePypi),
						Remarks: to.Ptr("PyPi packages."),
						Name:    to.Ptr("requests"),
						Version: to.Ptr("2.31.0"),
					},
				},
				{
					Properties: &armhdinsightcontainers.MavenLibraryProperties{
						Type:    to.Ptr(armhdinsightcontainers.TypeMaven),
						Remarks: to.Ptr("Maven packages."),
						Name:    to.Ptr("flink-connector-kafka"),
						GroupID: to.Ptr("org.apache.flink"),
						Version: to.Ptr("3.0.2-1.18"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/UninstallExistingClusterLibraries.json
func ExampleClusterLibrariesClient_BeginManageLibraries_uninstallExistingClusterLibraries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterLibrariesClient().BeginManageLibraries(ctx, "hiloResourceGroup", "clusterPool", "cluster", armhdinsightcontainers.ClusterLibraryManagementOperation{
		Properties: &armhdinsightcontainers.ClusterLibraryManagementOperationProperties{
			Action: to.Ptr(armhdinsightcontainers.LibraryManagementActionUninstall),
			Libraries: []*armhdinsightcontainers.ClusterLibrary{
				{
					Properties: &armhdinsightcontainers.PyPiLibraryProperties{
						Type: to.Ptr(armhdinsightcontainers.TypePypi),
						Name: to.Ptr("tensorflow"),
					},
				},
				{
					Properties: &armhdinsightcontainers.MavenLibraryProperties{
						Type:    to.Ptr(armhdinsightcontainers.TypeMaven),
						Name:    to.Ptr("flink-connector-hudi"),
						GroupID: to.Ptr("org.apache.flink"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}