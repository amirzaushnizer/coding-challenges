package hufftree

import (
	"container/heap"
	"fmt"
)

type HuffTreeNode struct {
	Weight  int
	Element byte
	Left    *HuffTreeNode
	Right   *HuffTreeNode
	IsLeaf  bool
}

type HuffHeap []*HuffTreeNode

func (huffHeap HuffHeap) Len() int {
	return len(huffHeap)
}

func (huffHeap HuffHeap) Less(i, j int) bool {
	return huffHeap[i].Weight < huffHeap[j].Weight
}

func (huffHeap HuffHeap) Swap(i, j int) {
	huffHeap[i], huffHeap[j] = huffHeap[j], huffHeap[i]
}

func (huffHeap *HuffHeap) Push(x interface{}) {
	*huffHeap = append(*huffHeap, x.(*HuffTreeNode))
}

func (huffHeap *HuffHeap) Pop() interface{} {
	old := *huffHeap
	n := len(old)
	x := old[n-1]
	*huffHeap = old[0 : n-1]
	return x
}

func CreateHuffHeap(frequencies map[byte]int) *HuffHeap {
	huffHeap := &HuffHeap{}
	for char, count := range frequencies {
		*huffHeap = append(*huffHeap, CreateLeaf(count, char))
	}
	heap.Init(huffHeap)
	return huffHeap
}

// print heap
func PrintHeap(huffHeap *HuffHeap) {
	for huffHeap.Len() > 0 {
		node := heap.Pop(huffHeap).(*HuffTreeNode)
		fmt.Printf("%q: %d\n", node.Element, node.Weight)
	}
}

func CreateLeaf(weight int, element byte) *HuffTreeNode {
	return &HuffTreeNode{weight, element, nil, nil, true}
}

func CreateInternalNode(weight int, left *HuffTreeNode, right *HuffTreeNode) *HuffTreeNode {
	return &HuffTreeNode{weight, 0, left, right, false}
}

// func CreateHuffTree(frequenciesHeap *HuffHeap) *HuffTreeNode {
// 	subTrees := make([]*HuffTreeNode, 0)

// 	for frequenciesHeap.Len() > 0 {
// 		frequency := heap.Pop(frequenciesHeap).(frequenciesminheap.Frequency)

// 		if len(subTrees) <= 1 {
// 			subTrees = append(subTrees, CreateLeaf(frequency.Count, frequency.Char))
// 			continue
// 		}

// 	}

// }
