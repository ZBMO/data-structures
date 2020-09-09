package main

import (
	"data-structures/linkedList"
	"fmt"
)

func main() {
	head := linkedList.CreateNode(5)
	head.Next = linkedList.CreateNode(100)

	fmt.Println(head.Data)
	fmt.Println(head.Next.Data)
}