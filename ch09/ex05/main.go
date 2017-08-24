package main

import "time"
import "fmt"

func main() {
	ch := make(chan int)
	t := time.After(1 * time.Second)
	var i int64

	// coutup go routine
	go func() {
		ch <- 1
		for {
			select {
			case <-t:
				close(ch)
			default:
				i++
				ch <- <-ch
			}
		}
	}()

	go func() {
		for {
			ch <- <-ch
		}
	}()

	<-t
	fmt.Printf("Count: %d", i)
}
