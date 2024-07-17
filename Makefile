generate:
	protoc ./proto/*.proto --go_out=./pkg && \
	protoc ./proto/*.proto --go-grpc_out=./pkg