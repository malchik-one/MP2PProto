package ByteType

import "testing"

func TestMarshalBinary(t *testing.T) {

	st := ByteSend{
		5,
	}

	slice, err := st.MarshalBinary()
	if err != nil {
		println(err.Error())
	}

	println("slice: ", slice[0])

}
