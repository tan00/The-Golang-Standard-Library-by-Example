package main

import (
	"container/heap"
	"fmt"
)

// type Interface interface {
// 	sort.Interface
// 	Push(x interface{}) // add x as element Len()
// 	Pop() interface{}   // remove and return element Len() - 1.
// }

//用使用堆, 需要实现五个函数 Push Pop Len Less Swap

func (h *Persons) Push(x interface{}) {
	*h = append(*h, x.(Person))
}

func (h *Persons) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap() {
	var (
		persons = &Persons{
			{18, 1, "James"},
			{5, 1, "Sfsaf"},
			{23, 0, "WEAFs"},
			{76, 1, "dvweg"},
			{2, 0, "Wefwrqw"},
			{36, 1, "Werfscd"},
		}
	)

	heap.Init(persons)
	fmt.Print(persons)

	for range *persons {
		v := heap.Pop(persons)
		fmt.Println("pop", v)
	}

}
