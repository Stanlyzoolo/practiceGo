package datatypes

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"container/list"
)

type dataHistoryHandler struct {
	Samples string
	Logger	*logrus.Logger
}

type ListNode struct {
	Val int
	Next *ListNode
}

func createQueueObj (d* dataHistoryHandler) {
	var stack []string

	stack = append(stack, "World!")
	stack = append(stack, "Hello ")
	fmt.Print(stack)

	n := len(stack)
	if n>0 {
		stack = stack [:n-1]
	}

	f := (d.DataSamples, stack)
	d.Samples = append(d.Samples, stack)
	d.Logger.Info("Info: ", d.Samples, stack)
}

func createListNodeObj () {
	obj1 := &ListNode{Val: 1}
	obj2 := &ListNode{Val: 2}
	obj1.Next = obj2
	obj3 := &ListNode{Val: 3}
	obj2.Next = obj3
	fmt.Println(obj1.Val, obj2.Val, obj3.Val)
	iter := obj1
	for iter != nil {
		fmt.Println(iter.Val)
		iter = iter.Next
	}
}


