package usecases

import (
	"strconv"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Celsius float32
type Fahrenheit float32
type HistoryHandler struct{
	Samples []Fahrenheit
	Logger	*logrus.Logger
}

func CheckAndSave(h *HistoryHandler, t interface{}, n int) {
	for i:= 0; i<n; i++ {
		f:= toFahrenheit(t) + Fahrenheit(i)
		h.Samples = append(h.Samples, f)
		h.Logger.Info("Предыдущие изменения и текущее значение - ", h.Samples, f)
	}
}

func toFahrenheit(t interface{}) Fahrenheit {
	var temp Fahrenheit
	switch c := t.(type) {
		case Celsius:
			temp = Fahrenheit(c*(9/5) + 32)	 
			fmt.Printf("Работаем с температурой в Цельсиях: %v\n", temp)
		case int:
			temp = Fahrenheit(c*(9/5) + 32)	 
			fmt.Printf("Работаем с температурой в интах: %v\n", temp)
		case nil:        fmt.Printf("Входные данные отсутствуют\n", temp)
		case string:
			conv, _ := strconv.Atoi(c)
			temp = Fahrenheit(conv*(9/5) + 32)	 
			fmt.Printf("Работаем с температурой со строками: %v\n", temp)
		default:
			fmt.Printf("Входные данные не идентифицируются\n", temp)
		}
	return temp
}
