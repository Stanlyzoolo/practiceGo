package channel

// import (
// 	"time"
// )

func readTheChannel(ch chan string) {
	<-ch
}

// func main() {
// 	ch := make(chan string)

// 	go readTheChannel(ch)
// 	time.Sleep(5 * time.Second)

// 	go func(ch chan string, s string) {
// 		ch <- s
// 	}(ch, "Work or not, tell me please, master?")
// }
