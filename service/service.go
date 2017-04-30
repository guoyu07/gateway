package service

import (
	"bytes"
	"github.com/hashicorp/go-msgpack/codec"
	"github.com/yangyuqian/gateway/types"
)

// Service represents a group of backend services.
type Service struct {
	Name           string                // service name should be unique, treated as an identity
	Labels         []string              // used for filtering
	Upstream       *types.Upstream       // upstream can be nil
	Authenticators []types.Authenticator // client attempt is accept when there is at least one Authenticator accepts it
	Throttlers     []types.Throttler     // client attempt is limited to min rate
}

// encode object to raw data as payload
func (svc *Service) Encode() (data []byte, err error) {
	buf := bytes.NewBuffer([]byte{})
	encoder := codec.NewEncoder(buf, &codec.SimpleHandle{})
	if err = encoder.Encode(svc); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// decode raw data into structural object
func (svc *Service) Decode(raw []byte) (err error) {
	decoder := codec.NewDecoder(bytes.NewReader(raw), &codec.SimpleHandle{})
	return decoder.Decode(svc)
}
