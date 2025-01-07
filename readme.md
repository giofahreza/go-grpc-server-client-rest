brew install protobuf
protoc --version

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc-gen-go --version

protoc --proto_path=pb pb/*.proto --go_out=paths=source_relative:pb --go-grpc_out=paths=source_relative:pb


grpcurl -plaintext -d '{"greeting": "Budi"}' localhost:50051 greet.GreetService/Greet