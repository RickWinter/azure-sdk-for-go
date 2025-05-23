//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListWorkspaceNamedValues.json
func ExampleWorkspaceNamedValueClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceNamedValueClient().NewListByServicePager("rg1", "apimService1", "wks1", &armapimanagement.WorkspaceNamedValueClientListByServiceOptions{Filter: nil,
		Top:                     nil,
		Skip:                    nil,
		IsKeyVaultRefreshFailed: nil,
	})
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
		// page.NamedValueCollection = armapimanagement.NamedValueCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.NamedValueContract{
		// 		{
		// 			Name: to.Ptr("592f1174cc83890dc4f32686"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/592f1174cc83890dc4f32686"),
		// 			Properties: &armapimanagement.NamedValueContractProperties{
		// 				Secret: to.Ptr(false),
		// 				DisplayName: to.Ptr("Logger-Credentials-592f1174cc83890dc4f32687"),
		// 				Value: to.Ptr("propValue"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testprop6"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testprop6"),
		// 			Properties: &armapimanagement.NamedValueContractProperties{
		// 				Secret: to.Ptr(true),
		// 				Tags: []*string{
		// 					to.Ptr("foo"),
		// 					to.Ptr("bar")},
		// 					DisplayName: to.Ptr("prop6namekv"),
		// 					KeyVault: &armapimanagement.KeyVaultContractProperties{
		// 						IdentityClientID: to.Ptr("2d2df842-44d8-4885-8dec-77cc1a984a31"),
		// 						SecretIdentifier: to.Ptr("https://contoso.vault.azure.net/secrets/aadSecret"),
		// 						LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
		// 							Code: to.Ptr("Success"),
		// 							TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-11T00:54:31.802Z"); return t}()),
		// 						},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementHeadWorkspaceNamedValue.json
func ExampleWorkspaceNamedValueClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspaceNamedValueClient().GetEntityTag(ctx, "rg1", "apimService1", "wks1", "testarmTemplateproperties2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetWorkspaceNamedValue.json
func ExampleWorkspaceNamedValueClient_Get_apiManagementGetWorkspaceNamedValue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceNamedValueClient().Get(ctx, "rg1", "apimService1", "wks1", "testarmTemplateproperties2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Name: to.Ptr("testarmTemplateproperties2"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testarmTemplateproperties2"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(false),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar")},
	// 			DisplayName: to.Ptr("propName"),
	// 			Value: to.Ptr("propValue"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetWorkspaceNamedValueWithKeyVault.json
func ExampleWorkspaceNamedValueClient_Get_apiManagementGetWorkspaceNamedValueWithKeyVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceNamedValueClient().Get(ctx, "rg1", "apimService1", "wks1", "testprop6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Name: to.Ptr("testprop6"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testarmTemplateproperties2"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(true),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar")},
	// 			DisplayName: to.Ptr("prop6namekv"),
	// 			KeyVault: &armapimanagement.KeyVaultContractProperties{
	// 				IdentityClientID: to.Ptr("2d2df842-44d8-4885-8dec-77cc1a984a31"),
	// 				SecretIdentifier: to.Ptr("https://rpbvtkeyvaultintegration.vault-int.azure-int.net/secrets/msitestingCert"),
	// 				LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
	// 					Code: to.Ptr("Success"),
	// 					TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-11T00:54:31.802Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceNamedValue.json
func ExampleWorkspaceNamedValueClient_BeginCreateOrUpdate_apiManagementCreateWorkspaceNamedValue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceNamedValueClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "wks1", "testprop2", armapimanagement.NamedValueCreateContract{
		Properties: &armapimanagement.NamedValueCreateContractProperties{
			Secret: to.Ptr(false),
			Tags: []*string{
				to.Ptr("foo"),
				to.Ptr("bar")},
			DisplayName: to.Ptr("prop3name"),
			Value:       to.Ptr("propValue"),
		},
	}, &armapimanagement.WorkspaceNamedValueClientBeginCreateOrUpdateOptions{IfMatch: nil})
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
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Name: to.Ptr("testprop2"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testprop2"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(false),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar")},
	// 			DisplayName: to.Ptr("prop3name"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Value: to.Ptr("propValue"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceNamedValueWithKeyVault.json
func ExampleWorkspaceNamedValueClient_BeginCreateOrUpdate_apiManagementCreateWorkspaceNamedValueWithKeyVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceNamedValueClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "wks1", "testprop6", armapimanagement.NamedValueCreateContract{
		Properties: &armapimanagement.NamedValueCreateContractProperties{
			Secret: to.Ptr(true),
			Tags: []*string{
				to.Ptr("foo"),
				to.Ptr("bar")},
			DisplayName: to.Ptr("prop6namekv"),
			KeyVault: &armapimanagement.KeyVaultContractCreateProperties{
				IdentityClientID: to.Ptr("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"),
				SecretIdentifier: to.Ptr("https://contoso.vault.azure.net/secrets/aadSecret"),
			},
		},
	}, &armapimanagement.WorkspaceNamedValueClientBeginCreateOrUpdateOptions{IfMatch: nil})
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
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Name: to.Ptr("testprop6"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testprop6"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(true),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar")},
	// 			DisplayName: to.Ptr("prop6namekv"),
	// 			KeyVault: &armapimanagement.KeyVaultContractProperties{
	// 				IdentityClientID: to.Ptr("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"),
	// 				SecretIdentifier: to.Ptr("https://contoso.vault.azure.net/secrets/aadSecret"),
	// 				LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
	// 					Code: to.Ptr("Success"),
	// 					TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-11T00:54:31.802Z"); return t}()),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateWorkspaceNamedValue.json
func ExampleWorkspaceNamedValueClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceNamedValueClient().BeginUpdate(ctx, "rg1", "apimService1", "wks1", "testprop2", "*", armapimanagement.NamedValueUpdateParameters{
		Properties: &armapimanagement.NamedValueUpdateParameterProperties{
			Secret: to.Ptr(false),
			Tags: []*string{
				to.Ptr("foo"),
				to.Ptr("bar2")},
			DisplayName: to.Ptr("prop3name"),
			Value:       to.Ptr("propValue"),
		},
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
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/namedValues"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testprop2"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(false),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar2")},
	// 			DisplayName: to.Ptr("prop3name"),
	// 			Value: to.Ptr("propValue"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteWorkspaceNamedValue.json
func ExampleWorkspaceNamedValueClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkspaceNamedValueClient().Delete(ctx, "rg1", "apimService1", "wks1", "testprop2", "*", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementWorkspaceNamedValueListValue.json
func ExampleWorkspaceNamedValueClient_ListValue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceNamedValueClient().ListValue(ctx, "rg1", "apimService1", "wks1", "testarmTemplateproperties2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamedValueSecretContract = armapimanagement.NamedValueSecretContract{
	// 	Value: to.Ptr("propValue"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementRefreshWorkspaceNamedValue.json
func ExampleWorkspaceNamedValueClient_BeginRefreshSecret() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceNamedValueClient().BeginRefreshSecret(ctx, "rg1", "apimService1", "wks1", "testprop2", nil)
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
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Name: to.Ptr("testprop6"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testprop6"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(true),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar")},
	// 			DisplayName: to.Ptr("prop6namekv"),
	// 			KeyVault: &armapimanagement.KeyVaultContractProperties{
	// 				IdentityClientID: to.Ptr("2d2df842-44d8-4885-8dec-77cc1a984a31"),
	// 				SecretIdentifier: to.Ptr("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
	// 				LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
	// 					Code: to.Ptr("Success"),
	// 					TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-11T00:54:31.802Z"); return t}()),
	// 				},
	// 			},
	// 		},
	// 	}
}
