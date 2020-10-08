package datatypes

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"container/list"
	"strings"
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

		for len(stack) > 0 {
			stack[0] = ""
			stack = stack[1:]
			
			f := (d.Samples, stack)
			d.Samples = f
			d.Logger.Info("Info: ", d.Samples)
		}
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

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseRange(s string) string {
	data :=  strings.Fields(s)
	for i,j := range data {
		runes := []rune(j)
		for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
}
