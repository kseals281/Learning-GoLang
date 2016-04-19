package main

import "fmt"

type node struct {
	num int
	// next nodeTemplate
}

type list struct {
}

func main() {
	node4 := node{4}
	// node3 := nodeTemplate{3, &node4}
	// node2 := nodeTemplate{2, &node3}
	// node1 := nodeTemplate{1, &node2}

	fmt.Println("Node 4: ", node4.num)

}
