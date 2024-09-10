
build-blockchain:
	protoc --go_out=../../../../  --go-grpc_out=../../../../  ./blockchain.proto

build-signer:build-blockchain
	protoc --go_out=../  --go-grpc_out=../  ./signer.proto


build-storePrivateKey:build-blockchain
	protoc --go_out=../  --go-grpc_out=../  ./storePrivateKey.proto


build-backend:build-blockchain
	protoc --go_out=../  --go-grpc_out=../ ./backend.proto
