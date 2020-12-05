package fanin

import "sync"

func NewFanIn(channelsIn []chan string, out chan string) {
	var wg sync.WaitGroup

	wg.Add(len(channelsIn))

	action := func(c chan string) {
		defer wg.Done()
		for val := range c {
			out <- val
		}
	}

	for _, c := range channelsIn {
		go action(c)
	}

	wg.Wait()
}
