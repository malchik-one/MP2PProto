package operations

var _ Operation = (*Payload)(nil)

type Payload struct {
	Tag     byte
	Payload []byte
}

func (b *Payload) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

func (b *Payload) UnmarshalBinary(data []byte) (n int, err error) {
	return 0, nil
}

func (b *Payload) GetTag() (tagByt byte) {
	return b.Tag
}
