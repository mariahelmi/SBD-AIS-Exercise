package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.Dial(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	ctx := context.Background()
	// todo
	// 1. List drinks
	fmt.Println("Requesting drinks 🍹🍺☕")
	drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	fmt.Println("Available drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d  name:%q  price:%d  description:%q\n",
			d.Id, d.Name, d.Price, d.Desc)
	}

	// 2. Order a few drinks
	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> Ordering: 2 x %s\n", d.Name)
		_, err = c.client.OrderDrink(ctx, &pb.OrderRequest{
			DrinkId: d.Id,
			Amount:  2,
		})
		if err != nil {
			return err
		}
	}
	// 3. Order more drinks
	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> Ordering: 6 x %s\n", d.Name)
		_, err = c.client.OrderDrink(ctx, &pb.OrderRequest{
			DrinkId: d.Id,
			Amount:  6,
		})
		if err != nil {
			return err
		}
	}
	fmt.Println("Getting the bill 💹💹💹")
	ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	// Aggregate totals
	totals := map[uint64]uint64{}
	for _, o := range ordersResp.Orders {
		totals[o.DrinkId] += o.Amount
	}

	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> Total: %d x %s\n", totals[d.Id], d.Name)
	}

	fmt.Println("Orders complete!")
	// print responses after each call
	return nil
}
