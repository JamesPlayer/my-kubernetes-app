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
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		color, err := getFileContents("/etc/app-config/color")
		if err != nil {
			fmt.Fprintf(w, "Error getting color config")
			return
		}

		logoUrl, err := getFileContents("/etc/app-config/logo_url")
		if err != nil {
			fmt.Fprintf(w, "Error getting logo_url config")
			return
		}
		
		json, err := json.Marshal(Response{
			Msg: "Hit API server.",
			Env: map[string]string{
				"MY_NODE_NAME": os.Getenv("MY_NODE_NAME"),
				"MY_POD_NAME": os.Getenv("MY_POD_NAME"),
				"MY_POD_IP": os.Getenv("MY_POD_IP"),
				"MY_SECRET": os.Getenv("MY_SECRET"),
			},
			Config: map[string]string{
				"color": color,
				"logoUrl": logoUrl,
			},
		})

		if err != nil {
			fmt.Fprintf(w, "Error marshalling response into json")
			return
		}

		fmt.Fprintf(w, "%s", string(json))
    })

    http.ListenAndServe(":8080", nil)
}
