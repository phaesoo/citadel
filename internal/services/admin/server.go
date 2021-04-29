package admin

import (
	"context"

	pb "github.com/phaesoo/keybox/gen/go/proto" // Update
)

type service interface {
	RegisterKey(ctx context.Context, userID string) (string, error)
}

type Server struct {
	pb.UnimplementedAdminServer
	service service
}

func NewServer(service service) *Server {
	return &Server{service: service}
}

func (s *Server) RegisterKey(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	keyID, err := s.service.RegisterKey(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		KeyId: keyID,
	}, nil
}

func (s *Server) DeregisterKey(context.Context, *pb.DeregisterRequest) (*pb.DeregisterReply, error) {
	return nil, nil
}
