package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
	"unsafe"

	"net"

	pb "ekstrah.com/go-protoBox-grpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type GenURLManagementServer struct {
	pb.UnimplementedGenURLManagementServer
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func genNewURL() string {
	return RandStringBytesMaskImprSrcUnsafe(7)
}

func saveURL(oriURL string, newURL string, userID string) int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to mongodb")
		return -1
	}
	collection := client.Database("users").Collection("free")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"oriURL", oriURL}, {"newURL", newURL}, {"userID", userID}})
	id := res.InsertedID
	fmt.Println("%v", id)
	return 0
}

// GenNewURL(context.Context, *ExURLReq) (*ExURLRes, error)
func (s *GenURLManagementServer) GenNewURL(ctx context.Context, req *pb.ExURLReq) (*pb.ExURLRes, error) {
	ReqNewURL := genNewURL()
	ReqOriURL := req.GetOriURL()
	ReqUserID := req.GetUserID()
	val := saveURL(ReqOriURL, ReqNewURL, ReqUserID)
	if val != 0 {
		return &pb.ExURLRes{OriURL: "error", NewURL: "error", UserID: "error", Count: 0}, nil
	}
	return &pb.ExURLRes{OriURL: ReqOriURL, NewURL: ReqNewURL, UserID: ReqUserID, Count: 0}, nil
}

func findURL(tURL string) string {
	filter := bson.D{{"newURL", tURL}}
	var result bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to mongodb")
		return "failed"
	}
	collection := client.Database("users").Collection("free")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return "failed"
		}
		panic(err)
	}
	retURL := result["oriURL"].(string)
	return retURL
}
func (s *GenURLManagementServer) ReDirURL(ctx context.Context, req *pb.ReDirReq) (*pb.ReDirRes, error) {
	testResult := findURL(req.GetReqURL())
	return &pb.ReDirRes{ResURL: testResult}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Server Failed to listen to port %v", port)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGenURLManagementServer(grpcServer, &GenURLManagementServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Server failed to start due to %v", err)
	}
}
