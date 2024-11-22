gen:
	@protoc --go_out=proto --go-grpc_out=proto proto/*.proto

run:
	go run server/main.go

test:
	go test ./test/... -v
