package request

import (
	"fmt"
	"testing"
)

func TestReq_Request(t *testing.T) {
	req := NewRequest()
	req.Url = "https://www.baidu.com/"
	resp, err := req.Request()
	fmt.Println(string(resp), err)
}
