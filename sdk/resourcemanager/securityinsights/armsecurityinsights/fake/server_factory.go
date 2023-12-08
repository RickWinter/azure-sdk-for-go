//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armsecurityinsights.ClientFactory type.
type ServerFactory struct {
	ActionsServer                            ActionsServer
	AlertRuleTemplatesServer                 AlertRuleTemplatesServer
	AlertRulesServer                         AlertRulesServer
	AutomationRulesServer                    AutomationRulesServer
	BookmarksServer                          BookmarksServer
	DataConnectorsServer                     DataConnectorsServer
	IncidentCommentsServer                   IncidentCommentsServer
	IncidentRelationsServer                  IncidentRelationsServer
	IncidentsServer                          IncidentsServer
	OperationsServer                         OperationsServer
	SentinelOnboardingStatesServer           SentinelOnboardingStatesServer
	ThreatIntelligenceIndicatorServer        ThreatIntelligenceIndicatorServer
	ThreatIntelligenceIndicatorMetricsServer ThreatIntelligenceIndicatorMetricsServer
	ThreatIntelligenceIndicatorsServer       ThreatIntelligenceIndicatorsServer
	WatchlistItemsServer                     WatchlistItemsServer
	WatchlistsServer                         WatchlistsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armsecurityinsights.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armsecurityinsights.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                        *ServerFactory
	trMu                                       sync.Mutex
	trActionsServer                            *ActionsServerTransport
	trAlertRuleTemplatesServer                 *AlertRuleTemplatesServerTransport
	trAlertRulesServer                         *AlertRulesServerTransport
	trAutomationRulesServer                    *AutomationRulesServerTransport
	trBookmarksServer                          *BookmarksServerTransport
	trDataConnectorsServer                     *DataConnectorsServerTransport
	trIncidentCommentsServer                   *IncidentCommentsServerTransport
	trIncidentRelationsServer                  *IncidentRelationsServerTransport
	trIncidentsServer                          *IncidentsServerTransport
	trOperationsServer                         *OperationsServerTransport
	trSentinelOnboardingStatesServer           *SentinelOnboardingStatesServerTransport
	trThreatIntelligenceIndicatorServer        *ThreatIntelligenceIndicatorServerTransport
	trThreatIntelligenceIndicatorMetricsServer *ThreatIntelligenceIndicatorMetricsServerTransport
	trThreatIntelligenceIndicatorsServer       *ThreatIntelligenceIndicatorsServerTransport
	trWatchlistItemsServer                     *WatchlistItemsServerTransport
	trWatchlistsServer                         *WatchlistsServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "ActionsClient":
		initServer(s, &s.trActionsServer, func() *ActionsServerTransport { return NewActionsServerTransport(&s.srv.ActionsServer) })
		resp, err = s.trActionsServer.Do(req)
	case "AlertRuleTemplatesClient":
		initServer(s, &s.trAlertRuleTemplatesServer, func() *AlertRuleTemplatesServerTransport {
			return NewAlertRuleTemplatesServerTransport(&s.srv.AlertRuleTemplatesServer)
		})
		resp, err = s.trAlertRuleTemplatesServer.Do(req)
	case "AlertRulesClient":
		initServer(s, &s.trAlertRulesServer, func() *AlertRulesServerTransport { return NewAlertRulesServerTransport(&s.srv.AlertRulesServer) })
		resp, err = s.trAlertRulesServer.Do(req)
	case "AutomationRulesClient":
		initServer(s, &s.trAutomationRulesServer, func() *AutomationRulesServerTransport {
			return NewAutomationRulesServerTransport(&s.srv.AutomationRulesServer)
		})
		resp, err = s.trAutomationRulesServer.Do(req)
	case "BookmarksClient":
		initServer(s, &s.trBookmarksServer, func() *BookmarksServerTransport { return NewBookmarksServerTransport(&s.srv.BookmarksServer) })
		resp, err = s.trBookmarksServer.Do(req)
	case "DataConnectorsClient":
		initServer(s, &s.trDataConnectorsServer, func() *DataConnectorsServerTransport {
			return NewDataConnectorsServerTransport(&s.srv.DataConnectorsServer)
		})
		resp, err = s.trDataConnectorsServer.Do(req)
	case "IncidentCommentsClient":
		initServer(s, &s.trIncidentCommentsServer, func() *IncidentCommentsServerTransport {
			return NewIncidentCommentsServerTransport(&s.srv.IncidentCommentsServer)
		})
		resp, err = s.trIncidentCommentsServer.Do(req)
	case "IncidentRelationsClient":
		initServer(s, &s.trIncidentRelationsServer, func() *IncidentRelationsServerTransport {
			return NewIncidentRelationsServerTransport(&s.srv.IncidentRelationsServer)
		})
		resp, err = s.trIncidentRelationsServer.Do(req)
	case "IncidentsClient":
		initServer(s, &s.trIncidentsServer, func() *IncidentsServerTransport { return NewIncidentsServerTransport(&s.srv.IncidentsServer) })
		resp, err = s.trIncidentsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SentinelOnboardingStatesClient":
		initServer(s, &s.trSentinelOnboardingStatesServer, func() *SentinelOnboardingStatesServerTransport {
			return NewSentinelOnboardingStatesServerTransport(&s.srv.SentinelOnboardingStatesServer)
		})
		resp, err = s.trSentinelOnboardingStatesServer.Do(req)
	case "ThreatIntelligenceIndicatorClient":
		initServer(s, &s.trThreatIntelligenceIndicatorServer, func() *ThreatIntelligenceIndicatorServerTransport {
			return NewThreatIntelligenceIndicatorServerTransport(&s.srv.ThreatIntelligenceIndicatorServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorServer.Do(req)
	case "ThreatIntelligenceIndicatorMetricsClient":
		initServer(s, &s.trThreatIntelligenceIndicatorMetricsServer, func() *ThreatIntelligenceIndicatorMetricsServerTransport {
			return NewThreatIntelligenceIndicatorMetricsServerTransport(&s.srv.ThreatIntelligenceIndicatorMetricsServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorMetricsServer.Do(req)
	case "ThreatIntelligenceIndicatorsClient":
		initServer(s, &s.trThreatIntelligenceIndicatorsServer, func() *ThreatIntelligenceIndicatorsServerTransport {
			return NewThreatIntelligenceIndicatorsServerTransport(&s.srv.ThreatIntelligenceIndicatorsServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorsServer.Do(req)
	case "WatchlistItemsClient":
		initServer(s, &s.trWatchlistItemsServer, func() *WatchlistItemsServerTransport {
			return NewWatchlistItemsServerTransport(&s.srv.WatchlistItemsServer)
		})
		resp, err = s.trWatchlistItemsServer.Do(req)
	case "WatchlistsClient":
		initServer(s, &s.trWatchlistsServer, func() *WatchlistsServerTransport { return NewWatchlistsServerTransport(&s.srv.WatchlistsServer) })
		resp, err = s.trWatchlistsServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}