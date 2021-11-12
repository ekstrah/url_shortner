package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "ekstrah.com/go-protoBox-grpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Send_GenMessage(client pb.GenURLManagementClient, msg *pb.ExURLReq) {
	resp, err := client.GenNewURL(context.Background(), msg)
	if err != nil {
		log.Fatal("Error Getting a response %v", err)
	}
	fmt.Println(resp)
}

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to 9000 with %v", err)
	}
	client := pb.NewGenURLManagementClient(conn)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		msg := &pb.ExURLReq{OriURL: "https://google.com", UserID: "ekstrah"}
		Send_GenMessage(client, msg)
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.Run(":8080") //
}
