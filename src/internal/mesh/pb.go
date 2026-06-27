package mesh

import (
	"context"
	"fmt"
	"log"
	pb "github.com/whotterre/odysseus/src/internal/grpc/proto"
)

type meshServer struct {
	pb.UnimplementedMeshTelemetryServer
}

func NewMeshServer() *meshServer {
	return &meshServer{}
}

func (s *meshServer) SyncTelemetry(ctx context.Context, req *pb.MeshReport) (*pb.TelemetryResponse, error) {
	log.Printf("Received Telemetry dump from Node: %s", req.ReportId)
	log.Printf("Total neighbors detected in cache: %d", len(req.GetContent()))

	return &pb.TelemetryResponse{
		Success: true,
		Message: fmt.Sprintf("Successfully processed %d peers from node %s", 9, req.ReportId),
	}, nil
}

func (s *meshServer) StreamNetworkUpdates(in *pb.NodeIdentity, stream pb.MeshTelemetry_StreamNetworkUpdatesServer) error {
	return nil
}
