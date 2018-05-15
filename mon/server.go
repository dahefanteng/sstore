package mon

import (
	"google.golang.org/grpc"
)

const port = ":2379"

func StartServer() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	RegisterMonServer(s, &MonServer{})
	s.Serve(ln)
}

type MonServer struct{}

func (m *MonServer) SetMap(ctx context.Context, req *SetMapRequest) (*SetMapResponse, error) {
	res := &SetMapResponse{Rtcode: 200}
	return res, nil
}
