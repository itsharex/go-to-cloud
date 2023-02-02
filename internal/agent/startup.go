package agent

import (
	"fmt"
	gotocloud "go-to-cloud/internal/agent/proto"
	"go-to-cloud/internal/agent/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"strings"
)

const HEALTHCHECK_SERVICE = "gotocloud.agent.v1.Health"

func Startup(port *string) error {
	if !strings.HasPrefix(*port, ":") {
		*port = ":" + *port
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(server.AccessTokenInterceptor))

	healthSrv := health.NewServer()
	healthSrv.SetServingStatus(HEALTHCHECK_SERVICE, healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthSrv)

	gotocloud.RegisterAgentServer(s, &server.Agent{})
	log.Printf("agent server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}
