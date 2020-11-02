package channels

import "time"

var ch = make(chan string)

func readAndWrite (ch chan string) string {
	time.Sleep(5*time.Second)
	message := <- ch
	go writeToChannel(ch, "Hard in study, easy in work")
	return message
}

func writeToChannel ( ch chan string, s string) {
	ch <- s
	close(ch)
}

