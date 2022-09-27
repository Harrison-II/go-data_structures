package main

import (
	"fmt"

	"github.com/harrison-ii/go/data-structures/heaps/max"
)

func main() {

	m := &max.MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 40, 50, 2, 3, 76, 7, 23, 6}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	fmt.Println("Max heap tree", m)

	for i := 0; i < 11; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
