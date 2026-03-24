package StringType

import "MP2PProto/transport/entity/operations"

var _ operations.Operation = (*StringSend)(nil)

type StringSend struct {
	Tag  byte
	Data string
}

func (s *StringSend) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (s *StringSend) UnmarshalBinary(data []byte) (n int, err error) {
	return 0, nil
}
