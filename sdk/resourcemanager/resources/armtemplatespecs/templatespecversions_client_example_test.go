//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtemplatespecs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/TemplateSpecVersionsCreate.json
func ExampleTemplateSpecVersionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTemplateSpecVersionsClient().CreateOrUpdate(ctx, "templateSpecRG", "simpleTemplateSpec", "v1.0", armtemplatespecs.TemplateSpecVersion{
		Location: to.Ptr("eastus"),
		Properties: &armtemplatespecs.TemplateSpecVersionProperties{
			Description: to.Ptr("This is version v1.0 of our template content"),
			MainTemplate: map[string]any{
				"$schema":        "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
				"contentVersion": "1.0.0.0",
				"parameters":     map[string]any{},
				"resources":      []any{},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TemplateSpecVersion = armtemplatespecs.TemplateSpecVersion{
	// 	Name: to.Ptr("v1.0"),
	// 	Type: to.Ptr("Microsoft.Resources/templateSpecs/versions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/templateSpecRG/providers/Microsoft.Resources/templateSpecs/simpleTemplateSpec/versions/v1.0"),
	// 	SystemData: &armtemplatespecs.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.1974346Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armtemplatespecs.TemplateSpecVersionProperties{
	// 		Description: to.Ptr("This is version v1.0 of our template content"),
	// 		MainTemplate: map[string]any{
	// 			"$schema": "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
	// 			"contentVersion": "1.0.0.0",
	// 			"parameters":map[string]any{
	// 			},
	// 			"resources":[]any{
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/TemplateSpecVersionsPatch.json
func ExampleTemplateSpecVersionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTemplateSpecVersionsClient().Update(ctx, "templateSpecRG", "simpleTemplateSpec", "v1.0", &armtemplatespecs.TemplateSpecVersionsClientUpdateOptions{TemplateSpecVersionUpdateModel: &armtemplatespecs.TemplateSpecVersionUpdateModel{
		Tags: map[string]*string{
			"myTag": to.Ptr("My Value"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TemplateSpecVersion = armtemplatespecs.TemplateSpecVersion{
	// 	Name: to.Ptr("v1.0"),
	// 	Type: to.Ptr("Microsoft.Resources/templateSpecs/versions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/templateSpecRG/providers/Microsoft.Resources/templateSpecs/simpleTemplateSpec/versions/v1.0"),
	// 	SystemData: &armtemplatespecs.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.1974346Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armtemplatespecs.TemplateSpecVersionProperties{
	// 		Description: to.Ptr("This is version v1.0 of our template content"),
	// 		MainTemplate: map[string]any{
	// 			"$schema": "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
	// 			"contentVersion": "1.0.0.0",
	// 			"parameters":map[string]any{
	// 			},
	// 			"resources":[]any{
	// 			},
	// 		},
	// 	},
	// 	Tags: map[string]*string{
	// 		"myTag": to.Ptr("My Value"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/TemplateSpecVersionsGet.json
func ExampleTemplateSpecVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTemplateSpecVersionsClient().Get(ctx, "templateSpecRG", "simpleTemplateSpec", "v1.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TemplateSpecVersion = armtemplatespecs.TemplateSpecVersion{
	// 	Name: to.Ptr("v1.0"),
	// 	Type: to.Ptr("Microsoft.Resources/templateSpecs/versions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/templateSpecRG/providers/Microsoft.Resources/templateSpecs/simpleTemplateSpec/versions/v1.0"),
	// 	SystemData: &armtemplatespecs.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.1974346Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armtemplatespecs.TemplateSpecVersionProperties{
	// 		Description: to.Ptr("This is version v1.0 of our template content"),
	// 		MainTemplate: map[string]any{
	// 			"$schema": "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
	// 			"contentVersion": "1.0.0.0",
	// 			"parameters":map[string]any{
	// 			},
	// 			"resources":[]any{
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/TemplateSpecVersionsDelete.json
func ExampleTemplateSpecVersionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTemplateSpecVersionsClient().Delete(ctx, "templateSpecRG", "simpleTemplateSpec", "v1.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/TemplateSpecVersionsList.json
func ExampleTemplateSpecVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTemplateSpecVersionsClient().NewListPager("templateSpecRG", "simpleTemplateSpec", nil)
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
		// page.TemplateSpecVersionsListResult = armtemplatespecs.TemplateSpecVersionsListResult{
		// 	Value: []*armtemplatespecs.TemplateSpecVersion{
		// 		{
		// 			Name: to.Ptr("v1.0"),
		// 			Type: to.Ptr("Microsoft.Resources/templateSpecs/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/templateSpecRG/providers/Microsoft.Resources/templateSpecs/simpleTemplateSpec/versions/v1.0"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.1974346Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecVersionProperties{
		// 				Description: to.Ptr("This is version v1.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("v2.0"),
		// 			Type: to.Ptr("Microsoft.Resources/templateSpecs/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/templateSpecRG/providers/Microsoft.Resources/templateSpecs/simpleTemplateSpec/versions/v2.0"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.1974346Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armtemplatespecs.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecVersionProperties{
		// 				Description: to.Ptr("This is another version (v2.0)"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/BuiltInTemplateSpecVersionsList.json
func ExampleTemplateSpecVersionsClient_NewListBuiltInsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTemplateSpecVersionsClient().NewListBuiltInsPager("nameOfTheBuiltIn", nil)
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
		// page.TemplateSpecVersionsListResult = armtemplatespecs.TemplateSpecVersionsListResult{
		// 	Value: []*armtemplatespecs.TemplateSpecVersion{
		// 		{
		// 			Name: to.Ptr("v1.0"),
		// 			Type: to.Ptr("Microsoft.Resources/builtInTemplateSpecs/versions"),
		// 			ID: to.Ptr("/providers/Microsoft.Resources/builtInTemplateSpecs/nameOfTheBuiltIn/versions/v1.0"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-01T01:01:01.1075056Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-01T01:01:01.1075056Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecVersionProperties{
		// 				Description: to.Ptr("This is version v1.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("v2.0"),
		// 			Type: to.Ptr("Microsoft.Resources/builtInTemplateSpecs/versions"),
		// 			ID: to.Ptr("/providers/Microsoft.Resources/builtInTemplateSpecs/nameOfTheBuiltIn/versions/v2.0"),
		// 			SystemData: &armtemplatespecs.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-02T02:03:01.1974346Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-02T02:03:01.1974346Z"); return t}()),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armtemplatespecs.TemplateSpecVersionProperties{
		// 				Description: to.Ptr("This is another version (v2.0)"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Resources/stable/2022-02-01/examples/BuiltInTemplateSpecVersionsGet.json
func ExampleTemplateSpecVersionsClient_GetBuiltIn() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtemplatespecs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTemplateSpecVersionsClient().GetBuiltIn(ctx, "nameOfTheBuiltIn", "v1.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TemplateSpecVersion = armtemplatespecs.TemplateSpecVersion{
	// 	Name: to.Ptr("v1.0"),
	// 	Type: to.Ptr("Microsoft.Resources/builtInTemplateSpecs/versions"),
	// 	ID: to.Ptr("/providers/Microsoft.Resources/builtInTemplateSpecs/nameOfTheBuiltIn/versions/v1.0"),
	// 	SystemData: &armtemplatespecs.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-01T01:01:01.1075056Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-01T01:01:01.1075056Z"); return t}()),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armtemplatespecs.TemplateSpecVersionProperties{
	// 		Description: to.Ptr("This is version v1.0"),
	// 	},
	// }
}
