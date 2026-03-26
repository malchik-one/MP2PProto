package initialize

import (
	"github.com/malchik-one/MP2PProto/transport/entity/operations"
	"github.com/malchik-one/MP2PProto/transport/entity/operations/ByteType"
)

var bytSen = &ByteType.ByteSend{
	TagSend: 5,
}

var MapOperationsString = map[string]operations.Operation{
	"ByteSend": bytSen,
}

var MapOperationsByte = map[byte]operations.Operation{
	5: bytSen,
}
