package speac

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type SpecOperations struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Tag  byte   `json:"tag"`
}

var TagsReg = make(map[string]byte)

func InitTags() error {

	_, b, _, _ := runtime.Caller(0)
	targetPath := filepath.Join(filepath.Dir(b), "speacs/st_speac.json")

	filestSpe, err1 := os.Open(targetPath)
	if err1 != nil {
		return err1
	}
	defer filestSpe.Close()

	var spOp []SpecOperations

	err2 := json.NewDecoder(filestSpe).Decode(&spOp)
	if err2 != nil {
		return err2
	}

	for i := 0; i <= len(spOp)-1; i++ {
		TagsReg[spOp[i].Name] = spOp[i].Tag
	}

	return nil
}
