package servers

import (
	"context"
	"crypto/rand"
	"crypto/rsa"

	pb "github.com/phaesoo/keybox/gen/go/proto" // Update
)

type AdminServer struct {
	pb.UnimplementedAdminServer
}

func (s *AdminServer) RegisterKey(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (s *AdminServer) DeregisterKey(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}
