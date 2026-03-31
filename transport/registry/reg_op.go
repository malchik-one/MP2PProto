package registry

import (
	"github.com/malchik-one/MP2PProto/transport/operations"
	"github.com/malchik-one/MP2PProto/transport/speac"
)

var OpRegString = make(map[string]func() operations.Operation)

var OpRegByte = map[byte]operations.Operation{}

func InitStructOp() {
	OpRegString["send"] = func() operations.Operation {
		return &operations.Send{Tag: speac.TagsReg["send"]}
	}
	OpRegString["payload"] = func() operations.Operation {
		return &operations.Payload{Tag: speac.TagsReg["payload"], Payload: []byte{}}
	}
}
