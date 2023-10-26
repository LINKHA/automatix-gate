#!/usr/bin/env bash
protoc -I. --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. \
		--grpc-gateway_opt=paths=source_relative \
		--grpc-gateway_opt=logtostderr=true \
		--openapiv2_out=. \
		--openapiv2_opt=logtostderr=true,allow_delete_body=true \
		--descriptor_set_out=dynamic_call.protoset \
		--include_imports \
		dynamic_call.proto
