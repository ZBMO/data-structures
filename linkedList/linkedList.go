package linkedList

type Node struct {
	Data int
	Next *Node
}

func CreateNode(data int) *Node {
	return &Node{data, nil}
}