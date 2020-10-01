package datatypes

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"container/list"
)

type dataHistoryHandler struct {
	DataSamples *list.List
	DataLogger	*logrus.Logger
}

func createQueueObj (d* dataHistoryHandler) {
	queue := list.New()
	var element1 string
	var element2 string

	fmt.Print("Входные данные №1")
	fmt.Scan(&element1)

	fmt.Print("Входные данные №2")
	fmt.Scan(&element2)

	
	e1 := queue.PushFront(element1)
	e2 := queue.PushBack(element2)

	fmt.Print(queue)
	
	for queue.Len() > 0 {
		queue.Remove(e1)
		queue.Remove(e2)		
	}

	for e := queue.Front(); e != nil; e = e.Next() {
		d.DataSamples = queue
		d.DataLogger.Info("Работа с очередью FIFO", d.DataSamples, queue)
		fmt.Print(d.DataSamples)
	}
}
