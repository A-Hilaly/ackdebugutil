package ackdebugutil

import (
	"encoding/json"
	"fmt"
)

func marshallObjectStr(i any) {
	b, _ := json.MarshalIndent(i, "", "    ")
	fmt.Println(string(b))
}

func DebugJSON(i any) {
	marshallObjectStr(i)
}

func DebugJSONAll(is ...any) {
	for _, i := range is {
		marshallObjectStr(i)
	}
}
