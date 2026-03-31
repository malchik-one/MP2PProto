package registry

import (
	"fmt"
	"testing"

	"github.com/malchik-one/MP2PProto/transport/speac"
)

func TestInOp(t *testing.T) {
	err := speac.InitTags()
	if err != nil {
		t.Fatal(err)
	}

	InitStructOp()

	fun := OpRegString["send"]()

	fmt.Println(fun.GetTag())
}
