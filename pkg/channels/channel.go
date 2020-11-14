package channels

import (
	"time"

	"github.com/Stanlyzoolo/practiceGo/pkg/usecases"
)

func ReadTheChannel(ch chan string, h *usecases.HistoryHandler) {
	time.Sleep(5 * time.Second)
	h.Info("В канале ", <-ch)
}
