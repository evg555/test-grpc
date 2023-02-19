build:
	protoc -I api/proto --go_out=./pkg/api --go-grpc_out=./pkg/api --go-grpc_opt=paths=source_relative api/proto/user.proto
	go build -o ./bin/server ./cmd/server/main.go
	go build -o ./bin/client ./cmd/client/main.go
run server:
	./bin/server

