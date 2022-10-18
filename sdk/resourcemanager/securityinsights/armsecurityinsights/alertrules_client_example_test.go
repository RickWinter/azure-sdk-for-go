//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/GetAllAlertRules.json
func ExampleAlertRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myRg", "myWorkspace", nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/GetFusionAlertRule.json
func ExampleAlertRulesClient_Get_getAFusionAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myRg", "myWorkspace", "myFirstFusionRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/GetMicrosoftSecurityIncidentCreationAlertRule.json
func ExampleAlertRulesClient_Get_getAMicrosoftSecurityIncidentCreationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myRg", "myWorkspace", "microsoftSecurityIncidentCreationRuleExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/GetScheduledAlertRule.json
func ExampleAlertRulesClient_Get_getAScheduledAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/GetNrtAlertRule.json
func ExampleAlertRulesClient_Get_getAnNrtAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
func ExampleAlertRulesClient_CreateOrUpdate_createsOrUpdatesAFusionAlertRuleWithScenarioExclusionPattern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "myFirstFusionRule", &armsecurityinsights.FusionAlertRule{
		Etag: to.Ptr("3d00c3ca-0000-0100-0000-5d42d5010000"),
		Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
		Properties: &armsecurityinsights.FusionAlertRuleProperties{
			AlertRuleTemplateName: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
			Enabled:               to.Ptr(true),
			SourceSettings: []*armsecurityinsights.FusionSourceSettings{
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Anomalies"),
				},
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Alert providers"),
					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Active Directory Identity Protection"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Defender"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Defender for IoT"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft 365 Defender"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Cloud App Security"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Endpoint"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Identity"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Office 365"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Sentinel scheduled analytics rules"),
						}},
				},
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Raw logs from other sources"),
					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
						{
							Enabled:           to.Ptr(true),
							SeverityFilters:   &armsecurityinsights.FusionSubTypeSeverityFilter{},
							SourceSubTypeName: to.Ptr("Palo Alto Networks"),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateFusionAlertRule.json
func ExampleAlertRulesClient_CreateOrUpdate_createsOrUpdatesAFusionAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "myFirstFusionRule", &armsecurityinsights.FusionAlertRule{
		Etag: to.Ptr("3d00c3ca-0000-0100-0000-5d42d5010000"),
		Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
		Properties: &armsecurityinsights.FusionAlertRuleProperties{
			AlertRuleTemplateName: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
			Enabled:               to.Ptr(true),
			SourceSettings: []*armsecurityinsights.FusionSourceSettings{
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Anomalies"),
				},
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Alert providers"),
					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Active Directory Identity Protection"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Defender"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Defender for IoT"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft 365 Defender"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Cloud App Security"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Endpoint"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Identity"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Microsoft Defender for Office 365"),
						},
						{
							Enabled: to.Ptr(true),
							SeverityFilters: &armsecurityinsights.FusionSubTypeSeverityFilter{
								Filters: []*armsecurityinsights.FusionSubTypeSeverityFiltersItem{
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
									},
									{
										Enabled:  to.Ptr(true),
										Severity: to.Ptr(armsecurityinsights.AlertSeverityInformational),
									}},
							},
							SourceSubTypeName: to.Ptr("Azure Sentinel scheduled analytics rules"),
						}},
				},
				{
					Enabled:    to.Ptr(true),
					SourceName: to.Ptr("Raw logs from other sources"),
					SourceSubTypes: []*armsecurityinsights.FusionSourceSubTypeSetting{
						{
							Enabled:           to.Ptr(true),
							SeverityFilters:   &armsecurityinsights.FusionSubTypeSeverityFilter{},
							SourceSubTypeName: to.Ptr("Palo Alto Networks"),
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateMicrosoftSecurityIncidentCreationAlertRule.json
func ExampleAlertRulesClient_CreateOrUpdate_createsOrUpdatesAMicrosoftSecurityIncidentCreationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "microsoftSecurityIncidentCreationRuleExample", &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRule{
		Etag: to.Ptr("\"260097e0-0000-0d00-0000-5d6fa88f0000\""),
		Kind: to.Ptr(armsecurityinsights.AlertRuleKindMicrosoftSecurityIncidentCreation),
		Properties: &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRuleProperties{
			ProductFilter: to.Ptr(armsecurityinsights.MicrosoftSecurityProductNameMicrosoftCloudAppSecurity),
			DisplayName:   to.Ptr("testing displayname"),
			Enabled:       to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateNrtAlertRule.json
func ExampleAlertRulesClient_CreateOrUpdate_createsOrUpdatesANrtAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.NrtAlertRule{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Kind: to.Ptr(armsecurityinsights.AlertRuleKindNRT),
		Properties: &armsecurityinsights.NrtAlertRuleProperties{
			Description: to.Ptr(""),
			DisplayName: to.Ptr("Rule2"),
			Enabled:     to.Ptr(true),
			EventGroupingSettings: &armsecurityinsights.EventGroupingSettings{
				AggregationKind: to.Ptr(armsecurityinsights.EventGroupingAggregationKindAlertPerResult),
			},
			IncidentConfiguration: &armsecurityinsights.IncidentConfiguration{
				CreateIncident: to.Ptr(true),
				GroupingConfiguration: &armsecurityinsights.GroupingConfiguration{
					Enabled: to.Ptr(true),
					GroupByEntities: []*armsecurityinsights.EntityMappingType{
						to.Ptr(armsecurityinsights.EntityMappingTypeHost),
						to.Ptr(armsecurityinsights.EntityMappingTypeAccount)},
					LookbackDuration:     to.Ptr("PT5H"),
					MatchingMethod:       to.Ptr(armsecurityinsights.MatchingMethodSelected),
					ReopenClosedIncident: to.Ptr(false),
				},
			},
			Query:               to.Ptr("ProtectionStatus | extend HostCustomEntity = Computer | extend IPCustomEntity = ComputerIP_Hidden"),
			Severity:            to.Ptr(armsecurityinsights.AlertSeverityHigh),
			SuppressionDuration: to.Ptr("PT1H"),
			SuppressionEnabled:  to.Ptr(false),
			Tactics: []*armsecurityinsights.AttackTactic{
				to.Ptr(armsecurityinsights.AttackTacticPersistence),
				to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
			Techniques: []*string{
				to.Ptr("T1037"),
				to.Ptr("T1021")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateScheduledAlertRule.json
func ExampleAlertRulesClient_CreateOrUpdate_createsOrUpdatesAScheduledAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.ScheduledAlertRule{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Kind: to.Ptr(armsecurityinsights.AlertRuleKindScheduled),
		Properties: &armsecurityinsights.ScheduledAlertRuleProperties{
			AlertDetailsOverride: &armsecurityinsights.AlertDetailsOverride{
				AlertDescriptionFormat: to.Ptr("Suspicious activity was made by {{ComputerIP}}"),
				AlertDisplayNameFormat: to.Ptr("Alert from {{Computer}}"),
			},
			CustomDetails: map[string]*string{
				"OperatingSystemName": to.Ptr("OSName"),
				"OperatingSystemType": to.Ptr("OSType"),
			},
			EntityMappings: []*armsecurityinsights.EntityMapping{
				{
					EntityType: to.Ptr(armsecurityinsights.EntityMappingTypeHost),
					FieldMappings: []*armsecurityinsights.FieldMapping{
						{
							ColumnName: to.Ptr("Computer"),
							Identifier: to.Ptr("FullName"),
						}},
				},
				{
					EntityType: to.Ptr(armsecurityinsights.EntityMappingTypeIP),
					FieldMappings: []*armsecurityinsights.FieldMapping{
						{
							ColumnName: to.Ptr("ComputerIP"),
							Identifier: to.Ptr("Address"),
						}},
				}},
			EventGroupingSettings: &armsecurityinsights.EventGroupingSettings{
				AggregationKind: to.Ptr(armsecurityinsights.EventGroupingAggregationKindAlertPerResult),
			},
			Query:            to.Ptr("Heartbeat"),
			QueryFrequency:   to.Ptr("PT1H"),
			QueryPeriod:      to.Ptr("P2DT1H30M"),
			Severity:         to.Ptr(armsecurityinsights.AlertSeverityHigh),
			TriggerOperator:  to.Ptr(armsecurityinsights.TriggerOperatorGreaterThan),
			TriggerThreshold: to.Ptr[int32](0),
			Description:      to.Ptr("An example for a scheduled rule"),
			DisplayName:      to.Ptr("My scheduled rule"),
			Enabled:          to.Ptr(true),
			IncidentConfiguration: &armsecurityinsights.IncidentConfiguration{
				CreateIncident: to.Ptr(true),
				GroupingConfiguration: &armsecurityinsights.GroupingConfiguration{
					Enabled: to.Ptr(true),
					GroupByAlertDetails: []*armsecurityinsights.AlertDetail{
						to.Ptr(armsecurityinsights.AlertDetailDisplayName)},
					GroupByCustomDetails: []*string{
						to.Ptr("OperatingSystemType"),
						to.Ptr("OperatingSystemName")},
					GroupByEntities: []*armsecurityinsights.EntityMappingType{
						to.Ptr(armsecurityinsights.EntityMappingTypeHost)},
					LookbackDuration:     to.Ptr("PT5H"),
					MatchingMethod:       to.Ptr(armsecurityinsights.MatchingMethodSelected),
					ReopenClosedIncident: to.Ptr(false),
				},
			},
			SuppressionDuration: to.Ptr("PT1H"),
			SuppressionEnabled:  to.Ptr(false),
			Tactics: []*armsecurityinsights.AttackTactic{
				to.Ptr(armsecurityinsights.AttackTacticPersistence),
				to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
			Techniques: []*string{
				to.Ptr("T1037"),
				to.Ptr("T1021")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/DeleteAlertRule.json
func ExampleAlertRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewAlertRulesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}