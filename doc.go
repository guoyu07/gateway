// Package gateway contains entrypoints for tests
package gateway

//go:generate protoc --gogo_out=import_path=github.com/yangyuqian/gateway/store,Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogopoto:. ./store/service.proto
