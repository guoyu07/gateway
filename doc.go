// Package gateway contains entrypoints for tests
package gateway

// goto https://developers.google.com/protocol-buffers/docs/gotutorial for more
// details on go protobuf, grpc is a protocol based on HTTP/2, and protobuf is
// a customized text format encoded for effient marshalling/unmarshalling.
//go:generate protoc --gogo_out=import_path=github.com/yangyuqian/gateway/store,Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogopoto:. ./store/service.proto
