package codec

import (
	"bytes"
	"github.com/hashicorp/go-msgpack/codec"
	"github.com/pkg/errors"
)

func Encode(obj interface{}) (raw []byte, err error) {
	buf := bytes.NewBuffer([]byte{})
	encoder := codec.NewEncoder(buf, &codec.SimpleHandle{})
	if err = encoder.Encode(obj); err != nil {
		return nil, errors.Wrapf(err, "can not encode obj(%+v)", obj)
	}

	return buf.Bytes(), nil
}

// decode raw data into structural object
func Decode(raw []byte, obj interface{}) (err error) {
	decoder := codec.NewDecoder(bytes.NewReader(raw), &codec.SimpleHandle{})
	return decoder.Decode(obj)
}
