package datatypes

import (
	"fmt"
	"strings"

	"github.com/Stanlyzoolo/practiceGo/pkg/usecases"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateQueueObj(h *usecases.HistoryHandler) {
	var stack []string

	stack = append(stack, "World!")
	stack = append(stack, "Hello ")
	fmt.Print(stack)

	for len(stack) > 0 {
		stack[0] = ""
		stack = stack[1:]
		h.Info(stack)
	}
}

func CreateListNodeObj(h *usecases.HistoryHandler) {
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

func ReverseWordOne(s string, h *usecases.HistoryHandler) string {
	newWord := make([]rune, len(s))
	for i, v := range s {
		_ = append(newWord[:len(s)-1-i], v)
	}
	h.Info(newWord)
	return string(newWord)
}

func ReverseWordTwo(s string, h *usecases.HistoryHandler) string {
	var newWord []rune
	for _, v := range s {
		newWord = append([]rune{v}, newWord...)
	}
	h.Info(newWord)
	return string(newWord)
}

func ReverseWordsinSentence(s string, h *usecases.HistoryHandler) []string {
	data := strings.Fields(s)
	var newSent []string
	for _, v := range data {
		fmt.Println(v)
		newSent = append([]string{v}, newSent...)
	}
	h.Info(newSent)
	return newSent
}