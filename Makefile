gen:
	protoc --go_out=. --go-grpc_out=. proto/order.proto proto/user.proto proto/tracking.proto

run:
	go run server/main.go

migrate:
	go run db/migrate.go

test:
	go test ./test/... -v
