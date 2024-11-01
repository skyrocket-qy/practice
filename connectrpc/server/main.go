package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	pb "server/api" // Adjust to your actual protobuf path
	"server/api/apiconnect"
)

type MainService struct {
	apiconnect.UnimplementedMainHandler
}

func (s *MainService) ListAccount(
	ctx context.Context,
	req *connect.Request[pb.ListAccountReq],
) (*connect.Response[pb.ListAccountResp], error) {
	fmt.Println("good")
	accounts := []*pb.Account{
		{UserName: "Jimmy1", DisplayName: "Jimmy1", State: 1},
		{UserName: "Jimmy2", DisplayName: "Jimmy2", State: 1},
	}
	return connect.NewResponse(
		&pb.ListAccountResp{List: accounts, Total: fmt.Sprintf("%d", len(accounts))},
	), nil
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(apiconnect.NewMainHandler(&MainService{}))

	// Setup reflection with the specified services
	reflector := grpcreflect.NewStaticReflector(
		"proto.Main", // Replace with actual service names
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	// Start the server with CORS and reflection support
	err := http.ListenAndServe(
		"localhost:50051",
		c.Handler(h2c.NewHandler(mux, &http2.Server{})),
	)
	if err != nil {
		log.Fatalf("listen failed: %v", err)
	}
}
