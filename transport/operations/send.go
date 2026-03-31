package operations

var _ Operation = (*Send)(nil)

type Send struct {
	Tag byte
}

func (s *Send) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (s *Send) UnmarshalBinary(data []byte) (n int, err error) {
	return 0, nil
}

func (s *Send) GetTag() (tag byte) {
	return s.Tag
}
