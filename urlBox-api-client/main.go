package main

import (
	"context"
	"log"
	"net/http"

	pb "ekstrah.com/go-protoBox-grpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Send_GenMessage(client pb.GenURLManagementClient, msg *pb.ExURLReq) string {
	resp, err := client.GenNewURL(context.Background(), msg)
	if err != nil {
		log.Fatal("Error Getting a response on GenURL %v", err)
	}
	return resp.GetNewURL()
}

func Get_ReDirURL(client pb.GenURLManagementClient, msg *pb.ReDirReq) string {
	resp, err := client.ReDirURL(context.Background(), msg)
	if err != nil {
		log.Fatal("Error getting a response on ReDirURL %v", err)
	}
	return resp.GetResURL()
}

func main() {
	conn, err := grpc.Dial("172.17.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to 9000 with %v", err)
	}
	client := pb.NewGenURLManagementClient(conn)

	// Gin Framework Initialization
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	//Gin Router Setup
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "Main web",
		})
	})
	r.POST("/", func(c *gin.Context) {
		uurl := c.PostForm("url")
		msg := &pb.ExURLReq{OriURL: uurl, UserID: "free"}
		NUrl := Send_GenMessage(client, msg)
		NUrl = "http://localhost:8080/" + NUrl
		c.HTML(http.StatusOK, "simple.tmpl.html", gin.H{
			"NURL": NUrl,
		})
	})
	r.GET("/ping", func(c *gin.Context) {

		ReURL := "WtNGsTC"
		msg := &pb.ReDirReq{ReqURL: ReURL}
		OURL := Get_ReDirURL(client, msg)
		c.Redirect(http.StatusMovedPermanently, OURL)
	})
	r.GET("/:NURL", func(c *gin.Context) {
		name := c.Param("NURL")
		ReURL := name
		msg := &pb.ReDirReq{ReqURL: ReURL}
		OURL := Get_ReDirURL(client, msg)
		c.Redirect(http.StatusMovedPermanently, OURL)
	})

	r.Run(":8080") //
}
