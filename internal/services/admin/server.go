package admin

import (
	"context"
	"crypto/rand"
	"crypto/rsa"

	"github.com/google/uuid"

	pb "github.com/phaesoo/keybox/gen/go/proto" // Update
)

type service interface {
}

type Server struct {
	pb.UnimplementedAdminServer
	s service
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) RegisterKey(context.Context, *pb.RegisterRequest) (*pb.RegisterReply, error) {
	keyID := uuid.NewString()
	_, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		KeyId: keyID,
	}, nil
}

func (s *Server) DeregisterKey(context.Context, *pb.DecryptionRequest) (*pb.DeregisterReply, error) {
	return nil, nil
}
