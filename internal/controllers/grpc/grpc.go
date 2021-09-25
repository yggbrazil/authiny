package grpc

import (
	"context"
	"log"
	"os"

	"authiny/internal/storages"
	"authiny/internal/usecases"
	"authiny/pkg/authiny_grpc"

	yg_context "github.com/yggbrazil/go-toolbox/context"
	"github.com/yggbrazil/go-toolbox/grpc"
)

type GrpcServer interface {
	Run()

	Login(ctx context.Context, req *authiny_grpc.LoginRequest) (*authiny_grpc.LoginResponse, error)

	CreateApplication(ctx context.Context, req *authiny_grpc.CreateApplicationRequest) (*authiny_grpc.CreateApplicationResponse, error)
}

type grpcServer struct {
	grpcServer *grpc.Server
	authiny_grpc.UnimplementedAuthinyServer
	applicationService usecases.ApplicationService
}

func NewGrpcServer(storage storages.Storage) (GrpcServer, error) {
	grpc_port := os.Getenv("AUTHINY_GRPC_PORT")

	s, err := grpc.NewServer(grpc_port)
	if err != nil {
		return &grpcServer{}, err
	}

	applicationService, err := usecases.NewApplicationService(storage)
	if err != nil {
		return &grpcServer{}, err
	}

	grpcServer := &grpcServer{
		grpcServer:         s,
		applicationService: applicationService,
	}

	authiny_grpc.RegisterAuthinyServer(s.GRPCServer, grpcServer)

	return grpcServer, nil
}

func (s *grpcServer) Run() {
	if err := s.grpcServer.Run(); err != nil {
		log.Fatalf("failed on GRPC server: %v", err)
	}
}

func (s *grpcServer) Login(ctx context.Context, req *authiny_grpc.LoginRequest) (*authiny_grpc.LoginResponse, error) {
	return &authiny_grpc.LoginResponse{
		Token: "token",
	}, nil
}

func (s *grpcServer) CreateApplication(ctx context.Context, req *authiny_grpc.CreateApplicationRequest) (*authiny_grpc.CreateApplicationResponse, error) {
	ctx = yg_context.AddTrace(ctx)

	applicationID, err := s.applicationService.Create(ctx, req.ApplicationName)
	if err != nil {
		return &authiny_grpc.CreateApplicationResponse{}, err
	}

	log.Println(applicationID)

	return &authiny_grpc.CreateApplicationResponse{
		ApplicationId: applicationID,
	}, nil
}
