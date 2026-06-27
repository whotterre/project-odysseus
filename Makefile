proto:
	rm ./src/internal/grpc/proto/*
	protoc --proto_path=proto --go_out=internal/grpc --go_opt=paths=source_relative --go-grpc_out=internal/grpc --go-grpc_opt=paths=source_relative `       --plugin=protoc-gen-go=C:\Users\USER\go\bin\protoc-gen-go.exe `       --plugin=protoc-gen-go-grpc=C:\Users\USER\go\bin\protoc-gen-go-grpc.exe proto/telemetry.proto
run:
	cd src/cmd && go build . -o server
	./server

PHONY: proto run