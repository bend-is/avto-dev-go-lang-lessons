package main

import (
	"context"
	"fmt"
	"time"

	"bendis/avto-dev-go-lang-lessons/task4/pkg/bufferedchan"
	"bendis/avto-dev-go-lang-lessons/task4/pkg/fanin"
	"bendis/avto-dev-go-lang-lessons/task4/pkg/fanout"
	"bendis/avto-dev-go-lang-lessons/task4/pkg/pipe"
)

func main() {
	fmt.Printf("Pipe run example:\n\n")

	execPipe()

	fmt.Printf("\nBufferedchan run example:\n\n")

	execBufferedChan()

	fmt.Printf("\nFanin run example:\n\n")

	execFunIn()

	fmt.Printf("\nFanout run example:\n\n")

	execFunOut()
}

func execPipe() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	output := make(chan string)
	input := gen("Hello", "How", "Are", "You")

	p := pipe.NewPipe(ctx, input, output, addTimestamps)

	go func() { defer close(output); p.Run() }()

	for val := range output {
		fmt.Printf("%s\n", val)
	}
}

func execBufferedChan() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := bufferedchan.NewChan(ctx, gen("Hello", "How", "Are", "You", "?"), 2)

	go ch.Run()

	for val := range ch.C() {
		fmt.Printf("%s\n", val)
	}
}

func execFunIn() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := make(chan string)
	in1 := gen("In1.1 Input", "In1.2 Input", "In1.3 Input")
	in2 := gen("In2.1 Input", "In2.2 Input", "In2.3 Input")
	in3 := gen("In3.1 Input", "In3.2 Input", "In3.3 Input")

	fIn := fanin.NewFanIn(ctx, []chan string{in1, in2, in3}, out)

	go func() {
		defer close(out)
		fIn.Detach(in2)
		fIn.Run()
	}()

	for val := range out {
		fmt.Printf("%s\n", val)
	}

}

func execFunOut() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out1 := make(chan string)
	out2 := make(chan string)
	out3 := make(chan string)
	in := gen("One", "Two", "Three", "Four", "Five", "Six")

	fOut := fanout.NewFanOut(ctx, in, []chan string{out1, out2, out3})

	go func() {
		defer func() { close(out1); close(out2); close(out3) }()
		fOut.Detach(out2)
		fOut.Run()
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
