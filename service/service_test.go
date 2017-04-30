package service

import (
	"reflect"
	"testing"
)

func TestEncodeDecode_valid(t *testing.T) {
	for _, cs := range []struct {
		name string
		svc  *Service
	}{
		{"service with only name", &Service{Name: "svc1"}},
	} {
		t.Run(cs.name, func(t *testing.T) {
			raw, err := cs.svc.Encode()
			if err != nil {
				t.Errorf("can not encode service(%+v) due to err(%+v)", cs.svc, err)
			}

			newsvc := &Service{}
			if err := newsvc.Decode(raw); err != nil {
				t.Errorf("can not decode raw data(%s) into structural object due to(%+v)", raw, err)
			}
			if !reflect.DeepEqual(cs.svc, newsvc) {
				t.Errorf("different data when encode and decode")
			}
		})
	}
}
