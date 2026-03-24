package Int64Type

import "github.com/malchik-one/MP2PProto/transport/entity/operations"

var _ operations.Operation = (*Int64Send)(nil)

type Int64Send struct {
	Tag  byte
	Data int64
}

func (i *Int64Send) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (i *Int64Send) UnmarshalBinary(data []byte) (n int, err error) {
	return 0, nil
}
