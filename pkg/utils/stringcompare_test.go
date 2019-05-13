package utils

import (
	"fmt"
	"testing"
)

func TestJsonToMapStr(t *testing.T) {
	mapStr := JsonToMapStr("", "", 1)
	fmt.Println(mapStr)
}
