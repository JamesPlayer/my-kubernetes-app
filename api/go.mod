module github.com/JamesPlayer/my-kubernetes-app/api

go 1.23.4

// replace github.com/JamesPlayer/my-kubernetes-app/microservice => ../microservice

require (
	github.com/JamesPlayer/my-kubernetes-app/microservice v0.0.0-20250116090509-5b7066d32ff6
	google.golang.org/grpc v1.69.4
)

require (
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)
