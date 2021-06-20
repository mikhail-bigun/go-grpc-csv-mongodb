package fetcher

import (
	"context"

	"google.golang.org/grpc"
)

type GRPCServer struct {
}

func (s *GRPCServer) Fetch(req *FetchRequest, resp_sender Fetcher_FetchServer) error {

	return nil
}
