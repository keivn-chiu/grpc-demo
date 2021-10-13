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

phone-classify-server:
	go build -o phone-classify-server ./cmd/phone-classify-server/server.go ; \
	mv phone-classify-server ./bin/

phone-classify-client:
	go build -o phone-classify-client ./cmd/phone-classify-client/client.go ; \
	mv phone-classify-client ./bin/

string-join-server:
	go build -o string-join-server ./cmd/string-join-server/server.go ; \
	mv string-join-server ./bin/

string-join-client:
	go build -o string-join-client ./cmd/string-join-client/client.go ; \
	mv string-join-client ./bin/