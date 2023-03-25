package main

import (
	"context"
	"fmt"
	"github.com/EdsonGustavoTofolo/gRPC-sample-golang/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client := pb.NewCategoryServiceClient(conn)

	list, err := client.ListCategories(ctx, &pb.Blank{})

	if err != nil {
		panic(err)
	}

	for _, category := range list.Categories {
		fmt.Printf("Category: %v\n", category.Name)
	}
}
