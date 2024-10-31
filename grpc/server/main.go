// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"server/api/apiconnect"

	pb "server/api" // Change to your actual import path for generated protobuf code

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type MainService struct {
	apiconnect.UnimplementedMainHandler
}

func (s *MainService) ListAccount(
	ctx context.Context,
	req *connect.Request[pb.ListAccountReq],
) (*connect.Response[pb.ListAccountResp], error) {
	fmt.Println("good")
	// Here you would typically fetch accounts from a database or other source
	accounts := []*pb.Account{
		{UserName: "Jimmy1", DisplayName: "Jimmy1", State: 1},
		{UserName: "Jimmy2", DisplayName: "Jimmy2", State: 1},
		// Add more accounts as needed
	}
	return connect.NewResponse(
		&pb.ListAccountResp{List: accounts, Total: string(len(accounts))},
	), nil
}

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	mux.Handle(apiconnect.NewMainHandler(&MainService{}))
	err := http.ListenAndServe(
		"localhost:50051",
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}
