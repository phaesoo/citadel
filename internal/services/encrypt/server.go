package encrypt

import (
	"context"

	pb "github.com/phaesoo/keybox/gen/go/proto" // Update
)

type service interface {
	Encrypt(ctx context.Context, userID, keyID string, plaintexts []string) ([]string, error)
}

type Server struct {
	pb.UnimplementedAdminServer
	service service
}

func NewServer(service service) *Server {
	return &Server{service: service}
}

func (s *Server) Encrypt(ctx context.Context, req *pb.EncryptRequest) (*pb.EncryptReply, error) {
	ciphertexts, err := s.service.Encrypt(ctx, req.UserId, req.KeyId, req.Plaintexts)
	if err != nil {
		return nil, err
	}
	return &pb.EncryptReply{
		Ciphertexts: ciphertexts,
	}, nil
}
