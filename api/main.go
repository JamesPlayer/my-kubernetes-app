package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/JamesPlayer/my-kubernetes-app/microservice/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "my-k8s-app-microservice.default.svc.cluster.local:50051", "address of microservice")

type Response struct {
	Msg    string            `json:"msg"`
	Env    map[string]string `json:"env"`
	Config map[string]string `json:"config"`
}

func getFileContents(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	microserviceClient := pb.NewPingPongServiceClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get config from volume
		color, _ := getFileContents("/etc/app-config/color")
		if err != nil {
			fmt.Fprintf(w, "Error getting color config")
			return
		}

		logoUrl, _ := getFileContents("/etc/app-config/logo_url")
		if err != nil {
			fmt.Fprintf(w, "Error getting logo_url config")
			return
		}

		response := []Response{
			{
				Msg: "Hit API server",
				Env: map[string]string{
					"MY_NODE_NAME": os.Getenv("MY_NODE_NAME"),
					"MY_POD_NAME":  os.Getenv("MY_POD_NAME"),
					"MY_POD_IP":    os.Getenv("MY_POD_IP"),
					"MY_SECRET":    os.Getenv("MY_SECRET"),
				},
				Config: map[string]string{
					"color":   color,
					"logoUrl": logoUrl,
				},
			},
		}

		// Get info from microservice
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		microserviceReply, err := microserviceClient.Ping(ctx, &pb.PingPongRequest{Msg: "Ping"})
		if err != nil {
			log.Fatalf("could not ping: %v", err)
		}

		response = append(response, Response{
			Msg: microserviceReply.GetMsg(),
			Env: microserviceReply.GetEnv(),
		})

		json, err := json.Marshal(response)

		if err != nil {
			fmt.Fprintf(w, "Error marshalling response into json")
			return
		}

		fmt.Fprintf(w, "%s", string(json))
	})

	http.ListenAndServe(":8080", nil)
}
