package ByteType

import "MP2PProto/transport/entity/operations"

var _ operations.Operation = (*ByteSend)(nil)

type ByteSend struct {
	Tag  byte
	Data []byte
}

func (b *ByteSend) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (b *ByteSend) UnmarshalBinary(data []byte) (n int, err error) {
	return 0, nil
}
