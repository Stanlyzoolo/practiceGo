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

func ReverseStringOne(s string) string {	
	newWord := make([]rune, len(s))
	for i, v  := range s {
		_ = append(newWord[:len(s)-1-i], v)
	}
	return string(newWord)
}

func ReverseStringTwo(s string) string {
	var newWord []rune
	for _, v := range s {
		newWord = append([]rune{v}, newWord...)
	}

	return string(newWord)
}


func ReverseRange(s string) []string {
	data := strings.Fields(s)
	var newSent []string
	for _, v := range data {
		fmt.Println(v)
		newSent = append([]string{v}, newSent...)
	}
	return newSent
}
