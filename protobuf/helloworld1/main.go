package main

import (
	"fmt"

	"./types"
	"github.com/golang/protobuf/proto"
)

func main() {
	req := &types.Request{Data: "Hello Dabin"}

	// Marshal
	encoded, err := proto.Marshal(req)
	if err != nil {
		fmt.Printf("Encode to protobuf data error: %v", err)
	}

	// Unmarshal
	var unmarshaledReq types.Request
	err = proto.Unmarshal(encoded, &unmarshaledReq)
	if err != nil {
		fmt.Printf("Unmarshal to struct error: %v", err)
	}

	fmt.Printf("req: %v\n", req.String())
	fmt.Printf("unmarshaledReq: %v\n", unmarshaledReq.String())
}

/* Output
req: data:"Hello Dabin"
unmarshaledReq: data:"Hello Dabin"
*/
