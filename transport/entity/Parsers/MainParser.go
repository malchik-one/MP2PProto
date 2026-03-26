package parsers

import (
	"github.com/malchik-one/MP2PProto/transport/entity/operations"
	"github.com/malchik-one/MP2PProto/transport/entity/operations/initialize"
)

func MainParseToBinary(sliceOp []operations.Operation) []byte {

	var b []byte = make([]byte, len(sliceOp))

	for l, i := range sliceOp {
		b[l] = i.GetTag()
	}

	return b
}

func MainParseToStruct(b []byte) []operations.Operation {

	var ops []operations.Operation = make([]operations.Operation, len(b))

	for l, i := range b {
		ops[l] = initialize.MapOperationsByte[i]
	}

	return ops

}
