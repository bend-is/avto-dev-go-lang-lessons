package fanout

import "sync"

func NewFanOut(in chan string, channelsOut []chan string) {
	var wg sync.WaitGroup

	wg.Add(len(channelsOut))

	action := func(out chan string) {
		defer wg.Done()
		for val := range in {
			out <- val
		}
	}

	for _, c := range channelsOut {
		go action(c)
	}

	wg.Wait()
}
