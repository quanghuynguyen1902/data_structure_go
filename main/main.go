package main

import (
	"github.com/data_structure_go/linked_list"
)

func main(){
	l := linked_list.NewLinkedList(2)
	l.Add(0, 3)
	l.Add(1, 5)
	l.Add(3, 10)
	l.Add(5, 10)
	l.List()
}
