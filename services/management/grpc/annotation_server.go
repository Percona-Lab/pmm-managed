// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package grpc

import (
	"context"
	"fmt"

	"github.com/percona/pmm/api/managementpb"

	"github.com/percona/pmm-managed/services/grafana"

	"google.golang.org/grpc/metadata"
)

// AnnotationServer is a server for making annotations in Grafana.
type AnnotationServer struct {
	grafanaClient *grafana.Client
}

// NewAnnotationServer creates Annotation Server.
func NewAnnotationServer(grafanaClient *grafana.Client) *AnnotationServer {
	return &AnnotationServer{
		grafanaClient: grafanaClient,
	}
}

// AddAnnotation adds annotation to Grafana.
func (as *AnnotationServer) AddAnnotation(ctx context.Context, req *managementpb.AddAnnotationRequest) (*managementpb.AddAnnotationResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("cannot get headers from metadata %v", headers)
	}
	// get authorization from headers.
	authorization := headers["authorization"][0]
	message, err := as.grafanaClient.CreateAnnotation(ctx, req.Tags, req.Text, authorization)
	if err != nil {
		return nil, err
	}
	return &managementpb.AddAnnotationResponse{Message: message}, nil
}
