package main

import (
	"fmt"
	"time"

	"bendis/avto-dev-go-lang-lessons/task4/pkg/fanin"
	"bendis/avto-dev-go-lang-lessons/task4/pkg/fanout"
	"bendis/avto-dev-go-lang-lessons/task4/pkg/pipe"
)

func main() {
	fmt.Printf("Pipe run example:\n\n")

	input := gen("Hello", "How", "Are", "You")
	output := make(chan string)

	go func() {
		defer close(output)
		pipe.NewPipe(input, output, addTimestamps)
	}()

	for val := range output {
		fmt.Printf("%s\n", val)
	}

	fmt.Printf("\n\nFanin run example:\n\n")

	out := make(chan string)
	in1 := gen("In1 Input", "In1 Input", "In1 Input")
	in2 := gen("In2 Input", "In2 Input", "In2 Input")
	in3 := gen("In3 Input", "In3 Input", "In3 Input")

	go func() {
		defer close(out)
		fanin.NewFanIn([]chan string{in1, in2, in3}, out)
	}()

	for val := range out {
		fmt.Printf("%s\n", val)
	}

	fmt.Printf("\n\nFanout run example:\n\n")

	in := gen("One", "Two", "Three", "Four", "Five", "Six")
	out1 := make(chan string)
	out2 := make(chan string)
	out3 := make(chan string)

	go func() {
		defer func() { close(out1); close(out2); close(out3) }()
		fanout.NewFanOut(in, []chan string{out1, out2, out3})
	}()

	for {
		select {
		case val, ok := <-out1:
			if !ok {
				return
			}
			fmt.Printf("response from out1: %s\n", val)
		case val, ok := <-out2:
			if !ok {
				return
			}
			fmt.Printf("response from out2: %s\n", val)
		case val, ok := <-out3:
			if !ok {
				return
			}
			fmt.Printf("response from out3: %s\n", val)
		}
	}
}

func addTimestamps(val string) string {
	return fmt.Sprintf("%s_%d", val, time.Now().Unix())
}

func gen(values ...string) chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for _, n := range values {
			ch <- n
		}
	}()

	return ch
}
