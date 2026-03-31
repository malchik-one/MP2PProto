package speac

import "testing"

func TestInTag(t *testing.T) {
	err := InitTags()
	if err != nil {
		t.Fatal(err)
	}

	println(TagsReg["send"])
}
