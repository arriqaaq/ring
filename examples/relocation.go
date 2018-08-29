package main

import (
	"fmt"
	"github.com/arriqaaq/ring"
	"math/rand"
)

func main() {
	nodes := []string{"a", "b", "c"}
	hashRing := ring.NewRing(nodes, 40)

	keyCount := 1000000
	distribution := make(map[string]int)
	key := make([]byte, 4)

	fmt.Println("with 3 nodes")
	for i := 0; i < keyCount; i++ {
		rand.Read(key)
		node, err := hashRing.Get(string(key))
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		distribution[node]++
	}
	for node, count := range distribution {
		fmt.Printf("node: %s, key count: %d\n", node, count)
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("with 4 nodes")
	hashRing.Add("d")
	distribution = make(map[string]int)

	for i := 0; i < keyCount; i++ {
		rand.Read(key)
		node, err := hashRing.Get(string(key))
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		distribution[node]++
	}
	for node, count := range distribution {
		fmt.Printf("node: %s, key count: %d\n", node, count)
	}

}
