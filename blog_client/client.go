package main

import (
	"context"
	"fmt"
	"log"

	"github.com/goodbsw/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Stephane",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	createRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createRes)
	blogID := createRes.GetBlog().GetId()

	// read blog
	fmt.Println("Reading the blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "slkdfjsdlk",
	})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: blogID,
	})
	if readBlogErr != nil {
		fmt.Printf("Error happend while reading: %v\n", readBlogErr)
	}
	fmt.Printf("Blog was read: %v", readBlogRes)
}
