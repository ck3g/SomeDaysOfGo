package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ck3g/SomeDaysOfGo/gRPC-Master-Class/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Running Blog RPC client...")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := blogpb.NewBlogServiceClient(conn)

	blogID := createBlog(c)
	readBlog(c, "1dkasdkh")
	readBlog(c, blogID)
}

func createBlog(c blogpb.BlogServiceClient) string {
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "Vitali",
			Title:    "My First Blog",
			Content:  "Content of the first blog",
		},
	}
	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CreateBlog RPC: %v", err)
	}

	log.Printf("Blog has been created: %v\n", res)

	blogID := res.GetBlog().GetId()
	return blogID
}

func readBlog(c blogpb.BlogServiceClient, blogID string) {
	fmt.Println("Reading the blog...")

	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blogID})
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
		return
	}

	fmt.Printf("Blog was read: %v\n", res)
}
