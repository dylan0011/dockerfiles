package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Hello Golang")

	// 服务器地址列表
	servers := []string{"192.168.3.2:2181", "192.168.3.2:2182", "192.168.3.2:2183"}

	println(servers)
	client, err := NewClient(servers, "/api", 10)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	node1 := &ServiceNode{"user", "192.168.3.2", 4000}
	node2 := &ServiceNode{"user", "192.168.3.2", 4001}
	node3 := &ServiceNode{"user", "192.168.3.2", 4002}
	if err := client.Register(node1); err != nil {
		panic(err)
	}
	if err := client.Register(node2); err != nil {
		panic(err)
	}
	if err := client.Register(node3); err != nil {
		panic(err)
	}
	nodes, err := client.GetNodes("user")
	if err != nil {
		panic(err)
	}
	for _, node := range nodes {
		fmt.Println(node.Host, node.Port)
	}

}