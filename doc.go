// Package gateway contains entrypoints for tests
package gateway

// goto https://developers.google.com/protocol-buffers/docs/gotutorial for more
// details on go protobuf, grpc is a protocol based on HTTP/2, and protobuf is
// a customized text format encoded for effient marshalling/unmarshalling.

//go:generate protoc -I protobuf/ protobuf/service.proto --go_out=plugins=grpc:protobuf
//go:generate protoc -I protobuf/ protobuf/cluster.proto --go_out=plugins=grpc:protobuf
