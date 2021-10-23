all:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative *.proto
	protoc --php_out=./php --grpc_out=./php --plugin=protoc-gen-grpc=/Users/mac/Dev/grpc/cmake/build/grpc_php_plugin *.proto
	protoc --java_out=./java *.proto
	#go generate .
