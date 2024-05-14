.PHONY: options
options:
	protoc --proto_path=./ --go_out=paths=source_relative:./  ./annotations/options.proto

.PHONY: example
# generate api proto
example:
	protoc --proto_path=./ --go_out=paths=source_relative:./tests/  --go-grpc_out=paths=source_relative:./tests  --grant_out=paths=source_relative:./tests ./example.proto


.PHONY: clean
clean:
	rm -rf tests/*.pb.go		
