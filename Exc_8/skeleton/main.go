package main

import (
    "exc8/client"
    "exc8/server"
	"time"
)

func main() {
	go func() {
		println("Starting gRPC server on :4000...")
		if err := server.StartGrpcServer(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)
	// todo start client
	c, err := client.NewGrpcClient()
	if err != nil {
		panic(err)
	}
	if err := c.Run(); err != nil {
		panic(err)
	}
	println("Orders complete!")
}
