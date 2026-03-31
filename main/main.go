package main

import (
	"fmt"

	"github.com/malchik-one/MP2PProto/transport/operations"
	"github.com/malchik-one/MP2PProto/transport/registry"
	"github.com/malchik-one/MP2PProto/transport/speac"
)

func main() {
	speac.InitTags()
	registry.InitStructOp()

	op := registry.OpRegString["payload"]()

	if s, ok := op.(*operations.Payload); ok {
		fmt.Println(s.Payload)
	} else {
		fmt.Println("Это не Send")
	}
}
