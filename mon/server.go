package mon

import (
	"google.golang.org/grpc"
	"net"
	"context"
)

const port = ":2379"

func StartServer() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	RegisterMonServer(s, &monServer{})
	s.Serve(ln)
}

type monServer struct{
	engine string
}

func (m *monServer) SetMap(ctx context.Context, req *SetMapRequest) (*SetMapResponse, error) {
	res := &SetMapResponse{Rtcode: 200}
	return res, nil
}
