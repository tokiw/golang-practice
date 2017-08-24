package memo

import (
	"fmt"
	"sync"
	"testing"
)

func TestCancel(t *testing.T) {
	cancelErr := fmt.Errorf("cancel")

	f := func(key string, done <-chan struct{}) (interface{}, error) {
		<-done
		return nil, cancelErr
	}

	m := New(Func(f))

	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		v, err := m.Get("key", done)
		wg.Done()
		if v != nil || err != cancelErr {
			t.Errorf("got %v, %v; want %v, %v", v, err, nil, cancelErr)
		}
	}()
	close(done)

	wg.Wait()
}
