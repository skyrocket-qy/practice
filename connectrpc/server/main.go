package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"server/api/apiconnect"

	pb "server/api" // Change to your actual import path for generated protobuf code

	"connectrpc.com/connect"
	"github.com/rs/cors"
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
	mux.Handle(apiconnect.NewMainHandler(&MainService{}))

	// Create a CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // Allow the frontend URL
			"http://127.0.0.1:3000", // Also allow 127.0.0.1 if needed
		},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"}, // Allow necessary methods
		AllowedHeaders:   []string{"*"},                      // Allow any headers
	})

	// Start the server with CORS support
	err := http.ListenAndServe(
		"localhost:50051",
		c.Handler(h2c.NewHandler(mux, &http2.Server{})),
	)
	log.Fatalf("listen failed: %v", err)
}
