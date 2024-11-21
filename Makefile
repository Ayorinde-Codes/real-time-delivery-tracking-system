gen:
	@protoc --go_out=. --go-grpc_out=. proto/*.proto

run:
	go run server/main.go

test:
	go test ./test/... -v
