package server

import (
	"context"
	"exc8/pb"
	// "fmt"
	// "log/slog"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	// wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer

	drinks []*pb.Drink
	orders []*pb.Order
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{
		drinks: []*pb.Drink{
            {Id: 1, Name: "Spritzer", Desc: "Wine with soda", Price: 2},
            {Id: 2, Name: "Beer", Desc: "Hagenberger Gold", Price: 3},
            {Id: 3, Name: "Coffee", Desc: "Mifare isn't that secure", Price: 0},
        },
        orders: []*pb.Order{},
    }
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	return srv.Serve(lis)
}

// todo implement functions

func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.GetDrinksResponse, error) {
	return &pb.GetDrinksResponse{Drinks: s.drinks}, nil
}

func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.GetOrdersResponse, error) {
	return &pb.GetOrdersResponse{Orders: s.orders}, nil
}

func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	order := &pb.Order{
		DrinkId: req.DrinkId,
		Amount:  req.Amount,
	}
	s.orders = append(s.orders, order)

	return &pb.OrderResponse{Order: order}, nil
}