package channels

import (
	"context"
	"fmt"
)

func doSomething() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go doSomethingElse(ctx, "аргумент")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped")
			return

		default:
			fmt.Println("All is good")
			cancel()
		}
	}
}

func doSomethingElse(ctx context.Context, arg string) {
	<-ctx.Done()
	var newWord []rune
	for _, v := range arg {
		newWord = append([]rune{v}, newWord...)
	}
	fmt.Println(newWord)
}
