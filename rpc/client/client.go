package main

import (
	"fmt"
	rpcdemo "go_crawler/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64

	err = client.Call("DemoService.Div", rpcdemo.Args{A: 10, B: 0}, &result)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("结果是：%v\n", result)
}
