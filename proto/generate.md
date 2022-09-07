## Generating Protobuf files
To generate or regenerate the protobuf files found here, use `buf`.
The tools and the plugins can be installed with:
```sh
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
go install github.com/yoheimuta/protolint/cmd/protolint@latest
```
The `PATH` variable may need to be updated in order to include the Go binaries.