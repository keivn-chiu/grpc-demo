.PHONY: clean, gen

clean:
	rm -rf api/*

api-gen:
	protoc -I pb pb/*.proto --go_out=plugins=grpc:api
