package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Msg string `json:"msg"`
	Env map[string]string `json:"env"`

	
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json, err := json.Marshal(Response{
			Msg: "Hit API server",
			Env: map[string]string{
				"MY_NODE_NAME": os.Getenv("MY_NODE_NAME"),
				"MY_POD_NAME": os.Getenv("MY_POD_NAME"),
				"MY_POD_IP": os.Getenv("MY_POD_IP"),
				"MY_SECRET": os.Getenv("MY_SECRET"),
			},
		})

		if (err != nil) {
			fmt.Fprintf(w, "Error marshalling response into json")
			return
		}

		fmt.Fprintf(w, "%s", string(json))
    })

    http.ListenAndServe(":8080", nil)
}
