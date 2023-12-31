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

func (huffHeap *HuffHeap) PrintHeap() {
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

func (huffHeap *HuffHeap) ToHuffTree() *HuffTreeNode {
	var tmp1, tmp2, tmp3 *HuffTreeNode = nil, nil, nil

	for huffHeap.Len() > 1 {
		tmp1 = heap.Pop(huffHeap).(*HuffTreeNode)
		tmp2 = heap.Pop(huffHeap).(*HuffTreeNode)
		tmp3 = CreateInternalNode(tmp1.Weight+tmp2.Weight, tmp1, tmp2)
		heap.Push(huffHeap, tmp3)
	}

	return tmp3
}

func encodeHuffTreeRec(root *HuffTreeNode, path string, encoding map[byte]string) {
	if root.IsLeaf {
		encoding[root.Element] = path
		return
	}

	encodeHuffTreeRec(root.Left, path+"0", encoding)
	encodeHuffTreeRec(root.Right, path+"1", encoding)
}

func (root *HuffTreeNode) Encode() map[byte]string {
	encodingTable := make(map[byte]string)
	encodeHuffTreeRec(root, "", encodingTable)

	return encodingTable
}
