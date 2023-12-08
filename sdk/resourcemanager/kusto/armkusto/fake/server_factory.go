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

// ServerFactory is a fake server for instances of the armkusto.ClientFactory type.
type ServerFactory struct {
	AttachedDatabaseConfigurationsServer AttachedDatabaseConfigurationsServer
	ClusterPrincipalAssignmentsServer    ClusterPrincipalAssignmentsServer
	ClustersServer                       ClustersServer
	DataConnectionsServer                DataConnectionsServer
	DatabaseServer                       DatabaseServer
	DatabasePrincipalAssignmentsServer   DatabasePrincipalAssignmentsServer
	DatabasesServer                      DatabasesServer
	ManagedPrivateEndpointsServer        ManagedPrivateEndpointsServer
	OperationsServer                     OperationsServer
	OperationsResultsServer              OperationsResultsServer
	OperationsResultsLocationServer      OperationsResultsLocationServer
	PrivateEndpointConnectionsServer     PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer           PrivateLinkResourcesServer
	SKUsServer                           SKUsServer
	SandboxCustomImagesServer            SandboxCustomImagesServer
	ScriptsServer                        ScriptsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armkusto.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armkusto.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                    *ServerFactory
	trMu                                   sync.Mutex
	trAttachedDatabaseConfigurationsServer *AttachedDatabaseConfigurationsServerTransport
	trClusterPrincipalAssignmentsServer    *ClusterPrincipalAssignmentsServerTransport
	trClustersServer                       *ClustersServerTransport
	trDataConnectionsServer                *DataConnectionsServerTransport
	trDatabaseServer                       *DatabaseServerTransport
	trDatabasePrincipalAssignmentsServer   *DatabasePrincipalAssignmentsServerTransport
	trDatabasesServer                      *DatabasesServerTransport
	trManagedPrivateEndpointsServer        *ManagedPrivateEndpointsServerTransport
	trOperationsServer                     *OperationsServerTransport
	trOperationsResultsServer              *OperationsResultsServerTransport
	trOperationsResultsLocationServer      *OperationsResultsLocationServerTransport
	trPrivateEndpointConnectionsServer     *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer           *PrivateLinkResourcesServerTransport
	trSKUsServer                           *SKUsServerTransport
	trSandboxCustomImagesServer            *SandboxCustomImagesServerTransport
	trScriptsServer                        *ScriptsServerTransport
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
	case "AttachedDatabaseConfigurationsClient":
		initServer(s, &s.trAttachedDatabaseConfigurationsServer, func() *AttachedDatabaseConfigurationsServerTransport {
			return NewAttachedDatabaseConfigurationsServerTransport(&s.srv.AttachedDatabaseConfigurationsServer)
		})
		resp, err = s.trAttachedDatabaseConfigurationsServer.Do(req)
	case "ClusterPrincipalAssignmentsClient":
		initServer(s, &s.trClusterPrincipalAssignmentsServer, func() *ClusterPrincipalAssignmentsServerTransport {
			return NewClusterPrincipalAssignmentsServerTransport(&s.srv.ClusterPrincipalAssignmentsServer)
		})
		resp, err = s.trClusterPrincipalAssignmentsServer.Do(req)
	case "ClustersClient":
		initServer(s, &s.trClustersServer, func() *ClustersServerTransport { return NewClustersServerTransport(&s.srv.ClustersServer) })
		resp, err = s.trClustersServer.Do(req)
	case "DataConnectionsClient":
		initServer(s, &s.trDataConnectionsServer, func() *DataConnectionsServerTransport {
			return NewDataConnectionsServerTransport(&s.srv.DataConnectionsServer)
		})
		resp, err = s.trDataConnectionsServer.Do(req)
	case "DatabaseClient":
		initServer(s, &s.trDatabaseServer, func() *DatabaseServerTransport { return NewDatabaseServerTransport(&s.srv.DatabaseServer) })
		resp, err = s.trDatabaseServer.Do(req)
	case "DatabasePrincipalAssignmentsClient":
		initServer(s, &s.trDatabasePrincipalAssignmentsServer, func() *DatabasePrincipalAssignmentsServerTransport {
			return NewDatabasePrincipalAssignmentsServerTransport(&s.srv.DatabasePrincipalAssignmentsServer)
		})
		resp, err = s.trDatabasePrincipalAssignmentsServer.Do(req)
	case "DatabasesClient":
		initServer(s, &s.trDatabasesServer, func() *DatabasesServerTransport { return NewDatabasesServerTransport(&s.srv.DatabasesServer) })
		resp, err = s.trDatabasesServer.Do(req)
	case "ManagedPrivateEndpointsClient":
		initServer(s, &s.trManagedPrivateEndpointsServer, func() *ManagedPrivateEndpointsServerTransport {
			return NewManagedPrivateEndpointsServerTransport(&s.srv.ManagedPrivateEndpointsServer)
		})
		resp, err = s.trManagedPrivateEndpointsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "OperationsResultsClient":
		initServer(s, &s.trOperationsResultsServer, func() *OperationsResultsServerTransport {
			return NewOperationsResultsServerTransport(&s.srv.OperationsResultsServer)
		})
		resp, err = s.trOperationsResultsServer.Do(req)
	case "OperationsResultsLocationClient":
		initServer(s, &s.trOperationsResultsLocationServer, func() *OperationsResultsLocationServerTransport {
			return NewOperationsResultsLocationServerTransport(&s.srv.OperationsResultsLocationServer)
		})
		resp, err = s.trOperationsResultsLocationServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "SKUsClient":
		initServer(s, &s.trSKUsServer, func() *SKUsServerTransport { return NewSKUsServerTransport(&s.srv.SKUsServer) })
		resp, err = s.trSKUsServer.Do(req)
	case "SandboxCustomImagesClient":
		initServer(s, &s.trSandboxCustomImagesServer, func() *SandboxCustomImagesServerTransport {
			return NewSandboxCustomImagesServerTransport(&s.srv.SandboxCustomImagesServer)
		})
		resp, err = s.trSandboxCustomImagesServer.Do(req)
	case "ScriptsClient":
		initServer(s, &s.trScriptsServer, func() *ScriptsServerTransport { return NewScriptsServerTransport(&s.srv.ScriptsServer) })
		resp, err = s.trScriptsServer.Do(req)
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