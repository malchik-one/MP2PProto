package operations

type Operation interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(data []byte) (n int, err error)
	GetTag() (tagByt byte)
}
