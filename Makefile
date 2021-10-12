.PHONY: init, clean, gen, product-client, product-server

init:
	mkdir bin ; \
	go mod tidy

clean:
	rm -rf api/* ; rm -rf bin/*

api-gen:
	protoc -I pb pb/*.proto --go_out=plugins=grpc:api

product-client:
	go build -o product-client ./cmd/product-client/client.go ; \
	mv product-client ./bin/


product-server:
	go build -o product-server ./cmd/product-server/server.go ; \
	mv product-server ./bin/