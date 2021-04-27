package servers

import (
	"context"
	"crypto/rand"
	"crypto/rsa"

	"github.com/google/uuid"

	pb "github.com/phaesoo/keybox/gen/go/proto" // Update
)

type AdminServer struct {
	pb.UnimplementedAdminServer
}

func (s *AdminServer) RegisterKey(context.Context, *pb.RegisterRequest) (*pb.RegisterReply, error) {
	keyID := uuid.NewString()
	_, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		KeyId: keyID,
	}, nil
}
func (s *AdminServer) DeregisterKey(context.Context, *pb.DecryptionRequest) (*pb.DeregisterReply, error) {
	return nil, nil
}
